#!/bin/bash

# Define your container ID
CONTAINER_ID="579fa69e620d"
PASSPHRASE="your_passphrase"

OUTPUT1=$(docker exec -it $CONTAINER_ID archwayd tx wasm store testdata/cw20_base.wasm --from fd --keyring-backend test --gas 10000000 --gas-prices 0.01stake --gas-adjustment 1.4 --chain-id localnet -y)

sleep 2
# Add keys and capture the output (automate passphrase input)
KEY_OUTPUT1=$(echo -e "$PASSPHRASE\n$PASSPHRASE" | docker exec -i $CONTAINER_ID archwayd keys add l1 --keyring-backend test )

sleep 2
echo "$KEY_OUTPUT1"

KEY_OUTPUT2=$(echo -e "$PASSPHRASE\n$PASSPHRASE" | docker exec -i $CONTAINER_ID archwayd keys add l2 --keyring-backend test)

sleep 2
echo "$KEY_OUTPUT2"

KEY_OUTPUT3=$(echo -e "$PASSPHRASE\n$PASSPHRASE" | docker exec -i $CONTAINER_ID archwayd keys add l3 --keyring-backend test)

sleep 2
echo "$KEY_OUTPUT3"

KEY_OUTPUT4=$(echo -e "$PASSPHRASE\n$PASSPHRASE" | docker exec -i $CONTAINER_ID archwayd keys add l4 --keyring-backend test)

sleep 2
echo "$KEY_OUTPUT4"

# Extract specific information (like addresses) from the output.
REWARD_ADDRESS=$(echo "$KEY_OUTPUT1" | grep "address:" | awk '{print $2}')

MINTER_ADDRESS=$(echo "$KEY_OUTPUT2" | grep "address:" | awk '{print $2}')


# Instantiate
OUTPUT=$(docker exec -it $CONTAINER_ID archwayd tx wasm instantiate 1 "{\"reward_address\": \"$REWARD_ADDRESS\", \"instantiation_request\": \"e30=\", \"gas_rebate_to_user\": false, \"collect_premium\": false, \"premium_percentage_charged\": 0, \"name\": \"aaaaa\", \"symbol\": \"TOK\", \"decimals\": 6, \"initial_balances\": [], \"mint\": {\"minter\": \"$MINTER_ADDRESS\"}}" --label test --from fd --keyring-backend test --chain-id localnet --fees 10000stake --no-admin -y)

sleep 5
echo "$OUTPUT" 

echo "Script execution completed."

txhash=$(echo "$OUTPUT" | grep -oP 'txhash: \K.*')



echo "$txhash"
txhash=$(echo $txhash | tr -dc '[:xdigit:]')
# Execute the command and retrieve the output
sleep 4
string=$(docker exec -it $CONTAINER_ID archwayd q tx $txhash)

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

sleep 2 

docker exec -t $CONTAINER_ID archwayd tx rewards set-contract-metadata "$contract_address" --contract-address "$contract_address" --from fd --keyring-backend test --chain-id localnet --owner-address $sender --rewards-address $reward_address --airdrop-recipient-address archway1wmuuy0eqvhq5s3j9e80p8judf55c8v8mgwfytr --airdrop-vesting-period 6000 --archway-reward-funds-transfer-address archway1gnvac03v6xgtz3vt00p25j2nq28j9c55jlfntt  --liquid-stake-interval 1 --liquidity-provider-address archway1wmuuy0eqvhq5s3j9e80p8judf55c8v8mgwfytr --liquidity-provider-commission 2 --liquidity-token-address archway1smd403gckfc4m3upzxfuwxkree5lr9854u4un9 --maximum-threshold 4 --minimum-reward-amount 100 --redemption-address archway18kpsdc76xg5884ey3qnesqtw8l9n06yw0u898p --redemption-interval 1 --redemption-interval-threshold 1 --redemption-rate-threshold 1 --rewards-withdrawal-interval 10 --fees 10000stake --gas 205000 -y

sleep 4

docker exec -it $CONTAINER_ID archwayd tx wasm execute "$contract_address" '{"increase_allowance": {"spender": "'$reward_address'", "amount": "1"}}' --from fd --chain-id localnet --keyring-backend test --gas 205000 --gas-prices 0.01stake --gas-adjustment 1.2 -y

sleep 4

docker exec -it $CONTAINER_ID archwayd tx wasm execute "$contract_address" '{"increase_allowance": {"spender": "'$reward_address'", "amount": "1"}}' --from fd --chain-id localnet --keyring-backend test --gas 205000 --gas-prices 0.01stake --gas-adjustment 1.2 -y

sleep 4 

docker exec -it $CONTAINER_ID archwayd tx wasm execute "$contract_address" '{"increase_allowance": {"spender": "'$reward_address'", "amount": "1"}}' --from fd --chain-id localnet --keyring-backend test --gas 205000 --gas-prices 0.01stake --gas-adjustment 1.2 -y

sleep 4

docker exec -it $CONTAINER_ID archwayd tx wasm execute "$contract_address" '{"increase_allowance": {"spender": "'$reward_address'", "amount": "1"}}' --from fd --chain-id localnet --keyring-backend test --gas 205000 --gas-prices 0.01stake --gas-adjustment 1.2 -y

sleep 4 
 
docker exec -it $CONTAINER_ID archwayd q rewards rewards-records "$reward_address"



