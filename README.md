# payment-go-sdk

A Golang SDK for monetha payment contracts.


### Contributing 
	
#### Build

Install dependencies:

    make dependencies

#### Making changes

Make your changes, then ensure that the linters pass:

    make lint	

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

### Setup

```golang
var (
    rawRPCURL = "https://www.etherscan.io" // Can you testnet as well (https://www.ropsten.etherscan.io)
    key = "your processing key"
)
paymentHandler := paymenthandler.New(rawRPCURL , key)
```
### Payment operations

#### AddOrder

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
// Use txHash for further operations.
```

#### CancelOrder

```golang
txHash, err := paymentHandler.CancelOrder(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int,
	cancelReason string)
if err != nil {
    // error handling
}
// Use txHash for further operations.
```

#### SecurePay

```golang
txHash, err := paymentHandler.SecurePay(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	walletKey *ecdsa.PrivateKey)
if err != nil {
    // error handling
}
// Use txHash for further operations.
```

#### SecureTokenPay

```golang
txHash, err := paymentHandler.SecureTokenPay(ctx context.Context,
	contractAddress, tokenAddress common.Address,
	gasPrice *big.Int,
	gasLimit *big.Int,
	orderID *big.Int,
	walletKey *ecdsa.PrivateKey)
if err != nil {
    // error handling
}
// Use txHash for further operations.
```

#### ProcessPayment

```golang
txHash, err := paymentHandler.ProcessPayment(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int)
if err != nil {
    // error handling
}
// Use txHash for further operations.
```

#### WithdrawRefund

```golang
txHash, err := paymentHandler.WithdrawRefund(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int)
if err != nil {
    // error handling
}
// Use txHash for further operations.
```

#### WithdrawTokenRefund

```golang
txHash, err := paymentHandler.WithdrawTokenRefund(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int)
if err != nil {
    // error handling
}
// Use txHash for further operations.
```

### Merchant Wallet operations

#### WithdrawToExchange

```golang
txHash, err := paymentHandler.WithdrawToExchange(ctx context.Context,
	walletAddress common.Address,
	depositAddress common.Address,
	amount *big.Int)
if err != nil {
    // error handling
}
// Use txHash for further operations.
```

#### WithdrawAllToExchange

```golang
txHash, err := paymentHandler.WithdrawAllToExchange(ctx context.Context,
	walletAddress common.Address,
	depositAddress common.Address,
	minAmount *big.Int)
if err != nil {
    // error handling
}
// Use txHash for further operations.
```

#### WithdrawAllTokensToExchange

```golang
txHash, err := paymentHandler.WithdrawAllTokensToExchange(ctx context.Context,
	tokenAddress common.Address,
	walletAddress common.Address,
	depositAddress common.Address,
	minAmount *big.Int)
if err != nil {
    // error handling
}
// Use txHash for further operations.
```


