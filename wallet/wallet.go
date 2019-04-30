package wallet

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	eth "github.com/monetha/go-ethereum"
	"github.com/monetha/payment-go-sdk/contracts"
)

var (
	// ErrContractNotFound error returned when MerchantWalletContract could not be instantiated on the given address
	ErrContractNotFound = errors.New("wallet: MerchantWallet contract not found")
	//ErrAlwaysFailingTransaction error returned when Transaction could not be submitted to blockchain
	ErrAlwaysFailingTransaction = errors.New("wallet: could not submit a transaction. Gas limit exceeds limits or always failing transaction")
)

// Wallet executes methods necessary for fund management in MerchantWallet
type Wallet struct {
	*eth.Session
	ContractAddress common.Address
	ContractHandler *contracts.MerchantWalletContract
}

// NewWallet converts session to Wallet
func NewWallet(s *eth.Session, contractAddress common.Address) *Wallet {

	merchantWalletContract, _ := contracts.NewMerchantWalletContract(contractAddress, s.Backend)

	p := &Wallet{
		s,
		contractAddress,
		merchantWalletContract,
	}
	return p
}

// WithdrawTo withdraws specific amount to a provided deposit address
func (w *Wallet) WithdrawTo(ctx context.Context,
	depositAddress common.Address,
	amount *big.Int) (txHash common.Hash, err error) {

	merchantWallet, err := w.initMerchantWalletContract()
	if err != nil {
		return err
	}

	tx, err := merchantWallet.WithdrawToExchange(&w.TransactOpts, depositAddress, amount)
	if err != nil {
		err = ErrAlwaysFailingTransaction
		return
	}

	txHash = tx.Hash()
	return
}

// WithdrawAllTo withdraws all balance or minAmount to a provided deposit address
func (w *Wallet) WithdrawAllTo(ctx context.Context,
	depositAddress common.Address,
	minAmount *big.Int) (txHash common.Hash, err error) {

	merchantWallet, err := w.initMerchantWalletContract()
	if err != nil {
		return err
	}

	tx, err := merchantWallet.WithdrawAllToExchange(&w.TransactOpts, depositAddress, minAmount)
	if err != nil {
		err = ErrAlwaysFailingTransaction
		return
	}

	txHash = tx.Hash()
	return
}

// WithdrawAllTokensTo withdraws all token balance or minAmount to a provided deposit address.
func (w *Wallet) WithdrawAllTokensTo(ctx context.Context,
	tokenAddress common.Address,
	depositAddress common.Address,
	minAmount *big.Int) (txHash common.Hash, err error) {

	merchantWallet, err := w.initMerchantWalletContract()
	if err != nil {
		return err
	}

	tx, err := merchantWallet.WithdrawAllTokensToExchange(&w.TransactOpts, tokenAddress, depositAddress, minAmount)
	if err != nil {
		err = ErrAlwaysFailingTransaction
		return
	}

	txHash = tx.Hash()
	return
}

// ChangeMerchant passes the administration of the MerchantWallet contract to another Address
func (w *Wallet) ChangeMerchant(ctx context.Context,
	to common.Address) (txHash common.Hash, err error) {

	merchantWallet, err := w.initMerchantWalletContract()
	if err != nil {
		return err
	}


	tx, err := merchantWallet.ChangeMerchantAccount(&w.TransactOpts, to)
	if err != nil {
		err = ErrAlwaysFailingTransaction
		return
	}

	txHash = tx.Hash()
	return
}

func (w *Wallet) initMerchantWalletContract() (merchantWallet *contracts.MerchantWalletContract, err error) {
	merchantWallet = w.ContractHandler

	if merchantWallet == nil {
		merchantWallet, err = contracts.NewMerchantWalletContract(w.ContractAddress, w.Backend)
		if err != nil {
			err = ErrContractNotFound
			return
		}
	}
	return
}
