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


FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redeemliquidityamountforDapps"
LiquidityFILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquiditydistributionforDapps"
# Read the file line by line
while IFS=',' read -r part1 part2 part3

do

  CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1  tx bank send  archway14zd6utea6u2zy5pd2yecphz8j9ydsq7x7qc8fu ${part1} 1000000uarch --from pval4 --keyring-backend=test --chain-id localnet --fees 17000uarch -y"
  echo "$CMD"
  OUTPUT3=$(eval $CMD)
  echo "$OUTPUT3"
  sleep 5

done < "$FILE"

