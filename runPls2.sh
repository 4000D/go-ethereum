#!/usr/bin/env bash

DATA_DIR=~/Library/Ethereum/plasma-root-2

make geth && rm -rf $DATA_DIR && ./build/bin/geth \
 --networkid 777 \
 --identity "pls_node1" \
 --datadir $DATA_DIR \
 --nodekeyhex 09bdf6985aabc696dc1fbeb5381aebd7a6421727343872eb2fadfc6d82486fd9 \
 --nodiscover \
 --port 30305 \
 --verbosity 3 \
 --pls \
 --bootnodes "enode://c6608c3bc9307fd2462b0710198926468e50cdeaf503b7092e5e9f97c827839305688ce34f6241417b0f6977860caad7cadb13654e6624f541dbdd59dac6eeab@127.0.0.1:30303"

ㅈ # --vmodule value miner=2,plasma/*=3 \

 : '
geth attach ~/Library/Ethereum/plasma-root-2/geth.ipc

admin.addPeer("enode://c6608c3bc9307fd2462b0710198926468e50cdeaf503b7092e5e9f97c827839305688ce34f6241417b0f6977860caad7cadb13654e6624f541dbdd59dac6eeab@127.0.0.1:30303")
net.peerCount

personal.newAccount("")
personal.unlockAccount(eth.accounts[0], "", 0)

web3.pls.getBlock(1)
 '
