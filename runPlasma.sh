#!/usr/bin/env bash

rm -rf ~/Library/Ethereum/plasma-root && ./build/bin/geth \
 --pls \
 --pls.operator \
 --pls.operatorPrivKey 9cd69f009ac86203e54ec50e3686de95ff6126d3b30a19f926a0fe9323c17181 \
 --rpc \
 --mine


: '
geth attach http://localhost:8545

# deploy tx
web3.eth.getTransactionReceipt("0xa12c74a0cc21d5441092af0a5afd324fb95ee3b59805c9ea118c5c0974357983")

# deposit
web3.eth.sendTransaction({
  from: web3.eth.accounts[0],
  to: "0x4Cf98814898Cd44fd86bBCA9D82eE8B7D61b19A1",
  value: 100,
  data: "0xd0e30db0",
  gas: 150000,
})

web3.eth.getTransaction("0x591471329bc4d5085ef438f65415ac801baa902492abd7089dcf0f7ab70f93e3)
web3.eth.getTransactionReceipt("0x591471329bc4d5085ef438f65415ac801baa902492abd7089dcf0f7ab70f93e3")


web3.eth.getTransactionCount(web3.eth.accounts[0])
web3.eth.getBalance("0x4Cf98814898Cd44fd86bBCA9D82eE8B7D61b19A1")
'
