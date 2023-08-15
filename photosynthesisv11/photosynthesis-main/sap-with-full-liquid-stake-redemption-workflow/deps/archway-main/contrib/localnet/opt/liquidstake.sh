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
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquidstakeparameters"

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

    echo "Liquid stake total amount: $AMOUNT" | jq -R '{"message": .}'
else
    echo "$FILE does not exist."
    exit 0
fi
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'


FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquidstakeparameters"
OUTPUT_FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/distributionepoch"
BACKUP_FILE="/tmp/backup_liquidstakeparameters"
TEMP_FILE="/tmp/temp_combined_file"

if [ -f "$FILE" ]; then
    last_line=$(tail -n 1 "$FILE")
    echo "Last line of file is: $last_line" | jq -R '{"message": .}'

    echo "$last_line" >> "$OUTPUT_FILE" | jq -R '{"message": .}'
    echo "Last line written to $OUTPUT_FILE" | jq -R '{"message": .}'
else
    echo "$FILE does not exist." | jq -R '{"message": .}'
fi

# Backup FILE to BACKUP_FILE
if [ -f "$FILE" ]; then
    cp "$FILE" "$BACKUP_FILE"
else
    echo "$FILE does not exist."
    exit 1
fi
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

# Truncate FILE
truncate -s 0 "$FILE"

sleep 5

echo "archwayd ibc transfer" | jq -R '{"message": .}'
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx ibc-transfer transfer transfer channel-0 --chain-id localnet stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 ${AMOUNT}uarch --from pval2 --keyring-backend test -y --fees 30000uarch"

# Initialize a variable to track if the command was successful
SUCCESS=0

# Loop to retry in case of sequence mismatch
for i in {1..10}; do
    echo "Attempt $i: Executing '$CMD'" | jq -R '{"message": .}'
    # Print the timestamp in a specific format
    echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
    echo "$CMD" | jq -R '{"message": .}'
    OUTPUT4=$(eval "$CMD")
    # Convert the YAML output to JSON using yq
    json_output=$(echo "$OUTPUT4" | yq eval -j -)

    # Embed the json_output in a new JSON structure using jq
    final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

    echo $final_json


    if [[ "$OUTPUT4" != *"account sequence mismatch"* ]]; then
        # Successful execution; break out of loop
        echo "Command executed successfully." | jq -R '{"message": .}'
        SUCCESS=1
        break
    else
        # Error encountered; sleep for 5 seconds before retry
        echo "Account sequence mismatch detected. Retrying in 5 seconds..." | jq -R '{"message": .}'
        sleep 5
    fi
done
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

# If the command was not successful after all retries
if [ "$SUCCESS" -eq 0 ]; then
    echo "Failed after max retries. Restoring original content from backup..." | jq -R '{"message": .}'

    # Create a temporary file that starts with the content of BACKUP_FILE and is followed by FILE's content
    cat "$BACKUP_FILE" > "$TEMP_FILE"
    cat "$FILE" >> "$TEMP_FILE"

    # Replace FILE with TEMP_FILE
    mv "$TEMP_FILE" "$FILE"
    rm "$BACKUP_FILE"  # Remove the backup after restoring
    echo "Exiting script after restoring content." | jq -R '{"message": .}'
    exit 1
fi
# Delay for 5 seconds
sleep 5
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

echo "strided track balance" | jq -R '{"message": .}'
# Run the first strided command
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 q bank balances --chain-id STRIDE stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8"
echo "$CMD" | jq -R '{"message": .}'
OUTPUT2=$(eval $CMD)
 # Convert the YAML output to JSON using yq
json_output=$(echo "$OUTPUT2" | yq eval -j -)

    # Embed the json_output in a new JSON structure using jq
final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

echo $final_json
sleep 5
# Delay for another 5 seconds
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

echo "strided liquid stake" | jq -R '{"message": .}'
# Define a success flag
SUCCESS=0

