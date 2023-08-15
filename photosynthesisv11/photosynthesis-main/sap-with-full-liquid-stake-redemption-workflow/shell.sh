#!/bin/bash

FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/rewardswithdrawaltodapps"

# Loop over the aggregated data and generate commands
while IFS=',' read -r part1 part2 part3
do
    echo $part1
    echo $part2
    echo $part3
    # Convert the line into array
    contractAddress=$part1 # First element is contract address
    keyName=${contractToKeyName[$contractAddress]}

    recordLimit=$part2 # Second element is the aggregated record limit
    totalAmount=$part3 # Third element is the total amount

    command="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd tx --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 rewards withdraw-rewards --from $keyName --records-limit $recordLimit --fees 17000uarch $contractAddress -y"
    echo $command
    OUTPUT=$(eval $command)
    echo "$OUTPUT"
    sleep 5

    command="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1  tx bank send ${contractAddress} archway1p4t985vch49tm632c9kd8qfag9gc2yfpcw845a ${totalAmount}uarch --from $keyName --keyring-backend=test --chain-id localnet --fees 17000uarch -y"
    echo "$command"
    OUTPUT3=$(eval $command)
    echo "$OUTPUT3"
    sleep 5
done < "$FILE"

