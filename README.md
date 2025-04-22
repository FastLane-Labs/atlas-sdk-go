# Atlas Go SDK

## Quickstart

Download the SDK package.

```bash
go get github.com/FastLane-Labs/atlas-sdk-go
```

Import the package and initialize the SDK.

```go
package main

import (
	"github.com/FastLane-Labs/atlas-sdk-go/core"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ethClient, _ := ethclient.Dial("https://rpc.sepolia.org")
	defer ethClient.Close()

	sdk, _ := core.NewAtlasSdk([]interface{}{ethClient}, nil)
}
```

## References

The full API reference is described in this document: https://atlas-docs.pages.dev/atlas/sdks/golang/methods.
