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


FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/enableredeemstake"

if [ -f "$FILE" ]; then
    echo "$FILE exists."

    AMOUNT=0
    while IFS= read -r line
    do
        AMOUNT=$((AMOUNT + line))
    done < "$FILE"

    echo "Total Liquidity amount: $AMOUNT"
else
    echo "$FILE does not exist."
    exit 0
fi

truncate -s 0 "$FILE"

echo "strided redeem stake"
# Run the first strided command
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 tx stakeibc redeem-stake ${AMOUNT} stuarch localnet archway15js809uedxqs2wl0lyt58httasr5rlplj3atd8 --from admin --keyring-backend test --chain-id STRIDE -y"
eval "$CMD"

# Delay for another 5 seconds
sleep 4


FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/enableredeemstake"
OUTPUT_FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redeemepoch"

if [ -f "$FILE" ]; then
    last_line=$(tail -n 1 "$FILE")
    echo "Last line of file is: $last_line"

    echo "$last_line" >> "$OUTPUT_FILE"
    echo "Last line written to $OUTPUT_FILE"
else
    echo "$FILE does not exist."
fi



rm "/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv4/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/enableredeemstake"
