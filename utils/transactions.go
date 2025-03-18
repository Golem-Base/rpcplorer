package utils

import (
	"github.com/ethereum/go-ethereum/core/types"
)

// TransactionWithBlock stores a transaction with its associated block
type TransactionWithBlock struct {
	Tx    *types.Transaction
	Block *types.Block
}

// ExtractRecentTransactions extracts transactions from multiple blocks
// and returns them in reverse chronological order, with a maximum limit
func ExtractRecentTransactions(blocks []*types.Block, limit int) []TransactionWithBlock {
	var result []TransactionWithBlock

	// Extract transactions from all blocks
	for _, block := range blocks {
		if block != nil {
			for _, tx := range block.Transactions() {
				result = append(result, TransactionWithBlock{
					Tx:    tx,
					Block: block,
				})

				// Limit the number of transactions
				if len(result) >= limit {
					return result[:limit]
				}
			}
		}
	}

	return result
}
