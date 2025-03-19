// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strconv"
	"time"
)

// Helper function to get transaction type name
func getTxTypeName(txType uint8) string {
	switch txType {
	case 0:
		return "Legacy"
	case 1:
		return "Access List"
	case 2:
		return "EIP-1559"
	default:
		return "Unknown"
	}
}

var topicToEventName = map[common.Hash]string{
	common.HexToHash("0xce4b4ad6891d716d0b1fba2b4aeb05ec20edadb01df512263d0dde423736bbb9"): "GolemBase Entity Created",
	common.HexToHash("0xf371f40aa6932ad9dacbee236e5f3b93d478afe3934b5cfec5ea0d800a41d165"): "GolemBase Entity Deleted",
	common.HexToHash("0x0297b0e6eaf1bc2289906a8123b8ff5b19e568a60d002d47df44f8294422af93"): "GolemBase Entity Updated",
}

// Helper function to identify and format Golem Base storage events
func getGolemBaseEventName(topic common.Hash) string {
	return topicToEventName[topic]
}

// Helper function to format Golem Base event data
func formatGolemBaseEventData(eventName string, topics []common.Hash, data []byte) string {
	switch eventName {
	case "GolemBase Entity Created", "GolemBase Entity Updated":
		entityKey := topics[1].Hex()
		expiresAt := new(big.Int).SetBytes(data).String()
		return fmt.Sprintf("Entity Key: %s\nExpires at block: %s", entityKey, expiresAt)
	case "GolemBase Entity Deleted":
		entityKey := topics[1].Hex()
		return fmt.Sprintf("Entity Key: %s", entityKey)
	default:
		return common.Bytes2Hex(data)
	}
}