for i in {1..5}; do
    echo "strided liquid stake" | jq -R '{"message": .}'
    CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 tx stakeibc liquid-stake ${AMOUNT} uarch --keyring-backend test --from admin --chain-id STRIDE -y"
    # Print the timestamp in a specific format
    echo "$CMD" | jq -R '{"message": .}'
    echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

    OUTPUT2=$(eval "$CMD")
    # Convert the YAML output to JSON using yq
    json_output=$(echo "$OUTPUT2" | yq eval -j -)

       # Embed the json_output in a new JSON structure using jq
    final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

    echo $final_json
    txhash=$(echo "$OUTPUT2" | grep -oP 'txhash: \K.*')

    # Check if txhash extraction was successful
    if [ -z "$txhash" ]; then
        echo "Error: Failed to extract txhash." | jq -R '{"message": .}'
        sleep 10
        continue
    fi

    echo "Transaction hash: $txhash" | jq -R '{"message": .}'
    txhash=$(echo "$txhash" | tr -dc '[:xdigit:]')
    sleep 4
    string=$(/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 q tx "$txhash")
    echo $string | jq -R '{"message": .}'
    if [[ "$string" != *"raw_log: 'failed to execute message"* ]]; then
        # Successful execution
        SUCCESS=1
        break
    else
        # Error encountered; sleep for 10 seconds before retry
        echo "Failed to execute message detected. Retrying in 30 seconds..." | jq -R '{"message": .}'
        sleep 30
    fi
done
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

if [ "$SUCCESS" -eq 0 ]; then
    echo "Failed after max retries." | jq -R '{"message": .}'
fi

sleep 5

echo "strided track balance" | jq -R '{"message": .}'
# Run the first strided command
OUTPUT1=$(/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 q bank balances --chain-id STRIDE stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8)
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
 # Convert the YAML output to JSON using yq
json_output=$(echo "$OUTPUT1" | yq eval -j -)

    # Embed the json_output in a new JSON structure using jq
final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

echo $final_json
sleep 5
echo "$OUTPUT1" | jq -R '{"message": .}'

while IFS= read -r line
do
  # If line contains 'denom: stuarch'
  if [ "$(echo $line | grep 'denom: stuarch')" ]; then
    # Get the amount from the previous line
    amount="${prev_line#*: \"}"
    # Remove trailing characters starting from '"'
    amount="${amount%\"}"
    echo "Amount: $amount" | jq -R '{"message": .}'
    break
  fi
  # If line contains 'amount:'
  if [ "$(echo $line | grep 'amount:')" ]; then
    prev_line="$line"
  fi
done << EOF
$OUTPUT1
EOF
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

echo "liquid_token_amount" | jq -R '{"message": .}'
echo "$amount" | jq -R '{"message": .}'
echo "Script execution completed." | jq -R '{"message": .}'
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

#rm "/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquidstakeparameters"
amount=$((amount - 30000))
echo "strided ibc liquid tokens transfer to Archway" | jq -R '{"message": .}'
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 tx ibc-transfer transfer transfer channel-1 --chain-id STRIDE archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m ${amount}stuarch --from admin --keyring-backend test -y --fees 30000stuarch"
echo "$CMD" | jq -R '{"message": .}'
OUTPUT1=$(eval $CMD)
 # Convert the YAML output to JSON using yq
json_output=$(echo "$OUTPUT1" | yq eval -j -)

    # Embed the json_output in a new JSON structure using jq
final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

echo $final_json
echo "$OUTPUT1" | jq -R '{"message": .}'
sleep 5

echo "Script execution completed." | jq -R '{"message": .}'

txhash=$(echo "$OUTPUT1" | grep -oP 'txhash: \K.*')

echo "$txhash" | jq -R '{"message": .}'
txhash=$(echo $txhash | tr -dc '[:xdigit:]')
# Execute the command and retrieve the output
sleep 4
string=$(build/archwayd q tx --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 $txhash)

sleep 2
echo "$string" | jq -R '{"message": .}'
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'


# Delay for 5 seconds
sleep 5

echo "archway liquidity account balance" | jq -R '{"message": .}'
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"
echo "$CMD" | jq -R '{"message": .}'
OUTPUT1=$(eval $CMD)
 # Convert the YAML output to JSON using yq
json_output=$(echo "$OUTPUT1" | yq eval -j -)

    # Embed the json_output in a new JSON structure using jq
final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

echo $final_json


# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

# Delay for 5 seconds
sleep 5


