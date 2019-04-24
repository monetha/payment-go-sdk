package wallet

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	eth "github.com/monetha/go-ethereum"
	"github.com/monetha/payment-go-sdk/contracts"
)

// Wallet executes methods necessary for fund management in MerchantWallet
type Wallet struct {
	*eth.Session
	ContractAddress common.Address
}

// NewWallet converts session to Wallet
func NewWallet(s *eth.Session, contractAddress common.Address) *Wallet {
	p := &Wallet{
		s,
		contractAddress,
	}
	return p
}

// WithdrawTo withdraws specific amount to an exchange deposit address.
func (w *Wallet) WithdrawTo(ctx context.Context,
	walletAddress common.Address,
	depositAddress common.Address,
	amount *big.Int) (txHash common.Hash, err error) {

	merchantWallet, err := contracts.NewMerchantWalletContract(walletAddress, w.Backend)
	if err != nil {
		return
	}

	tx, err := merchantWallet.WithdrawToExchange(&w.TransactOpts, depositAddress, amount)
	if err != nil {
		return
	}

	txHash = tx.Hash()
	return
}

// WithdrawAllTo withdraws all balance to an exchange deposit address.
func (w *Wallet) WithdrawAllTo(ctx context.Context,
	walletAddress common.Address,
	depositAddress common.Address,
	minAmount *big.Int) (txHash common.Hash, err error) {
	merchantWallet, err := contracts.NewMerchantWalletContract(walletAddress, w.Backend)
	if err != nil {
		return
	}

	tx, err := merchantWallet.WithdrawAllToExchange(&w.TransactOpts, depositAddress, minAmount)
	if err != nil {
		return
	}

	txHash = tx.Hash()
	return
}

// WithdrawAllTokensTo withdraws all token balance to an exchange deposit address.
func (w *Wallet) WithdrawAllTokensTo(ctx context.Context,
	tokenAddress common.Address,
	walletAddress common.Address,
	depositAddress common.Address,
	minAmount *big.Int) (txHash common.Hash, err error) {
	merchantWallet, err := contracts.NewMerchantWalletContract(walletAddress, w.Backend)
	if err != nil {
		return
	}

	tx, err := merchantWallet.WithdrawAllTokensToExchange(&w.TransactOpts, tokenAddress, depositAddress, minAmount)
	if err != nil {
		return
	}

	txHash = tx.Hash()
	return
}

// TODO: add `changeMerchantAccount`
