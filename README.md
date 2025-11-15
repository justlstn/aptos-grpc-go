# Aptos Go gRPC client

[![Generate client](https://github.com/justlstn/aptos-grpc-go/actions/workflows/main.yaml/badge.svg)](https://github.com/justlstn/aptos-grpc-go/actions/workflows/main.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/justlstn/aptos-grpc-go.svg)](https://pkg.go.dev/github.com/justlstn/aptos-grpc-go)

A community-based gRPC Golang client generator for Aptos blockchain protocol buffers.

## ⚠️ Important Notice

**This is NOT an official Aptos client.** This project is a community-driven initiative to provide Golang gRPC clients for the Aptos blockchain.

## Legal Notice

- All protocol buffer definitions (`*.proto` files) are the intellectual property of Aptos Labs and are subject to their original licensing terms
- The original proto files are licensed under the Apache License 2.0 by Aptos Foundation
- The proto files and related specifications remain the property of Aptos Labs
- This project only provides tooling and generated Go code based on the publicly available proto definitions
- Generated Go code in this repository is licensed under the MIT License (see LICENSE file)

### Apache 2.0 License Compliance

This project generates derivative works from Apache 2.0 licensed proto files. In compliance with the Apache 2.0 license terms:

- **Attribution**: Original proto files are Copyright © Aptos Foundation, licensed under Apache 2.0
- **Notice**: The original proto files can be found at: https://github.com/aptos-labs/aptos-core
- **Derivative Work**: The generated Go code constitutes a derivative work of the original proto files
- **License Grant**: The Apache 2.0 license grants permission to create and distribute derivative works

## What This Project Provides

This project generates Go gRPC client code from the official Aptos Core protocol buffer definitions, enabling Go developers to interact with the Aptos blockchain through strongly-typed client libraries.

## Installation

```bash
go get github.com/justlstn/aptos-grpc-go
```

## Usage

### Available gRPC Services

This package provides the following gRPC client interfaces:

- **RawDataClient** - Stream transactions from the Aptos indexer
- **GrpcManagerClient** - Manage indexer connections and requests
- **DataServiceClient** - Access indexer data services
- **FullnodeDataClient** - Connect to Aptos fullnode for transaction data
- **NetworkMessageServiceClient** - Handle network messaging

### Basic Example: Streaming Transactions

```go
package main

import (
    "context"
    "crypto/tls"
    "fmt"
    "io"
    "log"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    indexerv1 "github.com/justlstn/aptos-grpc-go/aptos/indexer/v1"
)

func main() {
    // Connect to Aptos indexer gRPC endpoint
    creds := credentials.NewTLS(&tls.Config{})
    conn, err := grpc.Dial(
        "grpc.mainnet.aptoslabs.com:443",
        grpc.WithTransportCredentials(creds),
    )
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Create RawData client
    client := indexerv1.NewRawDataClient(conn)

    // Request transactions starting from version 0
    req := &indexerv1.GetTransactionsRequest{
        StartingVersion: 0,
    }

    // Stream transactions
    stream, err := client.GetTransactions(context.Background(), req)
    if err != nil {
        log.Fatalf("Failed to get transactions: %v", err)
    }

    // Process streamed transactions
    for {
        resp, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("Error receiving transaction: %v", err)
        }

        fmt.Printf("Received %d transactions\n", len(resp.Transactions))
        for _, tx := range resp.Transactions {
            fmt.Printf("  Version: %d, Type: %v\n", tx.Version, tx.Type)
        }
    }
}
```

### Example: Using Fullnode Client

```go
package main

import (
    "context"
    "crypto/tls"
    "fmt"
    "log"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    fullnodev1 "github.com/justlstn/aptos-grpc-go/aptos/internal/fullnode/v1"
)

func main() {
    // Connect to Aptos fullnode
    creds := credentials.NewTLS(&tls.Config{})
    conn, err := grpc.Dial(
        "fullnode.mainnet.aptoslabs.com:443",
        grpc.WithTransportCredentials(creds),
    )
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Create Fullnode client
    client := fullnodev1.NewFullnodeDataClient(conn)

    // Ping the fullnode
    pingResp, err := client.Ping(context.Background(), &fullnodev1.PingFullnodeRequest{})
    if err != nil {
        log.Fatalf("Ping failed: %v", err)
    }

    fmt.Printf("Fullnode ping successful: %+v\n", pingResp)
}
```

### Testnet Example

```go
// For testnet, use:
conn, err := grpc.Dial(
    "grpc.testnet.aptoslabs.com:443",
    grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
)
```

## Automatic Updates

This repository is automatically updated when new proto definitions are published by Aptos Labs:

- Runs 3 times daily (6 AM, 12 PM, 6 PM UTC) on weekdays
- Checks for upstream changes in `aptos-labs/aptos-core`
- Automatically regenerates Go code when proto files change
- Creates pull requests for review before merging to main

You can view the latest sync status in the `.sync-state` file or check the [workflow runs](https://github.com/justlstn/aptos-grpc-go/actions).

## Contributing

Contributions are welcome! Please feel free to submit issues and pull requests.

## Disclaimer

This project is provided "as is" without warranty of any kind. Use at your own risk. The maintainers are not responsible for any issues that may arise from using this code.

## License

- **This project's code and tooling**: MIT License (see LICENSE file)
- **Original Aptos proto files (not included in the repository, but used to generate the code)**: Apache License 2.0 (Copyright © Aptos Foundation)
- **Generated Go code**: MIT License (derivative work of Apache 2.0 licensed proto files)

For the original Aptos Core license, see: https://github.com/aptos-labs/aptos-core/blob/main/LICENSE

### Apache 2.0 License Notice

```
Licensed under the Apache License, Version 2.0 (the "License");
you may not use the original proto files except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

## Related Projects

- [Aptos Core](https://github.com/aptos-labs/aptos-core) - The official Aptos blockchain implementation
- [Proto files](https://github.com/aptos-labs/aptos-core/tree/main/protos) - Original Aptos protocol buffer definitions
- [Aptos Developer Documentation](https://aptos.dev/) - Official developer resources
