package plasma

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/plasma/types"
)

var (
	// Block Error
	invalidOperator = errors.New("sender is not operator")

	// Transaction error
	invalidSenderSignature       = errors.New("sender signature is invalid")
	spentTransactionOutput       = errors.New("transaction output is already spent")
	mismatchedTransactionAmounts = errors.New("sum of transaction intputs and outputs are not matched")

	// Block / TX from peers
	invalidBlockNumber = errors.New("new block has invalid block number. plasma chain may not be synced")
)

type blockVerifier func(*types.Block) error

// BlockChain implements Plasma block chain service
type BlockChain struct {
	config *Config
	blocks map[uint64]*types.Block

	// TODO: store to DB
	currentBlock        *types.Block // block not mined yet
	currentBlockNumber  *big.Int     // block number of currentBlock
	blockInterval       *big.Int     // block submitted by operator
	pendingTransactions types.Transactions

	// Channels
	newBlock chan *types.Block
	quit     chan struct{}

	lock sync.RWMutex

	// callbacks
	verifyBlock blockVerifier
}

// NewBlockChain creates BlockChain instance
func NewBlockChain(config *Config, verifyBlock blockVerifier) *BlockChain {
	return &BlockChain{
		config:              config,
		blocks:              make(map[uint64]*types.Block),
		currentBlock:        &types.Block{},
		currentBlockNumber:  big.NewInt(1000),
		blockInterval:       big.NewInt(1000),
		pendingTransactions: types.Transactions{},
		newBlock:            make(chan *types.Block, 1),
		quit:                make(chan struct{}),
		verifyBlock:         verifyBlock,
	}
}

func (bc *BlockChain) GetCurrentBlock() *types.Block {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return bc.currentBlock
}

func (bc *BlockChain) GetCurrentBlockNumber() *big.Int {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return bc.currentBlockNumber
}

func (bc *BlockChain) GetBlock(blkNum *big.Int) (*types.Block, error) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	if blkNum.Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("No block with block number 0")
	}

	b, ok := bc.blocks[blkNum.Uint64()]

	if !ok {
		return nil, errors.New("No block with block number " + blkNum.String())
	}

	return b, nil
}

func (bc *BlockChain) GetPendingTransactions() (types.Transactions, error) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	return bc.pendingTransactions, nil
}

func (bc *BlockChain) GetTransaction(blkNum, txIndex *big.Int) (*types.Transaction, error) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	if blkNum.Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("No block with block number 0")
	}

	b, ok := bc.blocks[blkNum.Uint64()]

	if !ok {
		return nil, errors.New("No block with block number " + blkNum.String())
	}

	if txIndex.Cmp(big.NewInt(int64(len(b.Data.TransactionSet)))) > 0 {
		return nil, errors.New("No transaction with tx index " + txIndex.String())
	}

	tx := b.Data.TransactionSet[txIndex.Int64()]

	return tx, nil
}

// TODO: broadcast new transaction to peers
func (bc *BlockChain) ApplyTransaction(tx *types.Transaction) error {
	if err := bc.VerifyTransaction(tx); err != nil {
		log.Info("[Plasma Chain] Failed to verify transaction", "hash", tx.Hash(), "error", err)

		return err
	}

	bc.markUtxoSpent(tx.Data.BlkNum1, tx.Data.TxIndex1, tx.Data.OIndex1)
	bc.markUtxoSpent(tx.Data.BlkNum2, tx.Data.TxIndex2, tx.Data.OIndex2)

	bc.pendingTransactions = append(bc.pendingTransactions, tx)
	return nil
}

func (bc *BlockChain) VerifyTransaction(tx *types.Transaction) error {
	outputAmounts := big.NewInt(0)

	if tx.Data.Amount1 != nil {
		outputAmounts = big.NewInt(0).Add(outputAmounts, tx.Data.Amount1)
	}

	if tx.Data.Amount2 != nil {
		outputAmounts = big.NewInt(0).Add(outputAmounts, tx.Data.Amount2)
	}

	if tx.Data.Fee != nil {
		outputAmounts = big.NewInt(0).Add(outputAmounts, tx.Data.Fee)

	}

	inputAmounts := big.NewInt(0)

	if tx.Data.BlkNum1.Cmp(big.NewInt(0)) > 0 {
		preTX, err := bc.GetTransaction(tx.Data.BlkNum1, tx.Data.TxIndex1)

		if err != nil {
			return err
		}

		if err := verifyTxInput(tx, preTX, big0, tx.Data.OIndex1); err != nil {
			return err
		}

		var inputAmount *big.Int
		if tx.Data.OIndex1.Cmp(big0) == 0 {
			inputAmount = preTX.Data.Amount1
		} else {
			inputAmount = preTX.Data.Amount2
		}
		inputAmounts = big.NewInt(0).Add(inputAmounts, inputAmount)
	}

	if tx.Data.BlkNum2.Cmp(big.NewInt(0)) > 0 {
		preTX, err := bc.GetTransaction(tx.Data.BlkNum2, tx.Data.TxIndex2)

		if err != nil {
			return err
		}

		if err := verifyTxInput(tx, preTX, big1, tx.Data.OIndex2); err != nil {
			return err
		}

		var inputAmount *big.Int
		if tx.Data.OIndex1.Cmp(big0) == 0 {
			inputAmount = preTX.Data.Amount1
		} else {
			inputAmount = preTX.Data.Amount2
		}
		inputAmounts = big.NewInt(0).Add(inputAmounts, inputAmount)
	}

	if inputAmounts.Cmp(outputAmounts) != 0 {
		log.Info("[Plasma Chain] mismatched amount", "input", inputAmounts, "output", outputAmounts)
		return mismatchedTransactionAmounts
	}

	return nil
}

