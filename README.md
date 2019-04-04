# Monetha: Decentralized Reputation Framework

## Payment Layer: go-sdk 
[![GoDoc](https://godoc.org/github.com/monetha/payment-go-sdk?status.svg)](https://godoc.org/github.com/monetha/payment-go-sdk)

A Golang SDK for decentralized Monetha Payment Gateway usage.

- [Monetha: Decentralized Reputation Framework](#monetha-decentralized-reputation-framework)
	- [Payment Layer: go-sdk](#payment-layer-go-sdk)
	- [Building the source](#building-the-source)
		- [Prerequisites](#prerequisites)
		- [Build](#build)
	- [Contributing](#contributing)
		- [Making changes](#making-changes)
			- [Contracts update](#contracts-update)
			- [Formatting source code](#formatting-source-code)
	- [Usage](#usage)
		- [Installation & setup](#installation--setup)
		- [Payment operations](#payment-operations)
			- [AddOrder](#addorder)
			- [CancelOrder](#cancelorder)
			- [SecurePay](#securepay)
			- [SecureTokenPay](#securetokenpay)
			- [ProcessPayment](#processpayment)
			- [RefundPayment](#refundpayment)
			- [WithdrawRefund](#withdrawrefund)
			- [WithdrawTokenRefund](#withdrawtokenrefund)
		- [Merchant Wallet operations](#merchant-wallet-operations)
			- [Withdraw](#withdraw)
			- [WithdrawTo](#withdrawto)
			- [WithdrawAllTo](#withdrawallto)
			- [WithdrawAllTokensTo](#withdrawalltokensto)

## Building the source

### Prerequisites

1. Make sure you have Git installed.
2. Install Go 1.12
3. Setup $GOPATH environment variable as described here.
4. Clone the repository:

```bash
mkdir -p $GOPATH/src/github.com/monetha
cd $GOPATH/src/github.com/monetha
git clone git@github.com:monetha/payment-go-sdk.git
cd payment-go-sdk
```

**Note:** You can skip steps 2-3 on Linux and use the official docker image for Go after step 4 to build the project:

```bash
docker run -it --rm \
  -v "$PWD":/go/src/github.com/monetha/reputation-go-sdk \
  -w /go/src/github.com/monetha/reputation-go-sdk \
  golang:1.12 \
  /bin/bash
```

### Build

Install dependencies:

```bash
    make dependencies
```

## Contributing

Thank you for considering to help out with the source code! We welcome contributions from anyone on the internet, and are grateful for even the smallest of fixes! If you'd like to contribute to payment-go-sdk, please fork, fix, commit and send a pull request for the maintainers to review and merge into the main code base.  Feel free to register issues and suggestions 

### Making changes

Make your changes, then ensure that the linters pass:

```bash
    make lint
```

#### Contracts update

This SDK is dependent on [Monetha's payment layer smart contracts](https://github.com/monetha/payment-contracts). In case if contracts will be updated and method signatures will change `contracts/MerchantWallet.go` and `contracts/PaymentProcessor.go` files need to be re-generated. To do so follow the steps below:

1. Copy all necessary smart contract build artifacts of [github.com/monetha/payment-contracts](https://github.com/monetha/payment-contracts) to `contracts/code` folder.
2. Run `go generate` command in `contracts` folder to convert Ethereum contracts into Go package.
3. Commit new/updated files.

#### Formatting source code

`make fmt` command automatically formats Go source code of the entire project.

## Usage

In order to better understand the use cases of the `payment-go-sdk` please refer to [Monetha Payment layer: de-centralized  usage scenarios](https://github.com/monetha/payment-layer#de-centralized-payment-layer-usage-scenarios)

### Installation & setup

Install new dependency from you project's folder

```bash
go get github.com/monetha/payment-go-sdk
```

After package is installed import the package in your working file

```golang
import (
    "github.com/monetha/payment-go-sdk"
)
```

Initialize paymentHandler by providing 2 parameters

```golang
paymentHandler := paymenthandler.New(EthereumRPCURL , PrivateKey)
```

- `EthereumRPCURL` - Ethereum network node RPC url 
- `PrivateKey` - user's private key who will be covering operational costs of transaction execution (Customer or Merchant).

### Payment operations

#### AddOrder

This method can be executed by either Merchant or a Customer. This method initializes the deal between 2 parties Merchant and Customer.

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

Parameters:

- **ctx**
- **contractAddresses**
- **gasPrice**
- **orderID**
- **price**
- **paymentAcceptor**
- **originAddress**
- **fee**
- **tokenAddress**
- **vouchersApplied**

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

Parameters:

- **ctx**
- **contractAddresses**
- **gasPrice**
- **orderID**
- **clientReputation**
- **merchantReputation**
- **dealHash**
- **cancelReason**

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

Parameters:

- **ctx**
- **contractAddresses**
- **gasPrice**
- **orderID**
- **walletKey**

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

Parameters:

- **ctx**
- **contractAddresses**
- **gasPrice**
- **orderID**
- **walletKey**

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

Parameters:

- **ctx**
- **contractAddresses**
- **gasPrice**
- **orderID**
- **clientReputation**
- **merchantReputation**
- **dealHash**

#### RefundPayment

TBD

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

Parameters:

- **ctx**
- **contractAddresses**
- **gasPrice**
- **orderID**

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

Parameters:

- **ctx**
- **contractAddresses**
- **gasPrice**
- **orderID**

### Merchant Wallet operations

MerchantWallet contract operations allowing a Merchant to manage collected funds

#### Withdraw

TBD

#### WithdrawTo

TBD

#### WithdrawAllTo

TBD


#### WithdrawAllTokensTo

TBD
