package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// SearchHandler handles search requests and redirects to the appropriate URL
func SearchHandler(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			w.Header().Set("HX-Redirect", "/")
			return
		}

		// Trim any whitespace
		query = strings.TrimSpace(query)

		// Check if it's an Ethereum address
		if common.IsHexAddress(query) {
			w.Header().Set("HX-Redirect", "/address/"+query)
			return
		}

		// Check if it looks like a transaction hash (0x followed by 64 hex chars)
		if strings.HasPrefix(query, "0x") && len(query) == 66 {
			w.Header().Set("HX-Redirect", "/tx/"+query)
			return
		}

		// Check if it's a block number
		if blockNum, err := strconv.ParseUint(query, 10, 64); err == nil {
			w.Header().Set("HX-Redirect", "/block/"+strconv.FormatUint(blockNum, 10))
			return
		}

		// Check if it looks like a block hash (0x followed by 64 hex chars, same format as tx hash)
		if strings.HasPrefix(query, "0x") && len(query) == 66 {
			w.Header().Set("HX-Redirect", "/block/"+query)
			return
		}

		// If we can't determine, redirect to homepage
		w.Header().Set("HX-Redirect", "/")
	}
}
