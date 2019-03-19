package paymenthandler

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/payment-go-sdk/contracts"
)

// AddOrder creates order in Ethereum storage with the given parameters.
func (ph *PaymentHandler) AddOrder(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	price *big.Int,
	paymentAcceptor common.Address,
	originAddress common.Address,
	fee *big.Int,
	tokenAddress common.Address,
	vouchersApplied *big.Int) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(contractAddresses, ph.backend)
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != orderStateNull {
		err = fmt.Errorf("PaymentHandler: won't call addOrder, order ID=%v state(%v) is not Null", orderID, order.State)
		return
	}

	// using processing account key to call addOrder method
	processingAuth := bind.NewKeyedTransactor(ph.key)
	processingAuth.Context = ctx
	processingAuth.GasPrice = gasPrice

	tx, err := paymentProcessorContract.AddOrder(processingAuth, orderID, price, paymentAcceptor, originAddress, fee, tokenAddress, vouchersApplied)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return

}

// CancelOrder cancels the order. Canceling an order is only possible if securePay or payForOrder has not been called yet.
func (ph *PaymentHandler) CancelOrder(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int,
	cancelReason string) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(contractAddresses, ph.backend)
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != orderStateCreated {
		err = fmt.Errorf("PaymentHandler: won't call cancelOrder, order ID=%v state(%v) is not Created", orderID, order.State)
		return
	}

	processingAuth := bind.NewKeyedTransactor(ph.key)
	processingAuth.Context = ctx
	processingAuth.GasPrice = gasPrice

	tx, err := paymentProcessorContract.CancelOrder(processingAuth, orderID, clientReputation, merchantReputation, dealHash, cancelReason)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// SecurePay makes secure payment from given wallet to payment processor contract.
func (ph *PaymentHandler) SecurePay(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	walletKey *ecdsa.PrivateKey) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(contractAddresses, ph.backend)
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != orderStateCreated {
		err = fmt.Errorf("PaymentHandler: wont'c call securePay, order ID=%v state(%v) is not Created", orderID, order.State)
		return
	}

	walletAuth := bind.NewKeyedTransactor(walletKey)
	walletAuth.Context = ctx
	walletAuth.Value = order.Price
	walletAuth.GasPrice = gasPrice

	tx, err := paymentProcessorContract.SecurePay(walletAuth, orderID)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// SecureTokenPay makes secure payment from given wallet to payment processor contract.
func (ph *PaymentHandler) SecureTokenPay(ctx context.Context,
	contractAddress, tokenAddress common.Address,
	gasPrice *big.Int,
	gasLimit *big.Int,
	orderID *big.Int,
	walletKey *ecdsa.PrivateKey) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(contractAddress, ph.backend)
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != orderStateCreated {
		err = fmt.Errorf("PaymentHandler: wont'c call secureTokenPay, order ID=%v state(%v) is not Created", orderID, order.State)
		return
	}

	walletAuth := bind.NewKeyedTransactor(walletKey)

	walletAuth.Value = big.NewInt(0)
	walletAuth.Context = ctx
	walletAuth.GasPrice = gasPrice
	walletAuth.GasLimit = gasLimit.Uint64()

	tx, err := paymentProcessorContract.SecureTokenPay(walletAuth, orderID)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// ProcessPayment transfers funds to MonethaGateway, updates client/merchant reputation and completes the order.
func (ph *PaymentHandler) ProcessPayment(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(contractAddresses, ph.backend)
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != orderStatePaid {
		err = fmt.Errorf("PaymentHandler: wont'c call processPayment, order ID=%v state(%v) is not Paid", orderID, order.State)
		return
	}

	processingAuth := bind.NewKeyedTransactor(ph.key)
	processingAuth.Context = ctx
	processingAuth.GasPrice = gasPrice

	tx, err := paymentProcessorContract.ProcessPayment(processingAuth, orderID, clientReputation, merchantReputation, dealHash)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// RefundPayment calls refundPayment method that initiate process of funds refunding to the client.
func (ph *PaymentHandler) RefundPayment(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int,
	refundReason string) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(contractAddresses, ph.backend)
	if err != nil {
		return
	}

	processingAuth := bind.NewKeyedTransactor(ph.key)
	processingAuth.Context = ctx
	processingAuth.GasPrice = gasPrice

	tx, err := paymentProcessorContract.RefundPayment(processingAuth, orderID, clientReputation, merchantReputation, dealHash, refundReason)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// WithdrawRefund calls withdrawRefund method that performs fund transfer to the client's account.
func (ph *PaymentHandler) WithdrawRefund(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(contractAddresses, ph.backend)
	if err != nil {
		return
	}

	processingAuth := bind.NewKeyedTransactor(ph.key)
	processingAuth.Context = ctx
	processingAuth.GasPrice = gasPrice

	tx, err := paymentProcessorContract.WithdrawRefund(processingAuth, orderID)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// WithdrawTokenRefund calls withdrawRefund method that performs token fund transfer to the client's account.
func (ph *PaymentHandler) WithdrawTokenRefund(ctx context.Context,
	contractAddresses common.Address,
	gasPrice *big.Int,
	orderID *big.Int) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(contractAddresses, ph.backend)
	if err != nil {
		return
	}

	processingAuth := bind.NewKeyedTransactor(ph.key)
	processingAuth.Context = ctx
	processingAuth.GasPrice = gasPrice

	tx, err := paymentProcessorContract.WithdrawTokenRefund(processingAuth, orderID)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}
