package plasma

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

// Config represents the configuration state of a plasma node.
type Config struct {
	// Address of plasma contract on root chain
	ContractAddress common.Address

	// Address of plasma operator
	// TODO: load this address from Plasma contract on Ethereum network
	OperatorAddress common.Address

	// If this node is operator, specify the private key
	OperatorPrivateKey *ecdsa.PrivateKey

	// Plasma operator node
	OperatorNode    *Peer
	OperatorNodeURL string

	// BlockChain specific configs
	DataDir string
	OnDisk  bool // disk or memory
}

var DefaultConfig = Config{
	ContractAddress: common.StringToAddress("0x0"),
	OperatorAddress: params.PlasmaOperatorAddress,
}
