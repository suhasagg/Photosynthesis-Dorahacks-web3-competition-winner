#!/bin/bash

# Get the directory containing the script
script_dir=$(dirname "$(dirname "$(realpath "$0")")")
first_two_levels=$(echo "$script_dir" | cut -d'/' -f1-3)

binaryarchwayd="$first_two_levels/build/archwayd"
binarystrided="$first_two_levels/build/strided"
statearchwayd="$first_two_levels/dockernet/state/photo1"
statestrided="$first_two_levels/dockernet/state/stride1"


mkdir -p ~/bin
# Copy the strided binary
echo "strided binary"
cp $binarystrided ~/bin/strided
echo "archwayd binary"
cp $binaryarchwayd ~/bin/archwayd
# Give it execution permissions

chmod +x ~/bin/strided
chmod +x ~/bin/archwayd

# Add ~/bin to PATH only if it is not already in
if [[ ":$PATH:" != *":$HOME/bin:"* ]]; then
    echo 'export PATH="$HOME/bin:$PATH"' >> /home/photo/.bashrc
    source /home/photo/.bashrc
fi

FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquidstakeparameters"

if [ -f "$FILE" ]; then
    echo "$FILE exists."

    AMOUNT=0
    line_number=1

    while IFS= read -r line
    do
        if [ "$((line_number % 2))" -eq 1 ]; then
            AMOUNT=$((AMOUNT + line))
        fi
        line_number=$((line_number + 1))
    done < "$FILE"

    echo "Liquid stake total amount: $AMOUNT"
else
    echo "$FILE does not exist."
    exit 0
fi


FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquidstakeparameters"
OUTPUT_FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/distributionepoch"

if [ -f "$FILE" ]; then
    last_line=$(tail -n 1 "$FILE")
    echo "Last line of file is: $last_line"

    echo "$last_line" >> "$OUTPUT_FILE"
    echo "Last line written to $OUTPUT_FILE"
else
    echo "$FILE does not exist."
fi
truncate -s 0 "$FILE"

sleep 5

echo "archwayd ibc transfer"
# run the archwayd command
#archwayd --home $statearchwayd tx ibc-transfer transfer transfer channel-0 --node http://photo1:26657 --chain-id localnet stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 400uarch --from pval2 -y
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx ibc-transfer transfer transfer channel-0 --chain-id localnet stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 ${AMOUNT}uarch --from pval2 --keyring-backend test -y --fees 30000uarch"
eval "$CMD"

# Delay for 5 seconds
sleep 5

echo "strided track balance"
# Run the first strided command
/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 q bank balances --chain-id STRIDE stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8

# Delay for another 5 seconds
sleep 5

echo "strided liquid stake"
# Run the second strided command
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 tx stakeibc liquid-stake ${AMOUNT} uarch --keyring-backend test --from admin --chain-id STRIDE -y"
eval "$CMD"
sleep 5

echo "strided track balance"
# Run the first strided command
OUTPUT1=$(/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 q bank balances --chain-id STRIDE stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8)

sleep 5
echo "$OUTPUT1"

while IFS= read -r line
do
  # If line contains 'denom: stuarch'
  if [ "$(echo $line | grep 'denom: stuarch')" ]; then
    # Get the amount from the previous line
    amount="${prev_line#*: \"}"
    # Remove trailing characters starting from '"'
    amount="${amount%\"}"
    echo "Amount: $amount"
    break
  fi
  # If line contains 'amount:'
  if [ "$(echo $line | grep 'amount:')" ]; then
    prev_line="$line"
  fi
done << EOF
$OUTPUT1
EOF

echo "liquid_token_amount"
echo "$amount"
echo "Script execution completed."

#rm "/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquidstakeparameters"
amount=$((amount - 30000))
echo "strided ibc liquid tokens transfer to Archway"
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 tx ibc-transfer transfer transfer channel-1 --chain-id STRIDE archway17zyrf6u6fltey79rz4h02c6ydmcpctyn3nf9vf ${amount}stuarch --from admin --keyring-backend test -y --fees 30000stuarch"
OUTPUT1=$(eval $CMD)
echo "$OUTPUT1"
sleep 5

echo "Script execution completed."

txhash=$(echo "$OUTPUT1" | grep -oP 'txhash: \K.*')

echo "$txhash"
txhash=$(echo $txhash | tr -dc '[:xdigit:]')
# Execute the command and retrieve the output
sleep 4
string=$(build/archwayd q tx --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 $txhash)

sleep 2
echo "$string"


# Delay for 5 seconds
sleep 5

echo "archway liquidity account balance"
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet archway17zyrf6u6fltey79rz4h02c6ydmcpctyn3nf9vf"
eval "$CMD"

# Delay for 5 seconds
sleep 5


