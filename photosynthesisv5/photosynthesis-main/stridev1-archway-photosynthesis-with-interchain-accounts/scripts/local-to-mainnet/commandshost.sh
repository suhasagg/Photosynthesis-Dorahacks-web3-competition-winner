############################################
### WARNING: THIS FILE IS AUTOGENERATED. ###
###   ANY CHANGES WILL BE OVERWRITTEN.   ###
############################################
############################################
### WARNING: THIS FILE IS AUTOGENERATED. ###
###   ANY CHANGES WILL BE OVERWRITTEN.   ###
############################################
############################################
### WARNING: THIS FILE IS AUTOGENERATED. ###
###   ANY CHANGES WILL BE OVERWRITTEN.   ###
############################################
#### SETUP HOT WALLET (Only needs to be run once)
echo "$HOT_WALLET_1_MNEMONIC" | build/archwayd keys add hot --recover --keyring-backend test 


#### START RELAYER
# NOTE: CREATING CONNECTIONS WITH THE GO RELAYER 
# Create connections and channels
docker-compose -f scripts/local-to-mainnet/docker-compose.yml run --rm relayer rly transact link stride-host 

# (OR) If the go relayer isn't working, use hermes (NOTE: you'll have to add the connections to the relayer config in `scripts/state/relayer/config/config.yaml`)
# docker-compose -f scripts/local-to-mainnet/docker-compose.yml run --rm hermes hermes create connection --a-chain cosmoshub-4 --b-chain stride
# docker-compose -f scripts/local-to-mainnet/docker-compose.yml run --rm hermes hermes create channel --a-chain stride --a-connection connection-0 --a-port transfer --b-port transfer

# Ensure Relayer Config is updated (`scripts/local-to-mainnet/state/relayer/config/config.yaml`)
#    paths:
#     stride-host:
#       src:
#         chain-id: stride
#         client-id: 07-tendermint-0
#         connection-id: connection-0
#       dst:
#         chain-id: cosmoshub-4
#         client-id: {CLIENT-ID}
#         connection-id: {CONNECTION-ID}

# Get channel ID created on the host
build/strided --home scripts/state/stride1 q ibc channel channels 
transfer_channel=$(build/strided --home scripts/state/stride1 q ibc channel channels | grep channel-0 -A 4 | grep counterparty -A 1 | grep channel | awk '{print $2}') && echo $transfer_channel

# Start Go Relayer 
docker-compose -f scripts/local-to-mainnet/docker-compose.yml up -d relayer
docker-compose -f scripts/local-to-mainnet/docker-compose.yml logs -f relayer | sed -r -u "s/\x1B\[([0-9]{1,3}(;[0-9]{1,2})?)?[mGK]//g" >> scripts/logs/relayer.log 2>&1 &


#### REGISTER HOST
# IBC Transfer from HOST to stride (from relayer account)
build/archwayd tx ibc-transfer transfer transfer $transfer_channel stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 4000000uatom --from hot --chain-id cosmoshub-4 -y --keyring-backend test --node http://HOST_ENDPOINT:26657 --fees 150000uatom

# Confirm funds were recieved on stride and get IBC denom
build/strided --home scripts/state/stride1 q bank balances stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8

# Register host zone
IBC_DENOM=$(build/strided --home scripts/state/stride1 q bank balances stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 | grep ibc | awk '{print $2}' | tr -d '"') && echo $IBC_DENOM
build/strided --home scripts/state/stride1 tx stakeibc register-host-zone \
    connection-0 uatom archway $IBC_DENOM channel-0 1 \
    --from admin --gas 1000000 -y

# Add validator
echo '{"validators": [{"name": "HOST_VAL_NAME_1", "address": "HOST_VAL_ADDRESS_1", "weight": 10}]}' > validator.json
build/strided --home scripts/state/stride1 tx stakeibc add-validators cosmoshub-4 validator.json --gas 1000000 --chain-id stride --keyring-backend test --from admin -y
rm validator.json

# Confirm ICA channels were registered
build/strided --home scripts/state/stride1 q stakeibc list-host-zone


#### FLOW
## Go Through Flow
# Liquid stake (then wait and LS again)
build/strided --home scripts/state/stride1 tx stakeibc liquid-stake 1000000 uatom --keyring-backend test --from admin -y --chain-id stride -y

# Confirm stTokens, StakedBal, and Redemption Rate
build/strided --home scripts/state/stride1 q bank balances stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
build/strided --home scripts/state/stride1 q stakeibc list-host-zone

# Redeem
build/strided --home scripts/state/stride1 tx stakeibc redeem-stake 1000 cosmoshub-4 HOT_WALLET_ADDRESS --from admin --keyring-backend test --chain-id stride -y

# Confirm stTokens and StakedBal
build/strided --home scripts/state/stride1 q bank balances stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8
build/strided --home scripts/state/stride1 q stakeibc list-host-zone

# Add another validator
echo '{"validators": [{"name": "HOST_VAL_NAME_2", "address": "HOST_VAL_ADDRESS_2", "weight": 10}]}' > validator.json
build/strided --home scripts/state/stride1 tx stakeibc add-validators cosmoshub-4 validator.json --gas 1000000 --chain-id stride --keyring-backend test --from admin -y
rm validator.json

# Liquid stake and confirm the stake was split 50/50 between the validators
build/strided --home scripts/state/stride1 tx stakeibc liquid-stake 1000000 uatom --keyring-backend test --from admin -y --chain-id stride -y

# Change validator weights
build/strided --home scripts/state/stride1 tx stakeibc change-validator-weight cosmoshub-4 HOST_VAL_ADDRESS_1 1 --from admin -y
build/strided --home scripts/state/stride1 tx stakeibc change-validator-weight cosmoshub-4 HOST_VAL_ADDRESS_2 49 --from admin -y

# LS and confirm delegation aligned with new weights
build/strided --home scripts/state/stride1 tx stakeibc liquid-stake 1000000 uatom --keyring-backend test --from admin -y --chain-id stride -y

# Update delegations (just submit this query and confirm the ICQ callback displays in the stride logs)
# Must be submitted in ICQ window
build/strided --home scripts/state/stride1 tx stakeibc update-delegation cosmoshub-4 HOST_VAL_ADDRESS_1 --from admin -y

#### MISC 
# If a channel closes, restore it with:
build/strided --home scripts/state/stride1 tx stakeibc restore-interchain-account cosmoshub-4 {DELEGATION | WITHDRAWAL | FEE | REDEMPTION} --from admin