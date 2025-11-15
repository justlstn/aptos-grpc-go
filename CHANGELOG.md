# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2025-11-15

### Added
- Initial release of Aptos Go gRPC client
- Generated Go gRPC client code from Aptos protocol buffers
- Support for the following gRPC services:
  - `RawDataClient` - Stream transactions from the Aptos indexer
  - `GrpcManagerClient` - Manage indexer connections and requests
  - `DataServiceClient` - Access indexer data services
  - `FullnodeDataClient` - Connect to Aptos fullnode for transaction data
  - `NetworkMessageServiceClient` - Handle network messaging
- Automated GitHub workflow for syncing proto files from upstream
  - Runs 3 times daily (6 AM, 12 PM, 6 PM UTC) on weekdays
  - Automatically detects changes in `aptos-labs/aptos-core`
  - Creates PRs when updates are detected
- Custom patch system for maintaining Go package options
- Comprehensive README with usage examples
- MIT License with Apache 2.0 compliance notice

### Dependencies
- `google.golang.org/grpc` v1.73.0
- `google.golang.org/protobuf` v1.36.6

### Upstream Sync
- Synced with Aptos Core commit: `a61cfce4ab5515008a3a4142b8a462cbf835ed6c`

[0.1.0]: https://github.com/justlstn/aptos-grpc-go/releases/tag/v0.1.0
