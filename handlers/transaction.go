package handlers

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Golem-Base/rpcplorer/templates"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TransactionHandler returns an http.HandlerFunc for displaying transaction details
func TransactionHandler(client *ethclient.Client, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract transaction hash from URL
		txHash := r.PathValue("hash")
		if txHash == "" {
			http.NotFound(w, r)
			return
		}

		// Validate transaction hash format
		if len(txHash) != 66 && len(txHash) != 64 {
			log.Error("Invalid transaction hash", "hash", txHash)
			http.Error(w, "Invalid transaction hash format", http.StatusBadRequest)
			return
		}

		// Ensure hash starts with 0x
		if len(txHash) == 64 {
			txHash = "0x" + txHash
		}

		// Convert string to hash
		hash := common.HexToHash(txHash)

		// Get transaction
		tx, isPending, err := client.TransactionByHash(r.Context(), hash)
		if err != nil {
			log.Error("Failed to get transaction", "error", err, "hash", txHash)
			http.Error(w, "Transaction not found", http.StatusNotFound)
			return
		}

		if isPending {
			log.Info("Transaction is pending", "hash", txHash)
			http.Error(w, "Transaction is pending and not yet mined", http.StatusAccepted)
			return
		}

		// Get transaction receipt
		receipt, err := client.TransactionReceipt(r.Context(), hash)
		if err != nil {
			log.Error("Failed to get transaction receipt", "error", err, "hash", txHash)
			http.Error(w, "Failed to get transaction details", http.StatusInternalServerError)
			return
		}

		// Get the block containing this transaction
		block, err := client.BlockByHash(r.Context(), receipt.BlockHash)
		if err != nil {
			log.Error("Failed to get block", "error", err, "blockHash", receipt.BlockHash.Hex())
			http.Error(w, "Failed to get transaction details", http.StatusInternalServerError)
			return
		}

		// Get current gas price
		gasPrice, err := client.SuggestGasPrice(r.Context())
		if err != nil {
			log.Error("Failed to get current gas price", "error", err)
			// Not critical, continue with nil gas price
			gasPrice = nil
		}

		// Get sender address
		from, err := getSender(r.Context(), client, tx)
		if err != nil {
			log.Error("Failed to get sender address", "error", err, "hash", txHash)
			http.Error(w, "Failed to get transaction details", http.StatusInternalServerError)
			return
		}

		// Render the transaction template
		component := templates.Transaction(tx, receipt, block, gasPrice, from)
		err = component.Render(r.Context(), w)
		if err != nil {
			log.Error("Failed to render template", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

// getSender retrieves the sender address from a transaction
func getSender(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (common.Address, error) {
	chainID, err := client.NetworkID(ctx)
	if err != nil {
		return common.Address{}, err
	}

	signer := types.LatestSignerForChainID(chainID)
	from, err := types.Sender(signer, tx)
	if err != nil {
		return common.Address{}, err
	}

	return from, nil
}