func Transaction(tx *types.Transaction, receipt *types.Receipt, block *types.Block, currentGasPrice *big.Int, from common.Address) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"container mx-auto px-4 py-8\"><!-- Transaction Summary Card --><div class=\"bg-white rounded-lg shadow-md p-6 mb-8\"><h1 class=\"text-2xl font-semibold text-gray-800 mb-4\">Transaction Details</h1><div class=\"mb-4\"><div class=\"p-3 bg-blue-50 rounded-md flex items-center\"><div class=\"mr-3 bg-blue-100 rounded-full p-2\"><i class=\"fas fa-exchange-alt text-blue-500\"></i></div><span class=\"text-gray-600 break-all\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(tx.Hash().Hex())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 64, Col: 61}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</span> <button class=\"ml-2 text-blue-500 hover:text-blue-700\" data-value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(tx.Hash().Hex())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 67, Col: 35}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "\" onclick=\"copyToClipboard(this)\"><i class=\"fas fa-copy\"></i></button></div></div><div class=\"grid grid-cols-1 md:grid-cols-2 gap-6\"><!-- Status --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Status:</div><div class=\"flex items-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if receipt.Status == 1 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<span class=\"bg-green-100 text-green-800 text-sm font-medium mr-2 px-2.5 py-0.5 rounded\">Success</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "<span class=\"bg-red-100 text-red-800 text-sm font-medium mr-2 px-2.5 py-0.5 rounded\">Failed</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</div></div><!-- Block --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Block:</div><div class=\"flex items-center\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 templ.SafeURL = templ.SafeURL("/block/" + block.Number().String())
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var5)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\" class=\"text-blue-500 hover:text-blue-700\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(block.Number().String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 91, Col: 137}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</a> <span class=\"text-gray-500 ml-2\">(")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(block.Transactions().Len()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 92, Col: 83}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, " txns)</span></div></div><!-- Timestamp --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Timestamp:</div><div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var8 string
			templ_7745c5c3_Var8, templ_7745c5c3_Err = templ.JoinStringErrs(time.Unix(int64(block.Time()), 0).Format(time.RFC1123))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 100, Col: 63}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var8))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</div></div><!-- From --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">From:</div><div class=\"flex items-center\"><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var9 templ.SafeURL = templ.SafeURL("/address/" + from.Hex())
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var9)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "\" class=\"text-blue-500 hover:text-blue-700 break-all\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var10 string
			templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(from.Hex())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 109, Col: 20}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "</a> <button class=\"ml-2 text-blue-500 hover:text-blue-700\" data-value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var11 string
			templ_7745c5c3_Var11, templ_7745c5c3_Err = templ.JoinStringErrs(from.Hex())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 113, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var11))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\" onclick=\"copyToClipboard(this)\"><i class=\"fas fa-copy\"></i></button></div></div><!-- To --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">To:</div><div class=\"flex items-center\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if tx.To() != nil {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<a href=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var12 templ.SafeURL = templ.SafeURL("/address/" + tx.To().Hex())
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var12)))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "\" class=\"text-blue-500 hover:text-blue-700 break-all\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var13 string
				templ_7745c5c3_Var13, templ_7745c5c3_Err = templ.JoinStringErrs(tx.To().Hex())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 126, Col: 24}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var13))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 16, "</a> <button class=\"ml-2 text-blue-500 hover:text-blue-700\" data-value=\"")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var14 string
				templ_7745c5c3_Var14, templ_7745c5c3_Err = templ.JoinStringErrs(tx.To().Hex())
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 130, Col: 35}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var14))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 17, "\" onclick=\"copyToClipboard(this)\"><i class=\"fas fa-copy\"></i></button>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 18, "<span class=\"italic text-gray-500\">Contract Creation</span>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 19, "</div></div><!-- Value --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Value:</div><div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var15 string
			templ_7745c5c3_Var15, templ_7745c5c3_Err = templ.JoinStringErrs(tx.Value().String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 144, Col: 28}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var15))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 20, " ETH</div></div><!-- Transaction Fee --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Transaction Fee:</div><div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var16 string
			templ_7745c5c3_Var16, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%.8f", float64(receipt.GasUsed)*float64(tx.GasPrice().Uint64())/1e18))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 152, Col: 95}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var16))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 21, " ETH</div></div><!-- Gas Price --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Gas Price:</div><div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var17 string
			templ_7745c5c3_Var17, templ_7745c5c3_Err = templ.JoinStringErrs(tx.GasPrice().String())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 160, Col: 31}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var17))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 22, " Wei (")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var18 string
			templ_7745c5c3_Var18, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%.2f", float64(tx.GasPrice().Uint64())/1e9))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 160, Col: 99}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var18))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 23, " Gwei)</div></div><!-- Gas Limit & Usage --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Gas Limit & Usage:</div><div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var19 string
			templ_7745c5c3_Var19, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d | %d (%0.2f%%)", tx.Gas(), receipt.GasUsed, float64(receipt.GasUsed)/float64(tx.Gas())*100))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 168, Col: 120}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var19))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 24, "</div></div><!-- Nonce --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Nonce:</div><div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var20 string
			templ_7745c5c3_Var20, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", tx.Nonce()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 176, Col: 38}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var20))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 25, "</div></div><!-- Transaction Type --><div class=\"border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Type:</div><div>0x")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var21 string
			templ_7745c5c3_Var21, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%x", tx.Type()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 184, Col: 39}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var21))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 26, " (")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var22 string
			templ_7745c5c3_Var22, templ_7745c5c3_Err = templ.JoinStringErrs(getTxTypeName(tx.Type()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 184, Col: 69}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var22))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 27, ")</div></div><!-- Input Data --><div class=\"col-span-1 md:col-span-2 border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Input Data:</div><div class=\"bg-gray-50 p-3 rounded text-xs font-mono overflow-x-auto break-all\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var23 string
			templ_7745c5c3_Var23, templ_7745c5c3_Err = templ.JoinStringErrs(common.Bytes2Hex(tx.Data()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 192, Col: 36}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var23))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 28, "</div></div><!-- Transaction Logs --><div class=\"col-span-1 md:col-span-2 border-b pb-3\"><div class=\"text-gray-500 text-sm mb-1\">Transaction Logs:</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(receipt.Logs) > 0 {
				for i, log := range receipt.Logs {
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 29, "<div class=\"bg-gray-50 p-4 rounded mb-3\"><div class=\"flex items-center justify-between mb-2\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					if eventName := getGolemBaseEventName(log.Topics[0]); eventName != "" {
						templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 30, "<span class=\"text-sm font-medium text-blue-700\">")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var24 string
						templ_7745c5c3_Var24, templ_7745c5c3_Err = templ.JoinStringErrs(eventName)
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 204, Col: 70}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var24))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 31, "</span> ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					} else {
						templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 32, "<span class=\"text-sm font-medium text-gray-700\">Log #")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						var templ_7745c5c3_Var25 string
						templ_7745c5c3_Var25, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(i + 1))
						if templ_7745c5c3_Err != nil {
							return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 206, Col: 85}
						}
						_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var25))
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 33, "</span> ")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 34, "<span class=\"text-xs text-gray-500\">Address: <a href=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var26 templ.SafeURL = templ.SafeURL("/address/" + log.Address.Hex())
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var26)))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 35, "\" class=\"text-blue-500 hover:text-blue-700\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var27 string
					templ_7745c5c3_Var27, templ_7745c5c3_Err = templ.JoinStringErrs(log.Address.Hex())
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 208, Col: 175}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var27))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 36, "</a></span></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					if len(log.Topics) > 0 {
						templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 37, "<div class=\"mb-2\"><div class=\"text-xs text-gray-500 mb-1\">Topics:</div>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						for j, topic := range log.Topics {
							if j == 0 {
								if eventName := getGolemBaseEventName(topic); eventName != "" {
									templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 38, "<div class=\"text-xs font-mono break-all\"><span class=\"text-blue-600\">Event Signature:</span> ")
									if templ_7745c5c3_Err != nil {
										return templ_7745c5c3_Err
									}
									var templ_7745c5c3_Var28 string
									templ_7745c5c3_Var28, templ_7745c5c3_Err = templ.JoinStringErrs(topic.Hex())
									if templ_7745c5c3_Err != nil {
										return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 217, Col: 80}
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var28))
									if templ_7745c5c3_Err != nil {
										return templ_7745c5c3_Err
									}
									templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 39, "</div>")
									if templ_7745c5c3_Err != nil {
										return templ_7745c5c3_Err
									}
								} else {
									templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 40, "<div class=\"text-xs font-mono break-all\">")
									if templ_7745c5c3_Err != nil {
										return templ_7745c5c3_Err
									}
									var templ_7745c5c3_Var29 string
									templ_7745c5c3_Var29, templ_7745c5c3_Err = templ.JoinStringErrs(topic.Hex())
									if templ_7745c5c3_Err != nil {
										return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 220, Col: 68}
									}
									_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var29))
									if templ_7745c5c3_Err != nil {
										return templ_7745c5c3_Err
									}
									templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 41, "</div>")
									if templ_7745c5c3_Err != nil {
										return templ_7745c5c3_Err
									}
								}
							} else {
								templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 42, "<div class=\"text-xs font-mono break-all\">")
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
								var templ_7745c5c3_Var30 string
								templ_7745c5c3_Var30, templ_7745c5c3_Err = templ.JoinStringErrs(topic.Hex())
								if templ_7745c5c3_Err != nil {
									return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 223, Col: 67}
								}
								_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var30))
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
								templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 43, "</div>")
								if templ_7745c5c3_Err != nil {
									return templ_7745c5c3_Err
								}
							}
						}
						templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 44, "</div>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					if len(log.Data) > 0 {
						templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 45, "<div><div class=\"text-xs text-gray-500 mb-1\">Data:</div>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
						if eventName := getGolemBaseEventName(log.Topics[0]); eventName != "" {
							templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 46, "<div class=\"text-xs font-mono whitespace-pre-wrap\">")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var31 string
							templ_7745c5c3_Var31, templ_7745c5c3_Err = templ.JoinStringErrs(formatGolemBaseEventData(eventName, log.Topics, log.Data))
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 232, Col: 122}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var31))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 47, "</div>")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						} else {
							templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 48, "<div class=\"text-xs font-mono break-all\">")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							var templ_7745c5c3_Var32 string
							templ_7745c5c3_Var32, templ_7745c5c3_Err = templ.JoinStringErrs(common.Bytes2Hex(log.Data))
							if templ_7745c5c3_Err != nil {
								return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/transaction.templ`, Line: 234, Col: 81}
							}
							_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var32))
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
							templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 49, "</div>")
							if templ_7745c5c3_Err != nil {
								return templ_7745c5c3_Err
							}
						}
						templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 50, "</div>")
						if templ_7745c5c3_Err != nil {
							return templ_7745c5c3_Err
						}
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 51, "</div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			} else {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 52, "<div class=\"text-gray-500 text-sm\">No logs for this transaction</div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 53, "</div></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Layout("Transaction Details | RPCPlorer").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
