#!/bin/bash

LiquidityFILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquiditydistributiontodapps"
RedemptionFILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redemptiondistributionAmountforDapps"
RedemptionData="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redemptiondataforDapps"
declare -A tokensForContract
declare -A redemptionAddressforContract
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
# Read the RedemptionData file line by line
while IFS=':' read -r redemptionAddress contractAddress; do
    # Assign the redemption address to the contract address in the associative array
    redemptionAddressforContract["$contractAddress"]="$redemptionAddress"
done < "$RedemptionData"
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
# Print the redemption addresses for each contract address
for contract in "${!redemptionAddressforContract[@]}"; do
    echo "Contract $contract has redemption address: ${redemptionAddressforContract[$contract]}" | jq -R -c '{"message": .}'
done
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
totalTokens=0

# Read the file line by line
while IFS=',' read -r part1 part2 part3 part4; do

    case $part4 in
        ''|*[!0-9.]*)
            echo "part4 is not numeric: $part4" | jq -R -c '{"message": .}'
            continue ;;
        0|0.0|0.00)
            continue ;;
    esac
    echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
    part4=$(echo $part4 | cut -d '.' -f 1) # Considering only the integral part

    # Sum tokens for each contractAddress
    if [ -z "${tokensForContract[$part2]}" ]; then
        tokensForContract["$part2"]=$part4
    else
        tokensForContract["$part2"]=$(( tokensForContract["$part2"] + part4 ))
    fi

    totalTokens=$(( totalTokens + part4 ))
    echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
done < "$LiquidityFILE"
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
# Print the contents of the map
for contract in "${!tokensForContract[@]}"; do
    echo "Contract $contract has tokens: ${tokensForContract[$contract]}" | jq -R -c '{"message": .}'
done
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
echo "Total Tokens: $totalTokens" | jq -R -c '{"message": .}'
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m --chain-id localnet"
echo "$CMD" | jq -R -c '{"message": .}'
OUTPUT1=$(eval $CMD)
#echo "$OUTPUT" | jq -R -c '{"message": .}'
json_output=$(echo "$OUTPUT1" | yq eval -j -)
 # Embed the json_output in a new JSON structure using jq
#final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded' -c)

echo $json_output
#echo "$OUTPUT1" | jq -R -c '{"message": .}'
sleep 2
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
   # start by declaring prev_line as empty
prev_line=""

    # parse OUTPUT1 line by line
while IFS= read -r line
do
  echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
      # If line contains 'amount:', store it in prev_line
  if [ "$(echo $line | grep 'amount:')" ]; then
       prev_line="$line"
  fi

    # If line contains 'denom: uarch'
  if [ "$(echo $line | grep 'denom: uarch')" ]; then
        # Get the amount from the previous line
      amount="${prev_line#*: \"}"
        # Remove trailing characters starting from '"'
      cumulativeamountuarch="${amount%\"}"
      total_redemption_amount=$cumulativeamountuarch
      echo "total redemption amount: $total_redemption_amount" | jq -R -c '{"message": .}'
  fi

  if [ "$(echo $line | grep 'denom: ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145')" ]; then
          # Get the amount from the previous line
      amount="${prev_line#*: \"}"
          # Remove trailing characters starting from '"'
      cumulativeamountstuarchibc="${amount%\"}"
      echo "Liquid Amount: $cumulativeamountstuarchibc" | jq -R -c '{"message": .}'
  fi
  echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
  done << EOF
  $OUTPUT1
EOF

# Name of the log file
log_file="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redeemstake.log"

# Check if the log file exists
if [[ ! -f $log_file ]]; then
    echo "Error: $log_file does not exist."
    exit 1
fi

# Using tac to parse the file from bottom up
BEFORE_CLAIM_RAW=$(tac "$log_file" | grep -m1 "Amount before claim:" | awk '{print $NF}' | tr -d '"')
AFTER_CLAIM_RAW=$(tac "$log_file" | grep -m1 "Amount after claim:" | awk '{print $NF}' | tr -d '"')

