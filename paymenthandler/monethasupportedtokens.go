package paymenthandler

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/monetha/payment-go-sdk/contracts"
)

// GetAllTokens retrives addresses details of all supported tokens
func (ch *PaymentHandler) GetAllTokens(ctx context.Context,
	contractAddresses common.Address) (tokenAddresses []common.Address,
	tokenAcronymsBytes [][32]byte, err error) {
	monethaSupportedTokensContract, err := contracts.NewMonethaSupportedTokensContract(contractAddresses, ch.backend)
	if err != nil {
		return
	}

	tokenAddresses, tokenAcronymsBytes, err = monethaSupportedTokensContract.GetAll(nil)
	return
}
