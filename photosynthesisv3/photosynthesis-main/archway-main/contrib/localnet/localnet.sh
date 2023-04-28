#!/bin/sh

# this script instantiates localnet required genesis files

set -e

echo initting new chain
# init config files

archwayd init archwayd-id --chain-id localnet --overwrite

# create accounts
archwayd keys add fd --keyring-backend=test

addr=$(archwayd keys show fd -a --keyring-backend=test)
val_addr=$(archwayd keys show fd  --keyring-backend=test --bech val -a)

# give the accounts some money
archwayd add-genesis-account "$addr" 1000000000000stake --keyring-backend=test

archwayd keys add ad --keyring-backend=test

addr1=$(archwayd keys show ad -a --keyring-backend=test)
val_addr1=$(archwayd keys show ad  --keyring-backend=test --bech val -a)
# give the accounts some money
archwayd add-genesis-account "$addr1" 1000000000000stake --keyring-backend=test

# save configs for the daemon
archwayd gentx fd 10000000stake --chain-id localnet --keyring-backend=test

# input genTx to the genesis file
archwayd collect-gentxs
# verify genesis file is fine
archwayd validate-genesis
echo changing network settings
sed -i 's/127.0.0.1/0.0.0.0/g' $HOME/.archway/config/config.toml

echo test account address: "$addr"
echo test account private key: "$(yes | archwayd keys export fd --unsafe --unarmored-hex --keyring-backend=test)"
echo account for --from flag "fd"

echo starting network...
archwayd start
