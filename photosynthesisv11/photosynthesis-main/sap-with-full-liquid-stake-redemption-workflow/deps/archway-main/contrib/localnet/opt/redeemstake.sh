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
case ":$PATH:" in
    *":$HOME/bin:"*) ;;
    *) echo 'export PATH="$HOME/bin:$PATH"' >> /home/photo/.bashrc
       . /home/photo/.bashrc ;;
esac
# Variable to hold the total redemption amount
total_redemption_amount=0

# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'


FILE1="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/enableredeemstake"

if [ -f "$FILE1" ]; then
    echo "$FILE1 exists."
else
    echo "$FILE1 does not exist."
    exit 0
fi

FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redeemliquidityamountforDapps"
LiquidityFILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquiditydistributionforDapps"
# Read the file line by line
while IFS=',' read -r part1 part2 part3

do
    # Construct and execute the command
  CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances ${part1} --chain-id localnet"
  echo "$CMD" | jq -R '{"message": .}'
  OUTPUT1=$(eval $CMD)
   # Convert the YAML output to JSON using yq
  json_output=$(echo "$OUTPUT1" | yq eval -j -)

      # Embed the json_output in a new JSON structure using jq
  final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

  echo $final_json
  echo "$OUTPUT1" | jq -R '{"message": .}'
  sleep 2
  # Print the timestamp in a specific format
  echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

   # start by declaring prev_line as empty
  prev_line=""

    # parse OUTPUT1 line by line
  while IFS= read -r line
  do

      # If line contains 'amount:', store it in prev_line
    if [ "$(echo $line | grep 'amount:')" ]; then
        prev_line="$line"
    fi

    # If line contains 'denom: stuarch'
    if [ "$(echo $line | grep 'denom: stuarch')" ]; then
        # Get the amount from the previous line
        amount="${prev_line#*: \"}"
        # Remove trailing characters starting from '"'
        cumulativeamountstuarch="${amount%\"}"
        total_redemption_amount=$(echo "$total_redemption_amount + $cumulativeamountstuarch" | bc)
        echo "total redemption amount: $total_redemption_amount" | jq -R '{"message": .}'
        echo "Amount: $cumulativeamountstuarch" | jq -R '{"message": .}'
        CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx bank send ${part1} ${part2} ${cumulativeamountstuarch}stuarch --from l311 --keyring-backend=test --chain-id localnet --fees 17000arch -y"
        echo "$CMD" | jq -R '{"message": .}'
        OUTPUT2=$(eval $CMD)
         # Convert the YAML output to JSON using yq
        json_output=$(echo "$OUTPUT2" | yq eval -j -)

            # Embed the json_output in a new JSON structure using jq
        final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

        echo $final_json
        echo "$OUTPUT2" | jq -R '{"message": .}'
        echo "$part1,$part2,$part3,$cumulativeamountstuarch" >> $LiquidityFILE
        sleep 5
        break
    fi

    if [ "$(echo $line | grep 'denom: ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145')" ]; then
          # Get the amount from the previous line
        amount="${prev_line#*: \"}"
          # Remove trailing characters starting from '"'
        cumulativeamountibc="${amount%\"}"
        total_redemption_amount=$(echo "$total_redemption_amount + $cumulativeamountibc" | bc)
        echo "total redemption amount: $total_redemption_amount" | jq -R '{"message": .}'
        echo "Amount: $cumulativeamountibc" | jq -R '{"message": .}'
        CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1  tx bank send  ${part1} ${part2} ${cumulativeamountibc}ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145 --from l311 --keyring-backend=test --chain-id localnet --fees 17000uarch -y"
        echo "$CMD" | jq -R '{"message": .}'
        OUTPUT3=$(eval $CMD)
         # Convert the YAML output to JSON using yq
        json_output=$(echo "$OUTPUT3" | yq eval -j -)

            # Embed the json_output in a new JSON structure using jq
        final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

        echo $final_json
        echo "$OUTPUT3" | jq -R '{"message": .}'
        echo "$part1,$part2,$part3,$cumulativeamountibc" >> $LiquidityFILE
        sleep 5
        break
    fi
    # Print the timestamp in a specific format
    echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

    # Construct and execute the command
    # If line contains 'amount:'
   # if [ "$(echo $line | grep 'amount:')" ]; then
   #     prev_line="$line"
   # fi
  done << EOF
  $OUTPUT1
