package processor

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/monetha/go-ethereum"
	"github.com/monetha/payment-go-sdk/contracts"
)

var (
	// ErrTemplate error returned when ...
	ErrTemplate = errors.New("processor: template error message")
)

// Processor executes methods necessary for payment processing on chain
type Processor struct {
	*eth.Session
	ContractAddress common.Address
}

// NewProcessor converts session to Processor
func NewProcessor(s *eth.Session, contractAddress common.Address) *Processor {
	p := &Processor{
		s,
		contractAddress,
	}
	return p
}

// AddOrder creates order in Ethereum storage with the given parameters. It returns transaction hash
func (p *Processor) AddOrder(ctx context.Context,
	orderID *big.Int,
	price *big.Int,
	paymentAcceptor common.Address,
	originAddress common.Address,
	tokenAddress common.Address,
	vouchersApplied *big.Int) (txHash string, err error) {

	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(p.ContractAddress, p.Backend)
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != orderStateNull {
		err = fmt.Errorf("processor: won't call addOrder, order ID=%v state(%v) is not Null", orderID, order.State)
		return
	}

	// TODO: calculate the fee of 1.5%
	fee := big.NewInt(price.Int64() / 1000 * 15)

	tx, err := paymentProcessorContract.AddOrder(&p.TransactOpts, orderID, price, paymentAcceptor, originAddress, fee, tokenAddress, vouchersApplied)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return

}

// CancelOrder cancels the order. Canceling an order is only possible if securePay or payForOrder has not been called yet.
func (p *Processor) CancelOrder(ctx context.Context,
	gasPrice *big.Int,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int,
	cancelReason string) (txHash string, err error) {

	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(p.ContractAddress, p.Backend)
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != orderStateCreated {
		err = fmt.Errorf("processor: won't call cancelOrder, order ID=%v state(%v) is not Created", orderID, order.State)
		return
	}

	tx, err := paymentProcessorContract.CancelOrder(&p.TransactOpts, orderID, clientReputation, merchantReputation, dealHash, cancelReason)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// SecurePay makes secure payment from given wallet to payment processor contract.
func (p *Processor) SecurePay(ctx context.Context,
	gasPrice *big.Int,
	orderID *big.Int,
	walletKey *ecdsa.PrivateKey) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(p.ContractAddress, p.Backend)
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != orderStateCreated {
		err = fmt.Errorf("processor: could not call securePay, order ID=%v state(%v) is not Created", orderID, order.State)
		return
	}

	// TODO: remove wallet_key reference. KeyedTransactor comes from backend
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
func (p *Processor) SecureTokenPay(ctx context.Context,
	tokenAddress common.Address,
	gasPrice *big.Int,
	gasLimit *big.Int,
	orderID *big.Int,
	walletKey *ecdsa.PrivateKey) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(p.ContractAddress, p.Backend)
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != orderStateCreated {
		err = fmt.Errorf("processor: could not call secureTokenPay, order ID=%v state(%v) is not Created", orderID, order.State)
		return
	}

	// TODO: remove wallet_key reference. KeyedTransactor comes from backend
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
func (p *Processor) ProcessPayment(ctx context.Context,
	gasPrice *big.Int,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(p.ContractAddress, p.Backend)
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != orderStatePaid {
		err = fmt.Errorf("processor: will not call processPayment, order ID=%v state(%v) is not Paid", orderID, order.State)
		return
	}

	tx, err := paymentProcessorContract.ProcessPayment(&p.TransactOpts, orderID, clientReputation, merchantReputation, dealHash)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// RefundPayment calls refundPayment method that initiate process of funds refunding to the client.
func (p *Processor) RefundPayment(ctx context.Context,
	gasPrice *big.Int,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int,
	refundReason string) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(p.ContractAddress, p.Backend)
	if err != nil {
		return
	}

	tx, err := paymentProcessorContract.RefundPayment(&p.TransactOpts, orderID, clientReputation, merchantReputation, dealHash, refundReason)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// WithdrawRefund calls withdrawRefund method that performs fund transfer to the client's account.
func (p *Processor) WithdrawRefund(ctx context.Context,
	gasPrice *big.Int,
	orderID *big.Int) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(p.ContractAddress, p.Backend)
	if err != nil {
		return
	}

	tx, err := paymentProcessorContract.WithdrawRefund(&p.TransactOpts, orderID)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// WithdrawTokenRefund calls withdrawRefund method that performs token fund transfer to the client's account.
func (p *Processor) WithdrawTokenRefund(ctx context.Context,
	gasPrice *big.Int,
	orderID *big.Int) (txHash string, err error) {
	paymentProcessorContract, err := contracts.NewPaymentProcessorContract(p.ContractAddress, p.Backend)
	if err != nil {
		return
	}

	tx, err := paymentProcessorContract.WithdrawTokenRefund(&p.TransactOpts, orderID)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}
