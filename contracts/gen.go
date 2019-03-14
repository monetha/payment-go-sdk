package contracts

//go:generate go get github.com/ethereum/go-ethereum/cmd/abigen
//go:generate abigen --abi ./code/MerchantWallet.abi --bin ./code/MerchantWallet.bin --out MerchantWallet.go --pkg contracts --type MerchantWalletContract
//go:generate abigen --abi ./code/PaymentProcessor.abi --bin ./code/PaymentProcessor.bin --out PaymentProcessor.go --pkg contracts --type PaymentProcessorContract
