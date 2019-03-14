package paymenthandler

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/payment-go-sdk/contracts"
)

// WithdrawToExchange withdraws specific amount to an exchange deposit address.
func (ch *PaymentHandler) WithdrawToExchange(ctx context.Context,
	walletAddress common.Address,
	depositAddress common.Address,
	amount *big.Int) (txHash string, err error) {
	merchantWallet, err := contracts.NewMerchantWalletContract(walletAddress, ch.backend)
	if err != nil {
		return
	}

	walletAuth := bind.NewKeyedTransactor(ch.key)
	walletAuth.Context = ctx

	tx, err := merchantWallet.WithdrawToExchange(walletAuth, depositAddress, amount)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// WithdrawAllToExchange withdraws all balance to an exchange deposit address.
func (ch *PaymentHandler) WithdrawAllToExchange(ctx context.Context,
	walletAddress common.Address,
	depositAddress common.Address,
	minAmount *big.Int) (txHash string, err error) {
	merchantWallet, err := contracts.NewMerchantWalletContract(walletAddress, ch.backend)
	if err != nil {
		return
	}

	walletAuth := bind.NewKeyedTransactor(ch.key)
	walletAuth.Context = ctx

	tx, err := merchantWallet.WithdrawAllToExchange(walletAuth, depositAddress, minAmount)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}

// WithdrawAllTokensToExchange withdraws all token balance to an exchange deposit address.
func (ch *PaymentHandler) WithdrawAllTokensToExchange(ctx context.Context,
	tokenAddress common.Address,
	walletAddress common.Address,
	depositAddress common.Address,
	minAmount *big.Int) (txHash string, err error) {
	merchantWallet, err := contracts.NewMerchantWalletContract(walletAddress, ch.backend)
	if err != nil {
		return
	}

	walletAuth := bind.NewKeyedTransactor(ch.key)
	walletAuth.Context = ctx

	tx, err := merchantWallet.WithdrawAllTokensToExchange(walletAuth, tokenAddress, depositAddress, minAmount)
	if err != nil {
		return
	}

	txHash = tx.Hash().Hex()
	return
}
