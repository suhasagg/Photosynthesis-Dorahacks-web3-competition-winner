set -ex
DENOM=stake
CHAINID=icad
RLYKEY=cosmos1wt3khka7cmn5zd592x430ph4zmlhf5gfztgha6
icad version --long

# Setup Interchain Accounts Demo app
icad init --chain-id $CHAINID $CHAINID
sed -i 's#tcp://127.0.0.1:26657#tcp://0.0.0.0:26657#g' ~/.ica/config/config.toml
sed -i "s/\"stake\"/\"$DENOM\"/g" ~/.ica/config/genesis.json
sed -i 's/pruning = "syncable"/pruning = "nothing"/g' ~/.ica/config/app.toml
sed -i 's/enable = false/enable = true/g' ~/.ica/config/app.toml
sed -i -e 's/timeout_commit = "5s"/timeout_commit = "1s"/g' ~/.ica/config/config.toml
sed -i -e 's/timeout_propose = "3s"/timeout_propose = "1s"/g' ~/.ica/config/config.toml
icad keys --keyring-backend test add validator

icad add-genesis-account $(icad keys --keyring-backend test show validator -a) 1000000000000$DENOM
icad add-genesis-account $RLYKEY 1000000000000$DENOM
icad add-genesis-account cosmos1y8tcah6r989vna00ag65xcqn6mpasjjdd2e5u2 1000000000000$DENOM
# Update host chain genesis to allow x/bank/MsgSend ICA tx execution
sed -i -e 's/\"allow_messages\":.*/\"allow_messages\": [\"\/cosmos.bank.v1beta1.MsgSend\", \"\/cosmos.staking.v1beta1.MsgDelegate\"]/g' ~/.ica/config/genesis.json
icad gentx validator 900000000$DENOM --keyring-backend test --chain-id $CHAINID
icad collect-gentxs

icad start --pruning nothing
