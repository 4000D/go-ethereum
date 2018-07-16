#!/usr/bin/env bash

DATA_DIR=~/Library/Ethereum/plasma-root


make geth && rm -rf $DATA_DIR && ./build/bin/geth \
 --networkid 777 \
 --identity "pls_operator" \
 --nodekeyhex 9cd69f009ac86203e54ec50e3686de95ff6126d3b30a19f926a0fe9323c17181 \
 --etherbase 0x4c910ce23172578135467e20Bc2cF03e93B0d250 \
 --pls \
 --pls.operator \
 --pls.operatorPrivKey 9cd69f009ac86203e54ec50e3686de95ff6126d3b30a19f926a0fe9323c17181 \
 --rpc \
 --mine \
 --minerthreads 1 \
 --verbosity 3 \
 --debug


: '
geth attach ~/Library/Ethereum/plasma-root/geth.ipc

# compare rootchain.childChain vs plasmachain.getBlock
web3.eth.sendTransaction({
  from: web3.eth.accounts[0],
  to: "0x4Cf98814898Cd44fd86bBCA9D82eE8B7D61b19A1",
  value: 100,
  data: "0xd0e30db0",
  gas: 150000,
})
web3.pls.getBlock(1)
web3.pls.getChildChain(1)

# prepare other account
personal.newAccount("")
personal.unlockAccount(eth.accounts[1], "", 0)
web3.eth.sendTransaction({from: eth.accounts[0], to: eth.accounts[1], value: 1e18})

# deposit from account 0 & 1
web3.eth.sendTransaction({
  from: web3.eth.accounts[0],
  to: "0x4Cf98814898Cd44fd86bBCA9D82eE8B7D61b19A1",
  value: 100,
  data: "0xd0e30db0",
  gas: 150000,
})
web3.eth.sendTransaction({
  from: web3.eth.accounts[1],
  to: "0x4Cf98814898Cd44fd86bBCA9D82eE8B7D61b19A1",
  value: 100,
  data: "0xd0e30db0",
  gas: 150000,
})

# get deposit block
web3.pls.getBlock(1)
web3.pls.getBlock(2)

# apply tx 1
web3.pls.applyTransaction({
  blkNum1: 1,
  txIndex1: 0,
  oIndex1: 0,
  blkNum2: 0,
  txIndex2: 0,
  oIndex2: 0,
  newOwner1: web3.eth.accounts[1],
  amount1: 90,
  fee: 10,
  from1: web3.eth.accounts[0],
})
# apply tx 2
web3.pls.applyTransaction({
  blkNum1: 2,
  txIndex1: 0,
  oIndex1: 0,
  blkNum2: 0,
  txIndex2: 0,
  oIndex2: 0,
  newOwner1: web3.eth.accounts[1],
  amount1: 90,
  fee: 10,
  from1: web3.eth.accounts[1],
})

# submit block 3
web3.pls.submitBlock()

web3.pls.getBlock(1000)



# apply tx 3
web3.pls.applyTransaction({
  blkNum1: 1000,
  txIndex1: 0,
  oIndex1: 0,
  blkNum2: 1000,
  txIndex2: 1,
  oIndex2: 0,
  newOwner1: web3.eth.accounts[1],
  amount1: 180,
  fee: 0,
  from1: web3.eth.accounts[1],
  from2: web3.eth.accounts[1],
})

# submit block 4
web3.pls.submitBlock()

web3.pls.getBlock(2000)


web3.pls.applyTransaction({   blkNum1: 1,   txIndex1: 0,   oIndex1: 0,   blkNum2: 0,   txIndex2: 0,   oIndex2: 0,   newOwner1: web3.eth.accounts[0],   amount1: 90,   fee: 10,   from1: web3.eth.accounts[0], })
# submit
web3.pls.submitBlock()


web3.eth.sendTransaction({   from: web3.eth.accounts[0],   to: "0x4Cf98814898Cd44fd86bBCA9D82eE8B7D61b19A1",   value: 100,   data: "0xd0e30db0",   gas: 150000, })
web3.pls.applyTransaction({   blkNum1: 1,   txIndex1: 0,   oIndex1: 0,   blkNum2: 0,   txIndex2: 0,   oIndex2: 0,   newOwner1: web3.eth.accounts[0],   amount1: 90,   fee: 10,   from1: web3.eth.accounts[0], })


# check balance
web3.eth.getBalance("0x4Cf98814898Cd44fd86bBCA9D82eE8B7D61b19A1")




'
