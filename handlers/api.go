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

	// Fetch the latest blocks
	for i := 0; i < 5; i++ {
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
	// Fetch the latest block to get its transactions
	header, err := h.client.HeaderByNumber(r.Context(), nil)
	if err != nil {
		// Return empty transactions list if there's an error
		renderTransactionsList(w, r, []*types.Block{})
		return
	}

	block, err := h.client.BlockByNumber(r.Context(), header.Number)
	if err != nil {
		// Return empty transactions list if there's an error
		renderTransactionsList(w, r, []*types.Block{})
		return
	}

	// Render just the transactions list component
	renderTransactionsList(w, r, []*types.Block{block})
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
