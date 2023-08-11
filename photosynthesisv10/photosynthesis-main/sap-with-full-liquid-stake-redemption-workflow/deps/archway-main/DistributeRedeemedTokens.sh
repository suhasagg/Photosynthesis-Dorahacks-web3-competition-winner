#!/bin/bash

LiquidityFILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquiditydistributiontodapps"
RedemptionFILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redemptiondistributionAmountforDapps"
RedemptionData="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redemptiondataforDapps"
declare -A tokensForContract
declare -A redemptionAddressforContract

# Read the RedemptionData file line by line
while IFS=':' read -r redemptionAddress contractAddress; do
    # Assign the redemption address to the contract address in the associative array
    redemptionAddressforContract["$contractAddress"]="$redemptionAddress"
done < "$RedemptionData"

# Print the redemption addresses for each contract address
for contract in "${!redemptionAddressforContract[@]}"; do
    echo "Contract $contract has redemption address: ${redemptionAddressforContract[$contract]}"
done

totalTokens=0

# Read the file line by line
while IFS=',' read -r part1 part2 part3 part4; do

    case $part4 in
        ''|*[!0-9.]*)
            echo "part4 is not numeric: $part4"
            continue ;;
        0|0.0|0.00)
            echo "part4 should not be equal to zero"
            continue ;;
    esac

    part4=$(echo $part4 | cut -d '.' -f 1) # Considering only the integral part

    # Sum tokens for each contractAddress
    if [ -z "${tokensForContract[$part2]}" ]; then
        tokensForContract["$part2"]=$part4
    else
        tokensForContract["$part2"]=$(( tokensForContract["$part2"] + part4 ))
    fi

    totalTokens=$(( totalTokens + part4 ))

done < "$LiquidityFILE"

# Print the contents of the map
for contract in "${!tokensForContract[@]}"; do
    echo "Contract $contract has tokens: ${tokensForContract[$contract]}"
done

echo "Total Tokens: $totalTokens"

CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m --chain-id localnet"
OUTPUT1=$(eval $CMD)
echo "$OUTPUT1"
sleep 2

   # start by declaring prev_line as empty
prev_line=""

    # parse OUTPUT1 line by line
while IFS= read -r line
do

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
      echo "total redemption amount: $total_redemption_amount"
  fi

  if [ "$(echo $line | grep 'denom: ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145')" ]; then
          # Get the amount from the previous line
      amount="${prev_line#*: \"}"
          # Remove trailing characters starting from '"'
      cumulativeamountstuarchibc="${amount%\"}"
      echo "Liquid Amount: $cumulativeamountstuarchibc"
  fi

  done << EOF
  $OUTPUT1
EOF

# Calculate the ratio for each contractAddress
for contract in "${!tokensForContract[@]}"; do
    ratio=$(echo "scale=2; ${tokensForContract[$contract]}/$totalTokens" | bc)
    echo "Contract $contract has a ratio of $ratio"

    # Calculate amount based on the ratio
    amount=$(echo "scale=2; $ratio * $total_redemption_amount" | bc | cut -d '.' -f 1)

    CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd tx --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 bank send archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m ${redemptionAddressforContract[$contract]} ${amount}uarch --from pval2 --keyring-backend=test --chain-id localnet --fees 17000uarch -y"

    echo "Executing: $CMD"
    OUTPUT1=$(eval $CMD)
    echo "$OUTPUT1"
    echo "$contract,$amount" >> "$RedemptionFILE"
    sleep 5
done

truncate -s 0 "$LiquidityFILE"