EOF
done < "$FILE"


sleep 2
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

echo "archway liquidity account balance" | jq -R '{"message": .}'
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"
echo "$CMD" | jq -R '{"message": .}'
OUTPUT4=$(eval $CMD)
 # Convert the YAML output to JSON using yq
json_output=$(echo "$OUTPUT4" | yq eval -j -)

    # Embed the json_output in a new JSON structure using jq
final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

echo $final_json
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

sleep 2
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

echo "archwayd ibc transfer | jq -R '{"message": .}'"
# run the archwayd command
#archwayd --home $statearchwayd tx ibc-transfer transfer transfer channel-0 --node http://photo1:26657 --chain-id localnet stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 400uarch --from pval2 -y
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx ibc-transfer transfer transfer channel-0 --chain-id localnet stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 ${total_redemption_amount}ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145 --from pval5 --keyring-backend test -y --fees 30000uarch"
echo "$CMD" | jq -R '{"message": .}'
OUTPUT4=$(eval $CMD)
 # Convert the YAML output to JSON using yq
json_output=$(echo "$OUTPUT4" | yq eval -j -)

    # Embed the json_output in a new JSON structure using jq
final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

echo $final_json
echo "$OUTPUT4" | jq -R '{"message": .}'
# Delay for another 5 seconds
sleep 4
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'


echo "strided redeem stake" | jq -R '{"message": .}'
# Run the first strided command
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 tx stakeibc redeem-stake ${total_redemption_amount} localnet archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m --from admin --keyring-backend test --chain-id STRIDE -y"
# Initialize a variable to track if the command was successful
SUCCESS=0

# Loop to retry in case of sequence mismatch
for i in {1..5}; do
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

# Delay for another 5 seconds
sleep 4
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

#echo "Number of successfully redeemed records" >> $LiquidityFILE

echo "archway balance uarch" | jq -R '{"message": .}'
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"
echo "$CMD" | jq -R '{"message": .}'
OUTPUT5=$(eval $CMD)
 # Convert the YAML output to JSON using yq
json_output=$(echo "$OUTPUT5" | yq eval -j -)

    # Embed the json_output in a new JSON structure using jq
final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

echo $final_json
echo "$OUTPUT5" | jq -R '{"message": .}'
sleep 4
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

total_arch_balance_amount1=0
total_arch_balance_amount2=0
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

while IFS= read -r line
do
  # If line contains 'denom: stuarch'
  if [ "$(echo $line | grep 'denom: uarch')" ]; then
    # Get the amount from the previous line
    amount="${prev_line#*: \"}"
    # Remove trailing characters starting from '"'
    amount="${amount%\"}"
    total_arch_balance_amount1=amount
    echo "Amount before claim: $amount" | jq -R '{"message": .}'
    break
  fi
  # If line contains 'amount:'
  if [ "$(echo $line | grep 'amount:')" ]; then
    prev_line="$line"
  fi
done << EOF
$OUTPUT5
EOF
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'


echo "list user redemption  record" | jq -R '{"message": .}'
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 q records list-user-redemption-record -o=json"
echo "$CMD" | jq -R '{"message": .}'
output=$(eval $CMD)
 # Convert the YAML output to JSON using yq
json_output=$(echo "$output" | yq eval -j -)

    # Embed the json_output in a new JSON structure using jq
final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

echo $final_json
echo "$output" | jq -R '{"message": .}'

sleep 4

# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

echo "list epoch unbonding record" | jq -R '{"message": .}'
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 q records list-epoch-unbonding-record -o=json"
echo "$CMD" | jq -R '{"message": .}'
output=$(eval $CMD)
 # Convert the YAML output to JSON using yq
json_output=$(echo "$output" | yq eval -j -)

    # Embed the json_output in a new JSON structure using jq
