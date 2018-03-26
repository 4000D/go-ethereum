package plasma

import (
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// Block Error
	invalidOperator = errors.New("sender is not operator")

	// Transaction error
	invalidSenderSignature       = errors.New("sender signature is invalid")
	spentTransactionOutput       = errors.New("transaction output is already spent")
	mismatchedTransactionAmounts = errors.New("sum of transaction intputs and outputs are not matched")
)

// BlockChain implements Plasma block chain service
type BlockChain struct {
	config              BlockChainConfig
	blocks              []*Block
	currentBlock        *Block
	currentBlockNumber  *big.Int
	pendingTransactions []*Transaction

	lock sync.RWMutex
}

// NewBlockChain creates BlockChain instance
func NewBlockChain(config BlockChainConfig) *BlockChain {
	return &BlockChain{
		config:              config,
		blocks:              []*Block{},
		currentBlock:        &Block{},
		currentBlockNumber:  big.NewInt(1),
		pendingTransactions: []*Transaction{},
	}
}

func (bc *BlockChain) getCurrentBlock() *Block {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return bc.currentBlock
}

func (bc *BlockChain) getCurrentBlockNumber() *big.Int {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return bc.currentBlockNumber
}

func (bc *BlockChain) getBlock(blkNum *big.Int) *Block {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return bc.blocks[blkNum.Int64()]
}

func (bc *BlockChain) getTransaction(blkNum, txIndex *big.Int) *Transaction {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return bc.blocks[blkNum.Int64()].transactionSet[txIndex.Int64()]
}

func (bc *BlockChain) applyTransaction(tx *Transaction) error {
	if err := bc.verifyTransaction(tx); err != nil {
		return err
	}

	bc.markUtxoSpent(tx.blkNum1, tx.txIndex1, tx.oIndex1)
	bc.markUtxoSpent(tx.blkNum2, tx.txIndex2, tx.oIndex2)

	bc.currentBlock.transactionSet = append(bc.currentBlock.transactionSet, tx)
	return nil
}

func (bc *BlockChain) verifyTransaction(tx *Transaction) error {
	outputAmounts := big.NewInt(0).Add(tx.amount1, tx.amount2)
	outputAmounts = big.NewInt(0).Add(outputAmounts, tx.fee)

	inputAmounts := big.NewInt(0)

	if tx.blkNum1.Cmp(big.NewInt(0)) > 0 {
		preTX := bc.getTransaction(tx.blkNum1, tx.txIndex1)

		if err := verifyTxInput(tx, preTX, tx.blkNum1, tx.txIndex1, tx.oIndex1); err != nil {
			return err
		}

		inputAmount := preTX.amount1
		inputAmounts = big.NewInt(0).Add(inputAmounts, inputAmount)
	}

	if tx.blkNum2.Cmp(big.NewInt(0)) > 0 {

		preTX := bc.getTransaction(tx.blkNum2, tx.txIndex2)

		if err := verifyTxInput(tx, preTX, tx.blkNum2, tx.txIndex2, tx.oIndex2); err != nil {
			return err
		}

		inputAmount := preTX.amount2
		inputAmounts = big.NewInt(0).Add(inputAmounts, inputAmount)
	}

	if inputAmounts.Cmp(outputAmounts) != 0 {
		return mismatchedTransactionAmounts
	}

	return nil
}

// verify UTXO of preTX can be spent
func verifyTxInput(tx, preTx *Transaction, blkNum, txIndex, oIndex *big.Int) error {
	sender, err := tx.Sender(oIndex)

	if err != nil {
		return err
	}

	var spent bool
	var owner common.Address

	if oIndex.Cmp(big.NewInt(0)) == 0 {
		spent = tx.spent1
		owner = tx.newOwner1
	} else if oIndex.Cmp(big.NewInt(1)) == 0 {
		spent = tx.spent2
		owner = tx.newOwner2
	}

	if spent {
		return spentTransactionOutput
	}

	if sender != owner {
		return invalidSenderSignature
	}

	return nil

}

func (bc *BlockChain) markUtxoSpent(blkNum, txIndex, oIndex *big.Int) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	if blkNum.Cmp(big.NewInt(0)) == 0 {
		return
	}

	if oIndex.Cmp(big.NewInt(0)) == 0 {
		bc.blocks[blkNum.Int64()].transactionSet[txIndex.Int64()].spent1 = true
	} else {
		bc.blocks[blkNum.Int64()].transactionSet[txIndex.Int64()].spent2 = true
	}
}

func (bc *BlockChain) submitBlock(b *Block) error {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	if sender, err := b.Sender(); err != nil {
		return err
	} else {
		if bc.config.OperatorAddress != sender {
			return invalidOperator
		}
	}

	bc.blocks[bc.currentBlockNumber.Int64()] = bc.currentBlock
	bc.currentBlockNumber = big.NewInt(0).Add(bc.currentBlockNumber, big.NewInt(1))
	bc.currentBlock = &Block{}

	return nil
}

// read transaction with hash of `txHash` from root chain
func (bc *BlockChain) submitDeposit(txHash common.Hash) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	// tx := txHash
}