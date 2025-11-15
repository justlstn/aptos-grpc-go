# Stream Transactions Example

This example demonstrates how to stream transactions from the Aptos indexer using the gRPC client.

## Prerequisites

1. Get an API key from [geomi.dev](https://geomi.dev):
   - Sign in to geomi.dev
   - Navigate to "API Resource"
   - Create a new API key

2. Set the API key as an environment variable:
   ```bash
   export APTOS_API_KEY="your-api-key-here"
   ```

## Running the Example

```bash
cd examples/stream_transactions
go run main.go
```

Or build and run:

```bash
go build -o stream_transactions .
./stream_transactions
```

## Expected Output

```
Received 100 transactions
  Version: 0, Type: TRANSACTION_TYPE_GENESIS
  Version: 1, Type: TRANSACTION_TYPE_STATE_CHECKPOINT
  Version: 2, Type: TRANSACTION_TYPE_USER
  ...
```

## Customization

To connect to a different network, change the endpoint in `main.go`:

- **Testnet**: `grpc.testnet.aptoslabs.com:443`
- **Devnet**: `grpc.devnet.aptoslabs.com:443`
