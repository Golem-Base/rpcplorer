package handlers

import (
	"log/slog"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/Golem-Base/rpcplorer/templates"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BlockHandler returns an http.HandlerFunc that shows details for a specific block
func BlockHandler(client *ethclient.Client, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract block identifier from URL
		blockID := r.PathValue("id")
		if blockID == "" {
			http.NotFound(w, r)
			return
		}

		var block interface{} // Will hold the actual block
		var err error

		// Check if the blockID is a hash (starts with 0x) or a number
		if strings.HasPrefix(blockID, "0x") {
			// It's a hash
			hash := common.HexToHash(blockID)
			block, err = client.BlockByHash(r.Context(), hash)
		} else {
			// It's a number
			blockNum, err := strconv.ParseUint(blockID, 10, 64)
			if err != nil {
				log.Error("Invalid block number", "blockID", blockID, "error", err)
				http.Error(w, "Invalid block number format", http.StatusBadRequest)
				return
			}
			block, err = client.BlockByNumber(r.Context(), big.NewInt(int64(blockNum)))
		}

		if err != nil {
			log.Error("Failed to get block", "blockID", blockID, "error", err)
			http.Error(w, "Block not found", http.StatusNotFound)
			return
		}

		// Convert the generic interface{} to *types.Block
		typedBlock, ok := block.(*types.Block)
		if !ok {
			log.Error("Failed to convert block to *types.Block", "blockID", blockID)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Set default gas values
		gasUsed := typedBlock.GasUsed()
		gasLimit := typedBlock.GasLimit()

		// Calculate approximate block rewards (this is a simplified calculation)
		// In Ethereum, the block reward has changed over time. Currently it's 2 ETH in Ethereum
		// but this is a simplification
		rewards := big.NewInt(2 * 1e18) // 2 ETH in wei

		// Render the block template
		component := templates.Block(typedBlock, rewards, gasUsed, gasLimit)
		err = component.Render(r.Context(), w)
		if err != nil {
			log.Error("Failed to render block template", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
