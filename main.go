package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Golem-Base/rpcplorer/handlers"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/urfave/cli/v2"
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

			// Register API routes for HTMX
			handlers.RegisterAPIRoutes(mux, client)

			// Main page handler
			mux.HandleFunc("GET /", handlers.HomeHandler(client, log))

			// Transaction details page
			mux.HandleFunc("GET /tx/{hash}", handlers.TransactionHandler(client, log))

			// Block details page - can use either block number or hash
			mux.HandleFunc("GET /block/{id}", handlers.BlockHandler(client, log))

			// Blocks history page with pagination
			mux.HandleFunc("GET /blocks", handlers.BlocksHandler(client, log))

			// Address details page
			mux.HandleFunc("GET /address/{addr}", handlers.AddressHandler(client, log))

			// Search handler
			mux.HandleFunc("GET /search", handlers.SearchHandler(log))

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