# Strip non-numeric characters
BEFORE_CLAIM=$(echo "$BEFORE_CLAIM_RAW" | tr -cd '0-9')
AFTER_CLAIM=$(echo "$AFTER_CLAIM_RAW" | tr -cd '0-9')

# Debug: Print the values
echo "Debug: BEFORE_CLAIM = $BEFORE_CLAIM"
echo "Debug: AFTER_CLAIM = $AFTER_CLAIM"

# Compute the difference
DIFFERENCE=$(( AFTER_CLAIM - BEFORE_CLAIM ))

# Debug: Print the difference
echo "Debug: DIFFERENCE = $DIFFERENCE"

# Print the extracted and computed values
echo "Amount before claim: $BEFORE_CLAIM" | jq -R -c '{"message": .}'
echo "Amount after claim: $AFTER_CLAIM" | jq -R -c '{"message": .}'
echo "Difference between before claim and after claim: $DIFFERENCE" | jq -R -c '{"message": .}'

# Check if DIFFERENCE is zero
if [ $DIFFERENCE -eq 0 ]; then
    echo "No change between the amounts before and after the claim."
else
    echo "Difference detected."
fi

total_redemption_amount=$DIFFERENCE
echo "$total_redemption_amount" | jq -R -c '{"message": .}'
# Printing the result
echo "Total Redemption Amount: $total_redemption_amount" | jq -R -c '{"message": .}'
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
total_redemption_amount=$DIFFERENCE
echo "$total_redemption_amount" | jq -R -c '{"message": .}'
# Printing the result
echo "Total Redemption Amount: $total_redemption_amount" | jq -R -c '{"message": .}'
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'

if [ $total_redemption_amount -gt 0 ]; then
    # Calculate the ratio for each contractAddress
    for contract in "${!tokensForContract[@]}"; do
        echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
        ratio=$(echo "scale=2; ${tokensForContract[$contract]}/$totalTokens" | bc)
        echo "Contract $contract has a ratio of $ratio" | jq -R -c '{"message": .}'

        # Calculate amount based on the ratio
        amount=$(echo "scale=2; $ratio * $total_redemption_amount" | bc | cut -d '.' -f 1)

        CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd tx --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 bank send archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m ${redemptionAddressforContract[$contract]} ${amount}uarch --from pval2 --keyring-backend=test --chain-id localnet --fees 17000uarch -y"

        # Initialize a variable to track if the command was successful
        SUCCESS=0

        # Loop to retry in case of failure
        for i in {1..10}; do
            echo "Attempt $i: Executing '$CMD'" | jq -R -c '{"message": .}'
            echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
            OUTPUT1=$(eval $CMD)
            json_output=$(echo "$OUTPUT1" | yq eval -j -)

            # Check if the command was successful. 
            if [[ "$json_output" == *"success"* ]]; then
                echo "Command executed successfully." | jq -R -c '{"message": .}'
                SUCCESS=1
                break
            else
                # Error encountered; sleep for 5 seconds before retry
                echo "Error detected. Retrying in 15 seconds..." | jq -R -c '{"message": .}'
                sleep 15
            fi
        done

        if [ $SUCCESS -eq 0 ]; then
            echo "Failed to execute command after 10 attempts." | jq -R -c '{"message": .}'
            exit 1
        fi

        # Embed the json_output in a new JSON structure using jq
        #final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded' -c)

        echo $json_output
        #echo "$OUTPUT1" | jq -R -c '{"message": .}'
        echo "$contract,$amount" >> "$RedemptionFILE"
        sleep 5
    done
else
    echo "Total redemption amount is not greater than zero. Skipping transactions." | jq -R -c '{"message": .}'
fi

echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'

truncate -s 0 "$LiquidityFILE"
