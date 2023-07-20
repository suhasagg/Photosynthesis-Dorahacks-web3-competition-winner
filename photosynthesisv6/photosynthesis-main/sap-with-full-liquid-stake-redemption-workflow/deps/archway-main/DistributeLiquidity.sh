#!/bin/bash

# Get the directory containing the script
script_dir=$(dirname "$(dirname "$(realpath "$0")")")
first_two_levels=$(echo "$script_dir" | cut -d'/' -f1-3)

FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquiditydistributiontodapps"
LiquidityFILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquiditydistributionAmountforDapps"

# Read the file line by line
while IFS=',' read -r part1 part2 part3 part4
do
    # Check if part4 is numeric
    if ! [[ $part4 =~ ^[0-9]+([.][0-9]+)?$ ]]; then
        echo "part4 is not numeric: $part4"
        continue
    fi

    if [ "$(echo $part3 | grep 'ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145')" ]; then
        CMD='/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd tx --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 bank send '"$part1"' '"$part2"' ${part4} --from pval5 --keyring-backend=test --chain-id localnet  >> /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/DappLiquiditytransferlogtocentralArchwayaddress'
        OUTPUT3=$(eval $CMD)
        echo "$OUTPUT3"
        echo "$part1,$part2,$cumulativeamountibc" >> $LiquidityFILE
        sleep 2
        break
    fi

    if [ "$(echo $line | grep 'amount:')" ]; then
        prev_line="$line"
    fi
done < "$FILE"
