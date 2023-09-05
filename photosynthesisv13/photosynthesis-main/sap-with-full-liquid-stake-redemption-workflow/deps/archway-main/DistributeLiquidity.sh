#!/bin/sh

# Get the directory containing the script
script_dir=`dirname "$0"`
first_two_levels=`echo "$script_dir" | cut -d'/' -f1-3`

FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquiditydistributiontodapps"
LiquidityFILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquiditydistributionAmountforDapps"

# Read the file line by line
while IFS=',' read -r part1 part2 part3 part4
do

    case $part4 in
        ''|*[!0-9.]*)
            echo "part4 is not numeric: $part4" | jq -R -c '{"message": .}'
            continue ;;
        0|0.0|0.00)
            echo "part4 should not be equal to zero" | jq -R -c '{"message": .}'
            continue ;;
    esac


    part4=`echo $part4 | cut -d '.' -f 1`
    echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
    if echo $part3 | grep -q 'ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145'; then
        CMD='/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd tx --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 bank send '"$part1"' '"$part2"' '"$part4"'ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145 --from pval5 --keyring-backend=test --chain-id localnet --fees 17000uarch -y'
        echo "Executing: $CMD" | jq -R -c '{"message": .}'
        OUTPUT3=`$CMD`
        #echo "$OUTPUT3" | jq -R -c '{"message": .}'
        json_output=$(echo "$OUTPUT3" | yq eval -j -)

      # Embed the json_output in a new JSON structure using jq
        #final_json=$(echo '{}' | jq --argjson embedded "$json_output" '.message = $embedded' -c)

        echo $json_output
        #echo "$OUTPUT3" | jq -R -c '{"message": .}'
        echo "$part1,$part2,$cumulativeamountibc" >> $LiquidityFILE
        sleep 5
        break
    fi
    echo $(date +"%Y-%m-%d %H:%M:%S") | jq -R -c '{"message": .}'
    if echo $line | grep -q 'amount:'; then
        prev_line="$line"
    fi
    sed -i '1d' "$FILE"
done < "$FILE"
