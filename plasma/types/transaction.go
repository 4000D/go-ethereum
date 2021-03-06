package types

import (
	"crypto/ecdsa"
	"encoding/json"
	"io"
	"math/big"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"github.com/ethereum/go-ethereum/rlp"
)

// unsignTransaction only contains requierd fields for hash function
// TODO: change big.Int to uint64 to reduce txData size
type txData struct {
	BlkNum1   *big.Int        `json:"blkNum1"    rlp:"nil"`
	TxIndex1  *big.Int        `json:"txIndex1"   rlp:"nil"`
	OIndex1   *big.Int        `json:"oIndex1"    rlp:"nil"`
	BlkNum2   *big.Int        `json:"blkNum2"    rlp:"nil"`
	TxIndex2  *big.Int        `json:"txIndex2"   rlp:"nil"`
	OIndex2   *big.Int        `json:"oIndex2"    rlp:"nil"`
	NewOwner1 *common.Address `json:"newOwner1"  rlp:"nil"`
	Amount1   *big.Int        `json:"amount1"    rlp:"nil"`
	NewOwner2 *common.Address `json:"newOwner2"  rlp:"nil"`
	Amount2   *big.Int        `json:"amount2"    rlp:"nil"`
	Fee       *big.Int        `json:"fee"        rlp:"nil"`
}

// Transaction implements Plasma chain transaction
type Transaction struct {
	Data txData

	// included in deposit block
	isDeposit bool

	sig1 []byte
	sig2 []byte

	// whether TX output is spent or not
	spent1 bool
	spent2 bool

	// caches
	hash atomic.Value
	size atomic.Value
}

type Transactions []*Transaction

func (tx *Transaction) Sig1() []byte { return tx.sig1 }
func (tx *Transaction) Sig2() []byte { return tx.sig2 }
func (tx *Transaction) Spent1() bool { return tx.spent1 }
func (tx *Transaction) Spent2() bool { return tx.spent2 }

func (tx *Transaction) SetSig1(sig []byte) { tx.sig1 = sig }
func (tx *Transaction) SetSig2(sig []byte) { tx.sig2 = sig }
func (tx *Transaction) SetSpent1()         { tx.spent1 = true }
func (tx *Transaction) SetSpent2()         { tx.spent2 = true }

type txJSONData struct {
	blkNum1   *hexutil.Big    `json:"blkNum1"`
	txIndex1  *hexutil.Big    `json:"txIndex1"`
	oIndex1   *hexutil.Big    `json:"oIndex1"`
	blkNum2   *hexutil.Big    `json:"blkNum2"`
	txIndex2  *hexutil.Big    `json:"txIndex2"`
	oIndex2   *hexutil.Big    `json:"oIndex2"`
	newOwner1 *common.Address `json:"newOwner1"`
	amount1   *hexutil.Big    `json:"amount1"`
	newOwner2 *common.Address `json:"newOwner2"`
	amount2   *hexutil.Big    `json:"amount2"`
	fee       *hexutil.Big    `json:"fee"`
	sig1      hexutil.Bytes   `json:"sig1"`
	sig2      hexutil.Bytes   `json:"sig2"`
	spent1    *hexutil.Big    `json:"spent1"`
	spent2    *hexutil.Big    `json:"spent2"`
}

// NewTransaction creates Transaction instance
func NewTransaction(blkNum1, txIndex1, oIndex1, blkNum2, txIndex2, oIndex2 *big.Int, newOwner1 *common.Address, amount1 *big.Int, newOwner2 *common.Address, amount2, fee *big.Int) *Transaction {
	data := txData{
		blkNum1, txIndex1, oIndex1,
		blkNum2, txIndex2, oIndex2,
		newOwner1, amount1,
		newOwner2, amount2,
		fee,
	}

	if data.NewOwner1 == nil {
		data.NewOwner1 = &nullAddress
	}

	if data.NewOwner2 == nil {
		data.NewOwner2 = &nullAddress
	}

	if data.Amount1 == nil {
		data.Amount1 = big0
	}

	if data.Amount2 == nil {
		data.Amount2 = big0
	}

	if data.Fee == nil {
		data.Fee = big0
	}

	return &Transaction{
		Data:   data,
		sig1:   nil,
		sig2:   nil,
		spent1: false,
		spent2: false,
	}
}

// Hash returns sha3 hash of Transaction
func (tx *Transaction) Hash() (h common.Hash) {
	if hash := tx.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := rlpHash(tx)
	tx.hash.Store(v)
	return v
}

