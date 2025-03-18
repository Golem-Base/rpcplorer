package handlers

import (
	"log/slog"
	"net/http"

	"github.com/Golem-Base/rpcplorer/templates"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// AddressHandler returns an http.HandlerFunc for displaying address details
func AddressHandler(client *ethclient.Client, log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract address from URL
		addrHex := r.PathValue("addr")
		if addrHex == "" {
			http.NotFound(w, r)
			return
		}

		// Validate address format
		if !common.IsHexAddress(addrHex) {
			log.Error("Invalid address format", "address", addrHex)
			http.Error(w, "Invalid address format", http.StatusBadRequest)
			return
		}

		// Convert to common.Address
		address := common.HexToAddress(addrHex)

		// Get address balance
		balance, err := client.BalanceAt(r.Context(), address, nil)
		if err != nil {
			log.Error("Failed to get address balance", "address", addrHex, "error", err)
			http.Error(w, "Failed to get address details", http.StatusInternalServerError)
			return
		}

		// Check if the address is a contract
		code, err := client.CodeAt(r.Context(), address, nil)
		if err != nil {
			log.Error("Failed to check if address is a contract", "address", addrHex, "error", err)
			http.Error(w, "Failed to get address details", http.StatusInternalServerError)
			return
		}
		isContract := len(code) > 0

		// Get transaction count (nonce)
		txCount, err := client.NonceAt(r.Context(), address, nil)
		if err != nil {
			log.Error("Failed to get transaction count", "address", addrHex, "error", err)
			http.Error(w, "Failed to get address details", http.StatusInternalServerError)
			return
		}

		// Create the page data
		data := templates.AddressPageData{
			Address:          address,
			Balance:          balance,
			TransactionCount: txCount,
			IsContract:       isContract,
			ContractCode:     code,
		}

		// Render the template
		component := templates.Address(data)
		err = component.Render(r.Context(), w)
		if err != nil {
			log.Error("Failed to render address template", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}
