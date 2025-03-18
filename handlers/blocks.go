package handlers

import (
	"log/slog"
	"math/big"
	"net/http"
	"strconv"

	"github.com/Golem-Base/rpcplorer/templates"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/errgroup"
)

const (
	// Number of blocks to display per page
	blocksPerPage = 50
)

// BlocksHandler returns an http.HandlerFunc that shows a paginated list of blocks
func BlocksHandler(client *ethclient.Client, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the latest block number
		latestBlockNumber, err := client.BlockNumber(r.Context())
		if err != nil {
			log.Error("Failed to get latest block number", "error", err)
			http.Error(w, "Failed to get latest block data", http.StatusInternalServerError)
			return
		}

		// Parse the page parameter
		pageStr := r.URL.Query().Get("page")
		page := 1 // Default to first page
		if pageStr != "" {
			parsedPage, err := strconv.Atoi(pageStr)
			if err == nil && parsedPage > 0 {
				page = parsedPage
			}
		}

		// Calculate the range of blocks to fetch
		startBlock := int64(latestBlockNumber) - int64((page-1)*blocksPerPage) - int64(blocksPerPage) + 1
		endBlock := int64(latestBlockNumber) - int64((page-1)*blocksPerPage)

		// Ensure we don't go below block 0
		if startBlock < 0 {
			startBlock = 0
		}

		// Check if there are more pages
		hasNextPage := startBlock > blocksPerPage

		// Calculate how many blocks we can fetch
		numBlocksToFetch := int(endBlock - startBlock + 1)
		if numBlocksToFetch <= 0 {
			// No blocks to fetch for this page
			http.Error(w, "No blocks available for this page", http.StatusBadRequest)
			return
		}

		// Use errgroup to fetch blocks in parallel
		g, ctx := errgroup.WithContext(r.Context())
		blocks := make([]*types.Block, numBlocksToFetch)

		for i := 0; i < numBlocksToFetch; i++ {
			i := i
			blockNumber := uint64(endBlock - int64(i))

			g.Go(func() error {
				// Fetch block
				block, err := client.BlockByNumber(ctx, new(big.Int).SetUint64(blockNumber))
				if err != nil {
					return err
				}
				blocks[i] = block
				return nil
			})
		}

		// Wait for all goroutines to complete
		if err := g.Wait(); err != nil {
			log.Error("Failed to fetch blocks", "error", err)
			http.Error(w, "Failed to fetch block data", http.StatusInternalServerError)
			return
		}

		// Filter out nil blocks (in case we had any issues)
		validBlocks := make([]*types.Block, 0, numBlocksToFetch)
		for _, block := range blocks {
			if block != nil {
				validBlocks = append(validBlocks, block)
			}
		}

		log.Info("Retrieved blocks for page", "page", page, "count", len(validBlocks), "start", startBlock, "end", endBlock)

		// Render the blocks template
		component := templates.Blocks(validBlocks, page, hasNextPage, latestBlockNumber)
		err = component.Render(r.Context(), w)
		if err != nil {
			log.Error("Failed to render blocks template", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
