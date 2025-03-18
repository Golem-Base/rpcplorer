package main

import (
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"os"

	"github.com/Golem-Base/rpcplorer/templates"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg := struct {
		nodeURL string
	}{}

	app := &cli.App{
		Name: "rpcplorer",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "node-url",
				Usage:       "Node URL",
				Destination: &cfg.nodeURL,
				EnvVars:     []string{"NODE_URL"},
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {

			log.Info("Connecting to Ethereum node", "url", cfg.nodeURL)

			client, err := ethclient.Dial(cfg.nodeURL)
			if err != nil {
				log.Error("Failed to connect to the Ethereum node", "error", err)
				return err
			}
			defer client.Close()

			// Verify connection by getting network ID
			networkID, err := client.NetworkID(c.Context)
			if err != nil {
				log.Error("Failed to get network ID", "error", err)
				return err
			}
			log.Info("Connected to Ethereum network", "networkID", networkID)

			mux := http.NewServeMux()
			mux.HandleFunc(
				"/",
				func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path != "/" {
						http.NotFound(w, r)
						return
					}

					// Get the latest block number
					latestBlockNumber, err := client.BlockNumber(r.Context())
					if err != nil {
						log.Error("Failed to get latest block number", "error", err)
						http.Error(w, "Failed to get latest block data", http.StatusInternalServerError)
						return
					}

					// Determine how many blocks we can fetch
					numBlocksToFetch := 5
					if latestBlockNumber < 5 {
						numBlocksToFetch = int(latestBlockNumber) + 1 // +1 because block 0 (genesis) exists
					}

					// Use errgroup to get the blocks in parallel
					g, ctx := errgroup.WithContext(r.Context())
					blocks := make([]*types.Block, numBlocksToFetch)

					// Launch goroutines to fetch blocks in parallel
					for i := 0; i < numBlocksToFetch; i++ {
						// Capture loop variable
						i := i
						g.Go(func() error {
							blockNumber := latestBlockNumber - uint64(i)
							block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
							if err != nil {
								return fmt.Errorf("failed to get block %d: %w", blockNumber, err)
							}
							blocks[i] = block
							return nil
						})
					}

					// Wait for all goroutines to complete
					if err := g.Wait(); err != nil {
						log.Error("Failed to get blocks", "error", err)
						http.Error(w, "Failed to get latest block data", http.StatusInternalServerError)
						return
					}

					// Filter out nil blocks (in case we had any issues)
					validBlocks := make([]*types.Block, 0, numBlocksToFetch)
					for _, block := range blocks {
						if block != nil {
							validBlocks = append(validBlocks, block)
						}
					}

					log.Info("Retrieved latest blocks", "count", len(validBlocks))

					component := templates.Index(validBlocks)

					err = component.Render(r.Context(), w)
					if err != nil {
						log.Error("failed to render template", "error", err)
						http.Error(w, "Internal Server Error", http.StatusInternalServerError)
						return
					}
				},
			)

			log.Info("Starting server on :8080")
			return http.ListenAndServe(":8080", mux)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Error("error running app", "error", err)
		os.Exit(1)
	}
}
