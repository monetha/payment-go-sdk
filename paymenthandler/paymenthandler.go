package paymenthandler

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/monetha/go-ethereum/backend"
)

const (
	// order states in Ethereum storage
	orderStateNull = iota
	orderStateCreated
	orderStatePaid
	//orderStateFinalized
	//orderStateRefunding
	//orderStateRefunded
	//orderStateCanceled
)

// PaymentHandler is an implementation for payment processors and merchant wallet.
type PaymentHandler struct {
	backend chequebook.Backend
	key     *ecdsa.PrivateKey
	address common.Address
}

// New creates an instance of ContractHandler which makes JSON-RPC calls.
func New(rawRPCURL string, key *ecdsa.PrivateKey) (*PaymentHandler, error) {
	if key == nil {
		return nil, fmt.Errorf("non nil key expected")
	}

	cl, err := ethclient.Dial(rawRPCURL)
	if err != nil {
		return nil, fmt.Errorf("PaymentHandler: ethclient.Dial: %v", err)
	}

	return newWithBackend(cl, key), nil
}

// NewWithBackend creates a new ContractHandler instance with provided backend instead.
func newWithBackend(be chequebook.Backend, key *ecdsa.PrivateKey) *PaymentHandler {
	// Wrap backend to have up to date nonce value (transactions count) for processing account.
	address := crypto.PubkeyToAddress(key.PublicKey)
	be = backend.NewHandleNonceBackend(be, []common.Address{address})

	return &PaymentHandler{
		backend: be,
		key:     key,
		address: address,
	}
}
