package paymenthandler

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/payment-go-sdk/contracts"
)

// PayForOrder calls payForOrder method that transfers money to merchant and Monetha vault. It's part of cheap flow.
func (ch *PaymentHandler) PayForOrder(ctx context.Context,
	contractAddresses common.Address,
	walletKey *ecdsa.PrivateKey,
	orderID *big.Int,
	gasPrice *big.Int,
	gasLimit *big.Int,
	originAddress common.Address,
	price *big.Int,
	fee *big.Int,
	vouchersApplied *big.Int) (txHash string, err error) {
	privatePaymentProcessorContract, err := contracts.NewPrivatePaymentProcessorContract(contractAddresses, ch.backend)
	if err != nil {
		return
	}

	walletAuth := bind.NewKeyedTransactor(walletKey)
	walletAuth.Context = ctx
	walletAuth.Value = price
	walletAuth.GasPrice = gasPrice
	walletAuth.GasLimit = gasLimit.Uint64()

	tx, err := privatePaymentProcessorContract.PayForOrder(walletAuth, orderID, originAddress, fee, vouchersApplied)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// PayForOrderInTokens calls payForOrderInTokens method that transfers money to merchant and Monetha vault. It's part of cheap flow.
func (ch *PaymentHandler) PayForOrderInTokens(ctx context.Context,
	contractAddresses common.Address,
	walletKey *ecdsa.PrivateKey,
	gasPrice *big.Int,
	gasLimit *big.Int,
	orderID *big.Int,
	originAddress common.Address,
	tokenAddress common.Address,
	orderValue *big.Int,
	fee *big.Int) (txHash string, err error) {
	privatePaymentProcessorContract, err := contracts.NewPrivatePaymentProcessorContract(contractAddresses, ch.backend)
	if err != nil {
		return
	}

	walletAuth := bind.NewKeyedTransactor(walletKey)
	walletAuth.Context = ctx
	walletAuth.GasPrice = gasPrice
	walletAuth.GasLimit = gasLimit.Uint64()

	tx, err := privatePaymentProcessorContract.PayForOrderInTokens(walletAuth, orderID, originAddress, fee, tokenAddress, orderValue)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// RefundPaymentPrivate calls refundPayment method that initiate process of funds refunding to the client.
func (ch *PaymentHandler) RefundPaymentPrivate(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	clientAddress common.Address,
	refundReason string) (txHash string, err error) {
	privatePaymentProcessorContract, err := contracts.NewPrivatePaymentProcessorContract(contractAddresses, ch.backend)
	if err != nil {
		return
	}

	processingAuth := bind.NewKeyedTransactor(ch.key)
	processingAuth.Context = ctx
	processingAuth.GasPrice = gasPrice

	tx, err := privatePaymentProcessorContract.RefundPayment(processingAuth, orderID, clientAddress, refundReason)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// RefundTokenPaymentPrivate calls refundTokenPayment method that initiate process of token funds refunding to the client.
func (ch *PaymentHandler) RefundTokenPaymentPrivate(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	clientAddress common.Address,
	refundReason string,
	tokenAddress common.Address,
	orderValue *big.Int) (txHash string, err error) {
	privatePaymentProcessorContract, err := contracts.NewPrivatePaymentProcessorContract(contractAddresses, ch.backend)
	if err != nil {
		return
	}

	processingAuth := bind.NewKeyedTransactor(ch.key)
	processingAuth.Context = ctx
	processingAuth.GasPrice = gasPrice

	tx, err := privatePaymentProcessorContract.RefundTokenPayment(processingAuth, orderID, clientAddress, refundReason, orderValue, tokenAddress)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// WithdrawRefundPrivate calls withdrawRefund method that performs fund transfer to the client's account.
func (ch *PaymentHandler) WithdrawRefundPrivate(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int) (txHash string, err error) {
	privatePaymentProcessorContract, err := contracts.NewPrivatePaymentProcessorContract(contractAddresses, ch.backend)
	if err != nil {
		return
	}

	processingAuth := bind.NewKeyedTransactor(ch.key)
	processingAuth.Context = ctx
	processingAuth.GasPrice = gasPrice

	tx, err := privatePaymentProcessorContract.WithdrawRefund(processingAuth, orderID)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// WithdrawTokenRefundPrivate calls withdrawTokenRefund method that performs token fund transfer to the client's account.
func (ch *PaymentHandler) WithdrawTokenRefundPrivate(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	tokenAddress common.Address) (txHash string, err error) {
	privatePaymentProcessorContract, err := contracts.NewPrivatePaymentProcessorContract(contractAddresses, ch.backend)
	if err != nil {
		return
	}

	processingAuth := bind.NewKeyedTransactor(ch.key)
	processingAuth.Context = ctx
	processingAuth.GasPrice = gasPrice

	tx, err := privatePaymentProcessorContract.WithdrawTokenRefund(processingAuth, orderID, tokenAddress)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}
