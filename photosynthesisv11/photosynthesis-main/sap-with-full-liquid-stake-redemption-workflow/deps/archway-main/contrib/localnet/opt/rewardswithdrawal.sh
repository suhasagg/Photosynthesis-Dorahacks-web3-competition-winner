#!/bin/bash

declare -A contractToKeyName
key=""
address=""

# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'


while read -r line
do
    if echo "$line" | grep -q '^- name:'
    then
        key=$(echo "$line" | cut -d' ' -f3)
        echo "Read key: $key" | jq -R '{"message": .}'
    elif echo "$line" | grep -q '^address:'
    then
        address=$(echo "$line" | cut -d' ' -f2)
    fi

    if [ -n "$key" ] && [ -n "$address" ]
    then
        contractToKeyName[$address]=$key
        echo "Read key: $key for address: $address" | jq -R '{"message": .}'# Debugging line
        key=""
        address=""
    fi

done <<EOF
$(/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 keys list)
EOF

# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'


FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/rewardswithdrawaltodapps"
declare -A contractMap=()
declare -A contractrewardsMap=()
# First, populate the map with the total aggregated values per contract address
while IFS=',' read -r part1 part2 part3 part4
do
    totalAggregated=${contractMap[$part1]:-0} # get the total aggregated value for this contract, defaulting to 0 if not set
    totalAggregated=$((totalAggregated + part3))
    contractMap[$part1]="$totalAggregated"
    totalRewards=${contractrewardsMap[$part1]:-0} # get the total rewards for this contract, defaulting to 0 if not set
    totalRewards=$((totalRewards + part2))
    contractrewardsMap[$part1]="$totalRewards"
done < "$FILE"

# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'


# Now, go through                                                                                                                                                                                                                                                            the map and remove lines from the file for contract addresses with total aggregation > 30
for contract in "${!contractMap[@]}"; do
    if [ "${contractMap[$contract]}" -gt 30 ]; then
        echo "Total aggregate for contract $contract exceeded 30 ${contractMap[$contract]}. Removing lines..." | jq -R '{"message": .}'
        # remove lines for this contract address
        sed -i "/^$contract,/d" "$FILE"
    fi
done

# Print the timestamp in a specific format
echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'


