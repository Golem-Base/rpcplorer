package handlers

import (
	"fmt"
	"log/slog"
	"math/big"
	"net/http"

	"github.com/Golem-Base/rpcplorer/templates"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/errgroup"
)

// HomeHandler returns an http.HandlerFunc for the home page
func HomeHandler(client *ethclient.Client, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}
