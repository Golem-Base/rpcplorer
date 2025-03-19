package handlers

import (
	"math/big"
	"net/http"

	"github.com/Golem-Base/rpcplorer/templates"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type APIHandler struct {
	client *ethclient.Client
}

func NewAPIHandler(client *ethclient.Client) *APIHandler {
	return &APIHandler{
		client: client,
	}
}

// GetLatestBlocks fetches the latest blocks and returns them rendered as HTML
func (h *APIHandler) GetLatestBlocks(w http.ResponseWriter, r *http.Request) {
	// Fetch latest blocks (this is a simplified example)
	// In a real app, you would fetch multiple blocks
	blocks := make([]*types.Block, 0)

	// Get the latest block number
	header, err := h.client.HeaderByNumber(r.Context(), nil)
	if err != nil {
		// Return empty blocks list if there's an error
		renderBlocksList(w, r, blocks)
		return
	}

	// Calculate how many blocks we can fetch
	numBlocksToFetch := 5
	if header.Number.Cmp(big.NewInt(5)) < 0 {
		numBlocksToFetch = int(header.Number.Int64() + 1) // +1 because block 0 exists
	}

	// Fetch the latest blocks
	for i := 0; i < numBlocksToFetch; i++ {
		blockNum := new(big.Int).Sub(header.Number, big.NewInt(int64(i)))
		block, err := h.client.BlockByNumber(r.Context(), blockNum)
		if err != nil {
			continue
		}
		blocks = append(blocks, block)
	}

	// Render just the blocks list component
	renderBlocksList(w, r, blocks)
}

// GetLatestTransactions fetches transactions from the latest block
func (h *APIHandler) GetLatestTransactions(w http.ResponseWriter, r *http.Request) {
	// Fetch the latest blocks until we have at least 5 transactions
	header, err := h.client.HeaderByNumber(r.Context(), nil)
	if err != nil {
		// Return empty transactions list if there's an error
		renderTransactionsList(w, r, []*types.Block{})
		return
	}

	blocks := make([]*types.Block, 0)
	txCount := 0
	maxBlocks := 10 // Limit the number of blocks to check to avoid excessive fetching

	for i := 0; i < maxBlocks && txCount < 5; i++ {
		blockNum := new(big.Int).Sub(header.Number, big.NewInt(int64(i)))
		if blockNum.Sign() < 0 {
			break // Don't go below block 0
		}

		block, err := h.client.BlockByNumber(r.Context(), blockNum)
		if err != nil {
			continue
		}

		blocks = append(blocks, block)
		txCount += len(block.Transactions())

		if txCount >= 5 {
			break
		}
	}

	// Render the transactions list component with the collected blocks
	renderTransactionsList(w, r, blocks)
}

// RegisterRoutes registers all API routes
func RegisterAPIRoutes(mux *http.ServeMux, client *ethclient.Client) {
	handler := NewAPIHandler(client)

	// API routes for HTMX
	mux.HandleFunc("GET /api/blocks", handler.GetLatestBlocks)
	mux.HandleFunc("GET /api/transactions", handler.GetLatestTransactions)
}

// Render the blocks list component using templ
func renderBlocksList(w http.ResponseWriter, r *http.Request, blocks []*types.Block) {
	err := templates.BlocksList(blocks).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering blocks list", http.StatusInternalServerError)
	}
}

// Render the transactions list component using templ
func renderTransactionsList(w http.ResponseWriter, r *http.Request, blocks []*types.Block) {
	err := templates.TransactionsList(blocks).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Error rendering transactions list", http.StatusInternalServerError)
	}
}