for contractAddress in "${!contractMap[@]}"; do
    IFS="," read totalAggregated <<< "${contractMap[$contractAddress]}"
    keyName=${contractToKeyName[$contractAddress]}
    part2=${contractrewardsMap[$contractAddress]}
    echo "Cumulative Reward amount: $part2" | jq -R '{"message": .}'
   # Perform operations if totalAggregated is greater than 30
   # Print the timestamp in a specific format
    echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

    # Perform operations if totalAggregated is greater than 30
    if [ "$totalAggregated" -gt 0 ] && [ "$totalAggregated" -gt 30 ]; then
        iterations=$((totalAggregated / 2500))
        echo "Iterations: $iterations" | jq -R '{"message": .}'
        remaining=$((totalAggregated % 2500))

        echo "Remaining Record Limit: $remaining" | jq -R '{"message": .}'

        if ((iterations == 1)); then
            amount=$((part2 / iterations))
            echo "Reward Amount: $amount" | jq -R '{"message": .}'
            remainingamount=$((part2 % iterations))
            echo "Remaining Reward amount: $remainingamount" | jq -R '{"message": .}'
        fi


        if ((iterations != 0)); then
            amount=$((part2 / iterations))
            echo "Reward Amount: $amount" | jq -R '{"message": .}'
            remainingamount=$((part2 % iterations))
            echo "Remaining Reward amount: $remainingamount" | jq -R '{"message": .}'
        fi

        if ((iterations == 0)); then
           remainingamount=$part2
           echo "Reward Amount: $remainingamount" | jq -R '{"message": .}'
        fi
        echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
        # Perform operations for each iteration
        for i in $(seq 1 $iterations); do
            echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
            CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx bank send archway14zd6utea6u2zy5pd2yecphz8j9ydsq7x7qc8fu ${contractAddress} 1000000uarch --from pval4--keyring-backend=test --chain-id localnet --fees 17000uarch -y"
            echo "$CMD" | jq -R '{"message": .}'
            OUTPUT3=$(eval $CMD)
            json_output=$(echo "$OUTPUT3" | yq eval -j -)

                   # Embed the json_output in a new JSON structure using jq
            final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')
            echo $final_json
            echo "$OUTPUT3" | jq -R '{"message": .}'
            sleep 5
            echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
            echo "archway account balance" | jq -R '{"message": .}'
            CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet ${contractAddress}"
            echo "$CMD" | jq -R '{"message": .}'
            eval "$CMD"
            sleep 5
            echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
            command="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd tx --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 rewards withdraw-rewards --from $keyName --records-limit 35 --gas auto --gas-prices $(archwayd q rewards estimate-fees 1 --node http://localhost:26457 --output json | jq -r '.gas_unit_price | (.amount + .denom)') --gas-adjustment 1.4 -y"
            echo $command | jq -R '{"message": .}'
            OUTPUT=$(eval $command)
            json_output=$(echo "$OUTPUT" | yq eval -j -)

                               # Embed the json_output in a new JSON structure using jq
            final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')
            echo $final_json
            echo "$OUTPUT" | jq -R '{"message": .}'
            sleep 5
            echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
            echo "archway account balance" | jq -R '{"message": .}'
            CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet ${contractAddress}"
            echo "$CMD" | jq -R '{"message": .}'
            eval "$CMD"
            sleep 5
            echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
            command="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1  tx bank send ${contractAddress} archway1p4t985vch49tm632c9kd8qfag9gc2yfpcw845a ${amount}uarch --from $keyName --keyring-backend=test --chain-id localnet --fees 17000uarch -y"
            echo "$command" | jq -R '{"message": .}'
            OUTPUT3=$(eval $command)
            json_output=$(echo "$OUTPUT3" | yq eval -j -)

           # Embed the json_output in a new JSON structure using jq
            final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')
            echo $final_json
            echo "$OUTPUT3" | jq -R '{"message": .}'
            sleep 5
        done
        # Print the timestamp in a specific format
        echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

        if ((remaining != 0)); then
            CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx bank send archway14zd6utea6u2zy5pd2yecphz8j9ydsq7x7qc8fu ${contractAddress} 1000000uarch --from pval4--keyring-backend=test --chain-id localnet --fees 17000uarch -y"
            echo "$CMD" | jq -R '{"message": .}'
            OUTPUT3=$(eval $CMD)
            echo "$OUTPUT3" | jq -R '{"message": .}'
            json_output=$(echo "$OUTPUT3" | yq eval -j -)

                               # Embed the json_output in a new JSON structure using jq
            final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')
            echo $final_json
            sleep 5
            echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
            echo "archway account balance" | jq -R '{"message": .}'
            CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet ${contractAddress}"
            echo "$CMD" | jq -R '{"message": .}'
            OUTPUT3=$(eval "$CMD")
            echo "$OUTPUT3" | jq -R '{"message": .}'
            json_output=$(echo "$OUTPUT3" | yq eval -j -)

             # Embed the json_output in a new JSON structure using jq
            final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')
            sleep 5
            echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
            command="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd tx --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 rewards withdraw-rewards --from $keyName --records-limit $remaining --gas auto --gas-prices $(archwayd q rewards estimate-fees 1 --node http://localhost:26457 --output json | jq -r '.gas_unit_price | (.amount + .denom)') --gas-adjustment 1.4 -y"
            echo $command | jq -R '{"message": .}'
            OUTPUT=$(eval $command)
            echo "$OUTPUT" | jq -R '{"message": .}'
            json_output=$(echo "$OUTPUT" | yq eval -j -)

            # Embed the json_output in a new JSON structure using jq
            final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')
            echo "$OUTPUT" | jq -R '{"message": .}'
            sleep 5
            echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
            echo "archway account balance" | jq -R '{"message": .}'
            CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet ${contractAddress}"
            echo "$CMD" | jq -R '{"message": .}'
            OUTPUT=$(eval "$CMD")
            echo "$OUTPUT" | jq -R '{"message": .}'
            json_output=$(echo "$OUTPUT" | yq eval -j -)

                        # Embed the json_output in a new JSON structure using jq
            final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')
            sleep 5
            echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'
            command="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1  tx bank send ${contractAddress} archway1p4t985vch49tm632c9kd8qfag9gc2yfpcw845a ${remainingamount}uarch --from $keyName --keyring-backend=test --chain-id localnet --fees 17000uarch -y"
            echo "$command" | jq -R '{"message": .}'
            OUTPUT3=$(eval $command)
            echo "$OUTPUT3" | jq -R '{"message": .}'
            json_output=$(echo "$OUTPUT3" | yq eval -j -)

           # Embed the json_output in a new JSON structure using jq
            final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded')
            echo "$OUTPUT3" | jq -R '{"message": .}'
            sleep 5
       fi
       # Print the timestamp in a specific format
       echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R '{"message": .}'

   fi
done



