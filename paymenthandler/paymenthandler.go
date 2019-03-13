package paymenthandler

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/chequebook"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	ethereum "github.com/monetha/go-ethereum"
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
func New(rawRPCURL string, key string) (*PaymentHandler, error) {
	processingKey, err := ethereum.NewKeyFromPrivateKey(key)
	if err != nil {
		return nil, err
	}

	cl, err := ethclient.Dial(rawRPCURL)
	if err != nil {
		return nil, err
	}

	return newWithBackend(cl, processingKey.PrivateKey), nil
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
