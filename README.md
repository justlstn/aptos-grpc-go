# Aptos Golang Protos

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
