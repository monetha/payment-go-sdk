package contracts

//go:generate go get github.com/ethereum/go-ethereum/cmd/abigen
//go:generate abigen --abi ./code/GenericERC20.abi --out GenericERC20.go --pkg contracts --type GenericERC20
//go:generate abigen --abi ./code/MonethaGateway.abi --bin ./code/MonethaGateway.bin --out MonethaGateway.go --pkg contracts --type MonethaGatewayContract
//go:generate abigen --abi ./code/MerchantWallet.abi --bin ./code/MerchantWallet.bin --out MerchantWallet.go --pkg contracts --type MerchantWalletContract
//go:generate abigen --abi ./code/MerchantDealsHistory.abi --bin ./code/MerchantDealsHistory.bin --out MerchantDealsHistory.go --pkg contracts --type MerchantDealsHistoryContract
//go:generate abigen --abi ./code/PaymentProcessor.abi --bin ./code/PaymentProcessor.bin --out PaymentProcessor.go --pkg contracts --type PaymentProcessorContract
//go:generate abigen --abi ./code/PrivatePaymentProcessor.abi --bin ./code/PrivatePaymentProcessor.bin --out PrivatePaymentProcessor.go --pkg contracts --type PrivatePaymentProcessorContract
//go:generate abigen --abi ./code/MonethaTokenHoldersProgram.abi --bin ./code/MonethaTokenHoldersProgram.bin --out MonethaTokenHoldersProgram.go --pkg contracts --type MonethaTokenHoldersProgramContract
//go:generate abigen --abi ./code/MonethaVoucher.abi --bin ./code/MonethaVoucher.bin --out MonethaVoucher.go --pkg contracts --type MonethaVoucherContract
//go:generate abigen --abi ./code/MonethaToken.abi --bin ./code/MonethaToken.bin --out MonethaToken.go --pkg contracts --type MonethaTokenContract
//go:generate abigen --abi ./code/MonethaSupportedTokens.abi --bin ./code/MonethaSupportedTokens.bin --out MonethaSupportedTokens.go --pkg contracts --type MonethaSupportedTokensContract
