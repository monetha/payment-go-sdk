# payment-go-sdk

A Golang SDK for monetha payment contracts.

### Installation

```shell
go get github.com/monetha/payment-go-sdk
```

### Importing

```golang
import (
    "github.com/monetha/payment-go-sdk"
)
```

#### Setup

Init client for API services. Get APIKey/SecretKey from your binance account.

```golang
var (
    rawRPCURL = "https://www.etherscan.io" // Can you testnet as well (https://www.ropsten.etherscan.io)
    key = "your processing key"
)
paymentHandler := paymenthandler.New(rawRPCURL , key)
```

#### Add order

```golang
txHash, err := paymentHandler.AddOrder(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	price *big.Int,
	paymentAcceptor common.Address,
	originAddress common.Address,
	fee *big.Int,
	tokenAddress common.Address,
    vouchersApplied *big.Int)
if err != nil {
    // error handling
}
// Use txHash for confirmation if you want.
```

#### Cancel order

```golang
txHash, err := paymentHandler.CancelOrder(
	ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int,
	cancelReason string,
)
if err != nil {
    // error handling
}
// Use txHash for confirmation if you want.
```