// verify UTXO can be spent
func verifyTxInput(tx, preTx *types.Transaction, curOIndex, preOIndex *big.Int) error {
	sender, err := tx.Sender(curOIndex)

	if err != nil {
		return err
	}

	var spent bool
	var utxoOwner *common.Address

	if preOIndex.Cmp(big.NewInt(0)) == 0 {
		spent = preTx.Spent1()
		utxoOwner = preTx.Data.NewOwner1
	} else if preOIndex.Cmp(big.NewInt(1)) == 0 {
		spent = preTx.Spent2()
		utxoOwner = preTx.Data.NewOwner2
	}

	if spent {
		return spentTransactionOutput
	}

	if sender != *utxoOwner {
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
		bc.blocks[blkNum.Uint64()].Data.TransactionSet[txIndex.Int64()].SetSpent1()
	} else {
		bc.blocks[blkNum.Uint64()].Data.TransactionSet[txIndex.Int64()].SetSpent2()
	}
}

// SubmitBlock seals current block. Only operator can seal, broadcast to peers,
// and record it on root chain
// TODO: Check the submited block is correctly recorded
func (bc *BlockChain) SubmitBlock(privKey *ecdsa.PrivateKey) (common.Hash, error) {
	bc.lock.RLock()
	defer bc.lock.RUnlock()

	b := bc.currentBlock
	b.Data.BlockNumber = big.NewInt(bc.currentBlockNumber.Int64())
	b.Data.TransactionSet = bc.pendingTransactions

	if privKey == nil {
		privKey = bc.config.OperatorPrivateKey
	}

	_, err := b.Seal()
	if err != nil {
		return common.BytesToHash(nil), err
	}

	bc.blocks[bc.currentBlockNumber.Uint64()] = bc.currentBlock
	bc.currentBlockNumber = big.NewInt(0).Add(bc.currentBlockNumber, bc.blockInterval)
	bc.currentBlock = &types.Block{}
	bc.pendingTransactions = make(types.Transactions, 0)
	bc.newBlock <- b

	return b.Hash(), nil
}

func (bc *BlockChain) NewDeposit(amount *big.Int, depositor *common.Address, depositBlockNumber *big.Int) (common.Hash, error) {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	tx := types.NewTransaction(
		depositBlockNumber, big0, big0,
		big0, big0, big0,
		depositor, amount,
		&nullAddress, big0,
		big0)

	transactionSet := []*types.Transaction{tx}

	b := types.NewBlock(depositBlockNumber, transactionSet, nil)

	// TODO: save to db
	bc.blocks[depositBlockNumber.Uint64()] = b
	bc.newBlock <- b

	log.Info("[Plasma Chain] New Deposit added", "depositBlockNumber", depositBlockNumber, "blockHash", b.Hash())

	return b.Hash(), nil
}

// TODO: use event.Feed if needed.
func (bc *BlockChain) AddNewBlockListener(listener func(blk *types.Block) error) error {
	for {
		select {
		case blk := <-bc.newBlock:
			if err := listener(blk); err != nil {
				log.Info("[Plasma Chain] Faield to listen new block", "err", err)
			}
		case <-bc.quit:
			return nil
		}
	}
}

// add deposit block or synced block
func (bc *BlockChain) AddBlock(b *types.Block) (int, error) {
	bc.lock.Lock()
	defer bc.lock.Unlock()

	blkNum := b.NumberU64()

	if _, ok := bc.blocks[blkNum]; ok {
		return 0, nil
	}

	bc.blocks[blkNum] = b

	return 1, nil
}

// add batch of blocks
func (bc *BlockChain) AddBlocks(blocks types.Blocks) (int, error) {
	n := 0
	for _, block := range blocks {
		if n2, err := bc.AddBlock(block); err != nil {
			return n, err
		} else {
			n += n2
		}
	}

	return n, nil
}
