#!/usr/bin/env bash

DATA_DIR=~/Library/Ethereum/plasma-root-1

make geth && rm -rf $DATA_DIR && ./build/bin/geth \
 --networkid 777 \
 --identity "pls_operator" \
 --datadir $DATA_DIR \
 --nodekeyhex 9cd69f009ac86203e54ec50e3686de95ff6126d3b30a19f926a0fe9323c17181 \
 --nodiscover \
 --mine \
 --minerthreads 1 \
 --etherbase 0x4c910ce23172578135467e20Bc2cF03e93B0d250 \
 --verbosity 3 \
 --pls \
 --pls.operator \
 --pls.operatorPrivKey 9cd69f009ac86203e54ec50e3686de95ff6126d3b30a19f926a0fe9323c17181 \
 --debug
 
 # --lightserv 30 \
 # --lightpeers 25 \


 : '
geth attach ~/Library/Ethereum/plasma-root-1/geth.ipc
admin.addPeer("enode://3eab09a73a40ff703d138b623026adae25be09af3c733e323ed33dba47caf639d956cb60831192063057716fde1b1d4bd8e3dd73eb3b795d80d8654d40c1d292@127.0.0.1:30305")
admin.addPeer("enode://3eab09a73a40ff703d138b623026adae25be09af3c733e323ed33dba47caf639d956cb60831192063057716fde1b1d4bd8e3dd73eb3b795d80d8654d40c1d292@127.0.0.1:62734")

web3.eth.sendTransaction({
  from: web3.eth.accounts[0],
  to: "0x4Cf98814898Cd44fd86bBCA9D82eE8B7D61b19A1",
  value: 100,
  data: "0xd0e30db0",
  gas: 150000,
})

web3.pls.submitBlock()

web3.eth.sendTransaction({
  from: web3.eth.accounts[0],
  to: "0xd5f8d6067f49b6c9887bf473b87b5a1ef9991f71",
  value: 1e18
})
 '
