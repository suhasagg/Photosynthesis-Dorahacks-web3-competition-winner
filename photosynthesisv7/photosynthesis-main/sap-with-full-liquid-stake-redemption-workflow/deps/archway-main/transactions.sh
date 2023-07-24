echo "dapp transactions"

# run the archwayd command
#archwayd --home $statearchwayd tx ibc-transfer transfer transfer channel-0 --node http://photo1:26657 --chain-id localnet stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 400uarch --from pval2 -y
OUTPUT=$(build/archwayd --home dockernet/state/photo1  tx wasm execute "$contract_address" '{"increase_allowance": {"spender": "'$reward_address'", "amount": "1"}}' --from pval1 --chain-id localnet --node http://$output:26657 --keyring-backend=test --gas 205000 --gas-prices 0.01uarch --gas-adjustment 1.2 -y)

sleep 5
echo "$OUTPUT"

echo "Script execution completed."

txhash=$(echo "$OUTPUT" | grep -oP 'txhash: \K.*')


echo "$txhash"
txhash=$(echo $txhash | tr -dc '[:xdigit:]')
# Execute the command and retrieve the output
sleep 4
string=$(build/archwayd --home dockernet/state/photo1 q tx $txhash)

sleep 2
echo "$string"