final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

echo $final_json
echo "$output" | jq -R '{"message": .}'

# Delay for another 5 seconds
sleep 4

# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

# Parse JSON from the variable
json=$(echo "$output" | jq -c '.')

claimed_records=0
stride_epoch=0
# Iterate over json array
echo "${json}" | jq -r '.epoch_unbonding_record[] | @base64' | while read epoch; do
    # Print the timestamp in a specific format
    echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

    # Decode each row
    _jq_epoch() {
     echo ${epoch} | base64 --decode | jq -r ${1}
    }

    echo "$(_jq_epoch '.host_zone_unbondings')" | jq -r '.[] | @base64' | while read row; do

        # Decode each row
        _jq() {
            echo ${row} | base64 --decode | jq -r ${1}
        }

        # Get status
        status=$(_jq '.status')

        # If status is CLAIMABLE
        if [ "$status" = "CLAIMABLE" ]; then

            # Get user_redemption_records array
            records=$(_jq '.user_redemption_records')

            # Iterate over records array
            echo "${records}" | jq -r '.[] | @base64' | while read record; do
                _jq_record() {
                    echo ${record} | base64 --decode
                }

                # Split record into three parts
                record_val=$(_jq_record)
                part1=$(echo $record_val | cut -d'.' -f1)
                part2=$(echo $record_val | cut -d'.' -f2)
                part3=$(echo $record_val | cut -d'.' -f3)
                stride_epoch = part2
                echo "Hostzone: $part1" | jq -R '{"message": .}'
                echo "Epoch: $part2" | jq -R '{"message": .}'
                echo "Stride Address: $part3" | jq -R '{"message": .}'
                echo "claim undelegated tokens" | jq -R '{"message": .}'
                CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 tx stakeibc claim-undelegated-tokens ${part1} ${part2} ${part3} --from admin --keyring-backend test --chain-id STRIDE -y"
                echo "$CMD" | jq -R '{"message": .}'
                output=$(eval $CMD)
                 # Convert the YAML output to JSON using yq
                json_output=$(echo "$output" | yq eval -j -)

                    # Embed the json_output in a new JSON structure using jq
                final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

                echo $final_json
                echo "$output" | jq -R '{"message": .}'
                # Delay for another 5 seconds
                #echo "Number of successfully claimed records" >> /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquiditydistributionforDapps
                sleep 4
            done
        fi
    done
done
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

sleep 10

echo "archway redeemed uarch" | jq -R '{"message": .}'
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"
echo "$CMD" | jq -R '{"message": .}'
OUTPUT5=$(eval $CMD)
 # Convert the YAML output to JSON using yq
 # json_output=$(echo "$OUTPUT4" | yq eval -j -)
json_output=$(echo "$OUTPUT5" | yq eval -j -)
    # Embed the json_output in a new JSON structure using jq
final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')

echo $final_json
echo "$OUTPUT5" | jq -R '{"message": .}'
sleep 4
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

while IFS= read -r line
do
  # If line contains 'denom: stuarch'
  if [ "$(echo $line | grep 'denom: uarch')" ]; then
    # Get the amount from the previous line
    amount="${prev_line#*: \"}"
    # Remove trailing characters starting from '"'
    amount="${amount%\"}"
    total_arch_balance_amount2=amount
    echo "Amount after claim: $amount" | jq -R '{"message": .}'
    break
  fi
  # If line contains 'amount:'
  if [ "$(echo $line | grep 'amount:')" ]; then
    prev_line="$line"
  fi
done << EOF
$OUTPUT5
EOF
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'


FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/enableredeemstake"
OUTPUT_FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redeemepoch"

if [ -f "$FILE" ]; then
    last_line=$(tail -n 1 "$FILE")
    echo "Last line of file is: $last_line" | jq -R '{"message": .}'

    echo "$last_line" >> "$OUTPUT_FILE"
    echo "Last line written to $OUTPUT_FILE" | jq -R '{"message": .}'
else
    echo "$FILE does not exist." | jq -R '{"message": .}'
fi
# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

rm "/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/enableredeemstake"