// EncodeRLP implements rlp.Encoder
func (tx *Transaction) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &tx.Data)
}

// DecodeRLP implements rlp.Decoder
func (tx *Transaction) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&tx.Data)
	if err == nil {
		tx.size.Store(common.StorageSize(rlp.ListSize(size)))
	}

	return err
}

func (tx *Transaction) ToRPCResponse() map[string]interface{} {

	ret := map[string]interface{}{
		"hash":      tx.Hash(),
		"blkNum1":   tx.Data.BlkNum1,
		"txIndex1":  tx.Data.TxIndex1,
		"oIndex1":   tx.Data.OIndex1,
		"blkNum2":   tx.Data.BlkNum2,
		"txIndex2":  tx.Data.TxIndex2,
		"oIndex2":   tx.Data.OIndex2,
		"newOwner1": tx.Data.NewOwner1,
		"amount1":   tx.Data.Amount1,
		"newOwner2": tx.Data.NewOwner2,
		"amount2":   tx.Data.Amount2,
		"fee":       tx.Data.Fee,
	}

	if len(tx.sig1) > 0 {
		ret["v1"] = tx.sig1[64]
		ret["r1"] = tx.sig1[0:32]
		ret["s1"] = tx.sig1[32:64]
	}

	if len(tx.sig2) > 0 {
		ret["v1"] = tx.sig2[64]
		ret["r1"] = tx.sig2[0:32]
		ret["s1"] = tx.sig2[32:64]
	}

	return ret
}

func (txs Transactions) ToRPCResponse() []map[string]interface{} {

	ret := make([]map[string]interface{}, len(txs))

	for i, tx := range txs {
		ret[i] = tx.ToRPCResponse()
	}

	return ret
}

func (tx *Transaction) MarshalJSON() ([]byte, error) {

	var enc txJSONData

	enc.blkNum1 = (*hexutil.Big)(tx.Data.BlkNum1)
	enc.txIndex1 = (*hexutil.Big)(tx.Data.TxIndex1)
	enc.oIndex1 = (*hexutil.Big)(tx.Data.OIndex1)
	enc.blkNum2 = (*hexutil.Big)(tx.Data.BlkNum2)
	enc.txIndex2 = (*hexutil.Big)(tx.Data.TxIndex2)
	enc.oIndex2 = (*hexutil.Big)(tx.Data.OIndex2)
	enc.newOwner1 = tx.Data.NewOwner1
	enc.amount1 = (*hexutil.Big)(tx.Data.Amount1)
	enc.newOwner2 = tx.Data.NewOwner2
	enc.amount2 = (*hexutil.Big)(tx.Data.Amount2)
	enc.fee = (*hexutil.Big)(tx.Data.Fee)

	enc.sig1 = tx.sig1
	enc.sig2 = tx.sig2

	if tx.spent1 {
		enc.spent1 = (*hexutil.Big)(big1)
	} else {
		enc.spent1 = (*hexutil.Big)(big0)
	}
	if tx.spent2 {
		enc.spent2 = (*hexutil.Big)(big1)
	} else {
		enc.spent2 = (*hexutil.Big)(big0)
	}

	return json.Marshal(&enc)
}

// Sender returns owner address of TX input
func (tx *Transaction) Sender(oIndex *big.Int) (common.Address, error) {
	var sig []byte

	if oIndex.Cmp(big.NewInt(0)) == 0 {
		sig = tx.sig1
	} else if oIndex.Cmp(big.NewInt(1)) == 0 {
		sig = tx.sig2
	}

	return getSender(tx.Hash().Bytes(), sig)
}

func (tx *Transaction) Sign1(privKey *ecdsa.PrivateKey) error {
	sig, err := crypto.Sign(tx.Hash().Bytes(), privKey)

	if err != nil {
		return err
	}

	tx.sig1 = sig
	return nil
}

func (tx *Transaction) Sign2(privKey *ecdsa.PrivateKey) error {
	sig, err := crypto.Sign(tx.Hash().Bytes(), privKey)

	if err != nil {
		return err
	}

	tx.sig2 = sig
	return nil
}

func getSender(hash, sig []byte) (common.Address, error) {
	pubKeyBytes, err := crypto.Ecrecover(hash, sig)

	if err != nil {
		return common.Address{}, err
	}

	pubKey := crypto.ToECDSAPub(pubKeyBytes)
	return crypto.PubkeyToAddress(*pubKey), nil
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func sign(hash []byte, privKey *ecdsa.PrivateKey) (sig []byte, err error) {
	sig, err = crypto.Sign(hash, privKey)
	return
}
