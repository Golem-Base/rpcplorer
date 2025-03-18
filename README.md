# RPCPlorer

A lightweight Ethereum blockchain explorer that connects directly to any Ethereum JSON-RPC node. Built with Go, HTMX, and Tailwind CSS.

## Features

- **Block Explorer**: View detailed information about blocks, including transactions, gas usage, miner data, and more.
- **Transaction Explorer**: Examine transaction details, including status, gas costs, input data, and involved addresses.
- **Real-time Updates**: Auto-refreshing display of latest blocks and transactions.
- **Direct Node Connection**: Works with any Ethereum-compatible JSON-RPC endpoint.
- **Clean, Modern UI**: Responsive design with Tailwind CSS.
- **Fast Performance**: Server-side rendering with HTMX for smooth interactions.
- **Pagination**: Browse through blockchain history with efficient pagination controls.

## Installation

### Prerequisites

- Go 1.24 or later
- [Templ](https://github.com/a-h/templ) for templating

### Steps

1. Clone the repository:
   ```
   git clone https://github.com/Golem-Base/rpcplorer.git
   cd rpcplorer
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Generate template code:
   ```
   templ generate
   ```

4. Build the application:
   ```
   go build
   ```

## Usage

1. Run RPCPlorer with a valid Ethereum node URL:
   ```
   ./rpcplorer --node-url="https://your-ethereum-node-url"
   ```

   Or using an environment variable:
   ```
   NODE_URL="https://your-ethereum-node-url" ./rpcplorer
   ```

2. Open a web browser and navigate to `http://localhost:8080`

## Available Pages

- **Home**: `/` - Overview of latest blocks and transactions
- **Blocks History**: `/blocks` - Paginated list of blocks (50 per page) with age, transaction count, gas usage, and miner info
- **Block Details**: `/block/{number or hash}` - Detailed view of a specific block
- **Transaction Details**: `/tx/{hash}` - Detailed view of a specific transaction

## Development

### Structure

- `/handlers` - HTTP handlers for processing requests
- `/templates` - Templ templates for rendering HTML
- `main.go` - Application entry point

### Rebuild Templates

After modifying `.templ` files, regenerate the Go code:

```
templ generate
```

## Dependencies

- [go-ethereum](https://github.com/ethereum/go-ethereum) - Ethereum Go library
- [templ](https://github.com/a-h/templ) - Type-safe HTML templating
- [HTMX](https://htmx.org/) - Frontend interactivity
- [Tailwind CSS](https://tailwindcss.com/) - Styling
- [urfave/cli](https://github.com/urfave/cli) - Command-line interface

