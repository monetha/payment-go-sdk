package processor

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	eth "github.com/monetha/go-ethereum"
	"github.com/monetha/payment-go-sdk/contracts"
)

var (
	// ErrContractNotFound error returned when PaymentProcessorContract could not be instantiated on the given address
	ErrContractNotFound = errors.New("processor: PaymentProcessorContract not found")
	// ErrOrderStateNotNull error returned when Order's state is not Null
	ErrOrderStateNotNull = errors.New("processor: can not execute addOrder order's state is not Null")
	// ErrOrderNotCreated error returned when Order's state is not Created
	ErrOrderNotCreated = errors.New("processor: could not execute SecurePay order's state is not Created")
	//ErrAlwaysFailingTransaction error returned when Transaction could not be submitted to blockchain
	ErrAlwaysFailingTransaction = errors.New("processor: could not submit a transaction. Gas limit exceeds limits or always failing transaction")
)

// Processor executes methods necessary for payment processing on chain
type Processor struct {
	*eth.Session
	ContractAddress common.Address
	ContractHandler *contracts.PaymentProcessorContract
}

// NewProcessor converts session to Processor
func NewProcessor(s *eth.Session, contractAddress common.Address) *Processor {

	paymentProcessorContract, _ := contracts.NewPaymentProcessorContract(contractAddress, s.Backend)

	p := &Processor{
		s,
		contractAddress,
		paymentProcessorContract,
	}
	return p
}

// AddOrder creates order in Ethereum storage with the given parameters
func (p *Processor) AddOrder(ctx context.Context,
	orderID *big.Int,
	price *big.Int,
	originAddress common.Address,
	tokenAddress common.Address,
	vouchersApplied *big.Int) (txHash common.Hash, err error) {

	paymentProcessorContract, err := p.initPaymentProcessorContract()
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != OrderStateNull {
		err = ErrOrderStateNotNull
		return
	}

	fee := big.NewInt(price.Int64() / 1000 * 15)

	tx, err := paymentProcessorContract.AddOrder(&p.TransactOpts, orderID, price, originAddress, originAddress, fee, tokenAddress, vouchersApplied)
	if err != nil {
		err = ErrAlwaysFailingTransaction
		return
	}

	txHash = tx.Hash()
	return

}

// CancelOrder cancels the order. Canceling an order is possible only if Order was not paid yet.
func (p *Processor) CancelOrder(ctx context.Context,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int,
	cancelReason string) (txHash common.Hash, err error) {

	paymentProcessorContract, err := p.initPaymentProcessorContract()
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != OrderStateCreated {
		err = ErrOrderNotCreated
		return
	}

	tx, err := paymentProcessorContract.CancelOrder(&p.TransactOpts, orderID, clientReputation, merchantReputation, dealHash, cancelReason)
	if err != nil {
		return
	}

	txHash = tx.Hash()
	return
}

// SecurePay makes a secure transfer of funds from Customer's address to PaymentProcessor contract.
func (p *Processor) SecurePay(ctx context.Context,
	orderID,
	price *big.Int) (txHash common.Hash, err error) {

	paymentProcessorContract, err := p.initPaymentProcessorContract()
	if err != nil {
		return
	}

	// Set the fund value that will be transferred to a Contract
	p.TransactOpts.Value = price

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		err = fmt.Errorf("processor: %v", err)
		return
	}

	if order.State != OrderStateCreated {
		err = ErrOrderNotCreated
		return
	}

	tx, err := paymentProcessorContract.SecurePay(&p.TransactOpts, orderID)
	if err != nil {
		err = ErrAlwaysFailingTransaction
		return
	}

	txHash = tx.Hash()
	return
}

// SecureTokenPay makes a secure transfer of ERC20 funds from Customer's address to PaymentProcessor contract.
func (p *Processor) SecureTokenPay(ctx context.Context,
	tokenAddress common.Address,
	gasLimit *big.Int,
	orderID *big.Int) (txHash common.Hash, err error) {

	paymentProcessorContract, err := p.initPaymentProcessorContract()
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != OrderStateCreated {
		err = fmt.Errorf("processor: could not call secureTokenPay, order ID=%v state(%v) is not Created", orderID, order.State)
		return
	}

	tx, err := paymentProcessorContract.SecureTokenPay(&p.TransactOpts, orderID)
	if err != nil {
		return
	}

	txHash = tx.Hash()
	return
}

// ProcessPayment unlocks funds from the contract and transfers them to MerchantWallet.
func (p *Processor) ProcessPayment(ctx context.Context,
	orderID *big.Int,
	dealHash *big.Int) (txHash common.Hash, err error) {

	paymentProcessorContract, err := p.initPaymentProcessorContract()
	if err != nil {
		return
	}

	order, err := paymentProcessorContract.Orders(nil, orderID)
	if err != nil {
		return
	}

	if order.State != OrderStatePaid {
		err = fmt.Errorf("processor: will not call processPayment, order ID=%v state(%v) is not Paid", orderID, order.State)
		return
	}

	// reputation is passed as 0, as no reputation calculation is done in context of the payment
	tx, err := paymentProcessorContract.ProcessPayment(&p.TransactOpts, orderID, 0, 0, dealHash)
	if err != nil {
		return
	}

	txHash = tx.Hash()
	return
}

// RefundPayment calls refundPayment method that initiate process of funds refunding to the client.
func (p *Processor) RefundPayment(ctx context.Context,
	orderID *big.Int,
	clientReputation uint32,
	merchantReputation uint32,
	dealHash *big.Int,
	refundReason string) (txHash common.Hash, err error) {

	paymentProcessorContract, err := p.initPaymentProcessorContract()
	if err != nil {
		return
	}

	tx, err := paymentProcessorContract.RefundPayment(&p.TransactOpts, orderID, clientReputation, merchantReputation, dealHash, refundReason)
	if err != nil {
		return
	}

	txHash = tx.Hash()
	return
}

// WithdrawRefund calls withdrawRefund method that performs fund transfer to the client's account.
func (p *Processor) WithdrawRefund(ctx context.Context,
	orderID *big.Int) (txHash common.Hash, err error) {

	paymentProcessorContract, err := p.initPaymentProcessorContract()
	if err != nil {
		return
	}

	tx, err := paymentProcessorContract.WithdrawRefund(&p.TransactOpts, orderID)
	if err != nil {
		return
	}

	txHash = tx.Hash()
	return
}

// WithdrawTokenRefund calls withdrawRefund method that performs token fund transfer to the client's account.
func (p *Processor) WithdrawTokenRefund(ctx context.Context,
	orderID *big.Int) (txHash common.Hash, err error) {

	paymentProcessorContract, err := p.initPaymentProcessorContract()
	if err != nil {
		return
	}

	tx, err := paymentProcessorContract.WithdrawTokenRefund(&p.TransactOpts, orderID)
	if err != nil {
		return
	}

	txHash = tx.Hash()
	return
}

func (p *Processor) initPaymentProcessorContract() (paymentProcessorContract *contracts.PaymentProcessorContract, err error) {
	paymentProcessorContract = p.ContractHandler

	if paymentProcessorContract == nil {
		paymentProcessorContract, err = contracts.NewPaymentProcessorContract(p.ContractAddress, p.Backend)
		if err != nil {
			err = ErrContractNotFound
			return
		}
	}
	return
}
