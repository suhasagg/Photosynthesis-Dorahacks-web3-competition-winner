
# run the archwayd command
#archwayd --home $statearchwayd tx ibc-transfer transfer transfer channel-0 --node http://photo1:26657 --chain-id localnet stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 400uarch --from pval2 -
OUTPUT=$(/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1  tx wasm execute archway14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sy85n2u '{"increase_allowance": {"spender": "archway1f7js3pc9fs50hd6ttv3qwmt9fldh0xlthx6gkq", "amount": "1"}}' --from pval2 --chain-id localnet --keyring-backend=test --gas 205000 --fees 20000uarch -y)


sleep 5
echo "$OUTPUT"

echo "Script execution completed."

txhash=$(echo "$OUTPUT" | grep -oP 'txhash: \K.*')


echo "$txhash"
txhash=$(echo $txhash | tr -dc '[:xdigit:]')
# Execute the command and retrieve the output
sleep 4
string=$(/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q tx $txhash)

sleep 2
echo "$string"

# run the archwayd command
#archwayd --home $statearchwayd tx ibc-transfer transfer transfer channel-0 --node http://photo1:26657 --chain-id localnet stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 400uarch --from pval2 -
OUTPUT=$(/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1  tx wasm execute archway14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sy85n2u '{"increase_allowance": {"spender": "archway1f7js3pc9fs50hd6ttv3qwmt9fldh0xlthx6gkq", "amount": "1"}}' --from pval2 --chain-id localnet --keyring-backend=test --gas 205000 --fees 20000uarch -y)


sleep 5
echo "$OUTPUT"

echo "Script execution completed."

txhash=$(echo "$OUTPUT" | grep -oP 'txhash: \K.*')


echo "$txhash"
txhash=$(echo $txhash | tr -dc '[:xdigit:]')
# Execute the command and retrieve the output
sleep 4
string=$(/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q tx $txhash)

sleep 2
echo "$string"

