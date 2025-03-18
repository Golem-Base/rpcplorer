// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.833
package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/ethereum/go-ethereum/core/types"
	"strconv"
)

// Helper function to shorten hash
func shortenHash(hash string) string {
	if len(hash) <= 12 {
		return hash
	}
	return hash[:10] + "..." + hash[len(hash)-4:]
}

func Index(blocks []*types.Block) templ.Component {
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
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"container mx-auto px-4 py-8\"><!-- Hero Section --><div class=\"bg-white rounded-lg shadow-md p-6 mb-8\"><div class=\"mb-6\"><h2 class=\"text-2xl font-semibold text-gray-800 mb-2\">Welcome to RPCPlorer</h2><p class=\"text-gray-600\">Explore blockchain data with ease.</p></div><div class=\"md:hidden mb-6\"><div class=\"relative\"><input type=\"text\" placeholder=\"Search by Address / Txn Hash / Block / Token\" class=\"w-full py-3 px-4 pr-10 rounded-lg border border-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-500\"> <button class=\"absolute right-3 top-1/2 transform -translate-y-1/2 text-gray-400\"><i class=\"fas fa-search\"></i></button></div></div></div><!-- Latest Transactions and Blocks Section --><div class=\"grid grid-cols-1 md:grid-cols-2 gap-8 mb-8\"><!-- Latest Blocks --><div class=\"bg-white rounded-lg shadow-md\"><div class=\"border-b px-6 py-4 flex justify-between items-center\"><h3 class=\"text-lg font-medium text-gray-800\">Latest Blocks</h3><div class=\"flex items-center\"><div class=\"text-sm text-green-600 mr-3\"><span class=\"htmx-indicator\" id=\"block-poll-indicator\"><i class=\"fas fa-circle-notch fa-spin mr-1\"></i> Auto-refreshing</span></div><button class=\"text-blue-500 hover:text-blue-700 text-sm flex items-center\" hx-get=\"/api/blocks\" hx-target=\"#latest-blocks\" hx-trigger=\"click, every 2s\" hx-indicator=\"#block-spinner\"><i class=\"fas fa-sync-alt mr-1\"></i> Refresh <span id=\"block-spinner\" class=\"htmx-indicator ml-1\"><i class=\"fas fa-circle-notch fa-spin\"></i></span></button></div></div><div id=\"latest-blocks\" class=\"divide-y fade-me-in\" hx-get=\"/api/blocks\" hx-trigger=\"load, every 2s\" hx-indicator=\"#block-poll-indicator\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = BlocksList(blocks).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div></div><!-- Latest Transactions --><div class=\"bg-white rounded-lg shadow-md\"><div class=\"border-b px-6 py-4 flex justify-between items-center\"><h3 class=\"text-lg font-medium text-gray-800\">Latest Transactions</h3><div class=\"flex items-center\"><div class=\"text-sm text-green-600 mr-3\"><span class=\"htmx-indicator\" id=\"txn-poll-indicator\"><i class=\"fas fa-circle-notch fa-spin mr-1\"></i> Auto-refreshing</span></div><button class=\"text-blue-500 hover:text-blue-700 text-sm flex items-center\" hx-get=\"/api/transactions\" hx-target=\"#latest-transactions\" hx-trigger=\"click, every 2s\" hx-indicator=\"#txn-spinner\"><i class=\"fas fa-sync-alt mr-1\"></i> Refresh <span id=\"txn-spinner\" class=\"htmx-indicator ml-1\"><i class=\"fas fa-circle-notch fa-spin\"></i></span></button></div></div><div id=\"latest-transactions\" class=\"divide-y fade-me-in\" hx-get=\"/api/transactions\" hx-trigger=\"load, every 2s\" hx-indicator=\"#txn-poll-indicator\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = TransactionsList(blocks).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "</div></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return nil
		})
		templ_7745c5c3_Err = Layout("RPCPlorer - Home").Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func BlocksList(blocks []*types.Block) templ.Component {
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
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if len(blocks) > 0 {
			for i, block := range blocks {
				if i < 5 && block != nil {
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<div class=\"px-6 py-3 hover:bg-gray-50 fade-me-in\"><div class=\"flex items-center\"><div class=\"bg-red-100 rounded-full p-2 mr-3\"><i class=\"fas fa-cube text-red-500\"></i></div><div class=\"flex-1\"><p class=\"text-blue-500 font-medium\"><a href=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var4 templ.SafeURL = templ.SafeURL("/block/" + strconv.FormatUint(block.NumberU64(), 10))
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var4)))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var5 string
					templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.FormatUint(block.NumberU64(), 10))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/index.templ`, Line: 112, Col: 52}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "</a></p><p class=\"text-sm text-gray-500\">Hash: <a href=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var6 templ.SafeURL = templ.SafeURL("/block/" + block.Hash().Hex())
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var6)))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var7 string
					templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(shortenHash(block.Hash().Hex()))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/index.templ`, Line: 116, Col: 105}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "</a></p></div><div class=\"text-right\"><p class=\"text-gray-800\">Miner: <a href=\"")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var8 templ.SafeURL = templ.SafeURL("/address/" + block.Coinbase().Hex())
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var8)))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "\" class=\"text-blue-500\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var9 string
					templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(shortenHash(block.Coinbase().Hex()))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/index.templ`, Line: 120, Col: 162}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</a></p><p class=\"text-sm text-gray-500\">")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					var templ_7745c5c3_Var10 string
					templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(len(block.Transactions())))
					if templ_7745c5c3_Err != nil {
						return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/index.templ`, Line: 121, Col: 81}
					}
					_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
					templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, " txns</p></div></div></div>")
					if templ_7745c5c3_Err != nil {
						return templ_7745c5c3_Err
					}
				}
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, " ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if len(blocks) == 0 {
				templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "<div class=\"px-6 py-4 text-center text-gray-500\">No blocks available</div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		} else {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<div class=\"px-6 py-4 text-center text-gray-500\">No blocks available</div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 15, "<div class=\"border-t px-6 py-3 text-center\"><a href=\"/blocks\" class=\"text-blue-500 hover:text-blue-700\">View all blocks</a></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
