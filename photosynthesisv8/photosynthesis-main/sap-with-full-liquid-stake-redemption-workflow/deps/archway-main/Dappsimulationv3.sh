#!/bin/bash

# Define your container ID
#CONTAINER_ID="22ec3aa32a5d"
PASSPHRASE="your_passphrase"
output=$(ps -aux | grep 'host-port 26457' | head -n 1 | awk '{ print $19 }' | cut -d- -f3)
OUTPUT1=$(build/archwayd --home dockernet/state/photo1 tx wasm store deps/archway-main/contrib/localnet/testdata/cw20_base.wasm --from pval3 --keyring-backend test --fees 500000uarch --gas 2000000 --chain-id localnet -y)

sleep 5
echo "$OUTPUT1"

echo "Script execution completed."

txhash=$(echo "$OUTPUT1" | grep -oP 'txhash: \K.*')



echo "$txhash"
txhash=$(echo $txhash | tr -dc '[:xdigit:]')
# Execute the command and retrieve the output
sleep 4
string=$(build/archwayd --home dockernet/state/photo1 q tx $txhash)

sleep 2
echo "$string"


sleep 2
# Add keys and capture the output (automate passphrase input)
KEY_OUTPUT1=$(echo -e "$PASSPHRASE\n$PASSPHRASE" | build/archwayd  --home dockernet/state/photo1 keys add l111  --keyring-backend=test)

sleep 2
echo "$KEY_OUTPUT1"

KEY_OUTPUT2=$(echo -e "$PASSPHRASE\n$PASSPHRASE" | build/archwayd --home dockernet/state/photo1 keys add l211  --keyring-backend=test)

sleep 2
echo "$KEY_OUTPUT2"

KEY_OUTPUT3=$(echo -e "$PASSPHRASE\n$PASSPHRASE" | build/archwayd --home dockernet/state/photo1 keys add l311 --keyring-backend=test)

sleep 2
echo "$KEY_OUTPUT3"

KEY_OUTPUT4=$(echo -e "$PASSPHRASE\n$PASSPHRASE" | build/archwayd --home dockernet/state/photo1 keys add l411  --keyring-backend=test)

sleep 2
echo "$KEY_OUTPUT4"

# Extract specific information (like addresses) from the output.
REWARD_ADDRESS=$(echo "$KEY_OUTPUT1" | grep "address:" | awk '{print $2}')

MINTER_ADDRESS=$(echo "$KEY_OUTPUT2" | grep "address:" | awk '{print $2}')

LIQUIDITY_ADDRESS=$(echo "$KEY_OUTPUT3" | grep "address:" | awk '{print $2}')

REDEMPTION_ADDRESS=$(echo "$KEY_OUTPUT4" | grep "address:" | awk '{print $2}')

# Instantiate
OUTPUT=$(build/archwayd tx --home dockernet/state/photo1 wasm instantiate 1 "{\"reward_address\": \"$REWARD_ADDRESS\", \"instantiation_request\": \"e30=\", \"gas_rebate_to_user\": false, \"collect_premium\": false, \"premium_percentage_charged\": 0, \"name\": \"aaaaa\", \"symbol\": \"TOK\", \"decimals\": 6, \"initial_balances\": [], \"mint\": {\"minter\": \"$MINTER_ADDRESS\"}}" --label test --from pval3  --keyring-backend=test --chain-id localnet --fees 30000uarch --no-admin -y)

sleep 5
echo "$OUTPUT"

echo "Script execution completed."

txhash=$(echo "$OUTPUT" | grep -oP 'txhash: \K.*')



echo "$txhash"
txhash=$(echo $txhash | tr -dc '[:xdigit:]')
# Execute the command and retrieve the output
sleep 4
string=$(build/archwayd q tx --home dockernet/state/photo1 $txhash)

sleep 2
echo "$string"

# Pattern to search
pattern="archway"

# Extract sender
sender=$(echo "$string" | grep -oP 'sender:\s*\Karchway[\w]*')

# Extract reward address
reward_address=$(echo "$string" | grep -oP 'reward_address:\s*\Karchway[\w]*')

# Extract contract address
contract_address=$(echo "$string" | grep -oP '"_contract_address","value":"\Karchway[\w]*')

# Pattern to search
pattern="archway"

# Check if pattern exists in each string
for variable in "$sender" "$reward_address" "$contract_address"
do
  if [[ $variable == *"$pattern"* ]]; then
    echo "Pattern found in $variable"
  else
    echo "Pattern not found in $variable"
  fi
done


# Print the variables
echo "Sender: $sender"
echo "Reward Address: $reward_address"
echo "Contract Address: $contract_address"
echo "Liquidity Provider Address: $LIQUIDITY_ADDRESS"

sleep 4

OUTPUT=$(build/archwayd  tx --home dockernet/state/photo1 rewards set-contract-metadata "$contract_address" --contract-address "$contract_address" --owner-address $sender --rewards-address $reward_address --airdrop-recipient-address archway1wmuuy0eqvhq5s3j9e80p8judf55c8v8mgwfytr --airdrop-vesting-period 6000 --archway-reward-funds-transfer-address archway1gnvac03v6xgtz3vt00p25j2nq28j9c55jlfntt  --liquid-stake-interval 1 --liquidity-provider-address archway1wmuuy0eqvhq5s3j9e80p8judf55c8v8mgwfytr --liquidity-provider-address "$LIQUIDITY_ADDRESS" --liquidity-provider-commission 2 --liquidity-token-address archway1smd403gckfc4m3upzxfuwxkree5lr9854u4un9 --maximum-threshold 4 --minimum-reward-amount 100 --redemption-address "$REDEMPTION_ADDRESS" --redemption-interval 1 --redemption-interval-threshold 1 --redemption-rate-threshold 1 --rewards-withdrawal-interval 1  --from pval3 --keyring-backend test --chain-id localnet --fees 30000uarch --gas 205000 -y)
sleep 5
echo "$OUTPUT"

echo "Script execution completed."

#txhash=$(echo "$OUTPUT" | grep -oP 'txhash: \K.*')



#echo "$txhash"
#txhash=$(echo $txhash | tr -dc '[:xdigit:]')
## Execute the command and retrieve the output
#sleep 4
#string=$(build/archwayd q tx --home dockernet/state/photo1 $txhash)

#sleep 2
#echo "$string"

#sleep 4
