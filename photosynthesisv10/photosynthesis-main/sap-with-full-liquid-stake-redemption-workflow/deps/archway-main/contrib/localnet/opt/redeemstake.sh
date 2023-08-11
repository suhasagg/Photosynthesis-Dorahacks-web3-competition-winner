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

    # If line contains 'denom: stuarch'
    if [ "$(echo $line | grep 'denom: stuarch')" ]; then
        # Get the amount from the previous line
        amount="${prev_line#*: \"}"
        # Remove trailing characters starting from '"'
        cumulativeamountstuarch="${amount%\"}"
        total_redemption_amount=$(echo "$total_redemption_amount + $cumulativeamountstuarch" | bc)
        echo "total redemption amount: $total_redemption_amount"
        echo "Amount: $cumulativeamountstuarch"
        CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx bank send ${part1} ${part2} ${cumulativeamountstuarch}stuarch --from l311 --keyring-backend=test --chain-id localnet --fees 17000arch -y"
        echo "$CMD"
        OUTPUT2=$(eval $CMD)
        echo "$OUTPUT2"
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
        echo "total redemption amount: $total_redemption_amount"
        echo "Amount: $cumulativeamountibc"
        CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1  tx bank send  ${part1} ${part2} ${cumulativeamountibc}ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145 --from l311 --keyring-backend=test --chain-id localnet --fees 17000uarch -y"
        echo "$CMD"
        OUTPUT3=$(eval $CMD)
        echo "$OUTPUT3"
        echo "$part1,$part2,$part3,$cumulativeamountibc" >> $LiquidityFILE
        sleep 5
        break
    fi

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

echo "archway liquidity account balance"
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"
eval "$CMD"

sleep 2

echo "archwayd ibc transfer"
# run the archwayd command
#archwayd --home $statearchwayd tx ibc-transfer transfer transfer channel-0 --node http://photo1:26657 --chain-id localnet stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 400uarch --from pval2 -y
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 tx ibc-transfer transfer transfer channel-0 --chain-id localnet stride1u20df3trc2c2zdhm8qvh2hdjx9ewh00sv6eyy8 ${total_redemption_amount}ibc/15CE03505E1F9891F448F53C9A06FD6C6AF9E5BE7CBB0A4B45F7BE5C9CBFC145 --from pval5 --keyring-backend test -y --fees 30000uarch"
echo "$CMD"
OUTPUT4=$(eval $CMD)
echo "$OUTPUT4"
# Delay for another 5 seconds
sleep 4


echo "strided redeem stake"
# Run the first strided command
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 tx stakeibc redeem-stake ${total_redemption_amount} localnet archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m --from admin --keyring-backend test --chain-id STRIDE -y"
echo "$CMD"
OUTPUT5=$(eval $CMD)
echo "$OUTPUT5"
# Delay for another 5 seconds
sleep 4

#echo "Number of successfully redeemed records" >> $LiquidityFILE

echo "archway balance uarch"
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"
echo "$CMD"
OUTPUT5=$(eval $CMD)
echo "$OUTPUT5"
sleep 4

total_arch_balance_amount1=0
total_arch_balance_amount2=0

while IFS= read -r line
do
  # If line contains 'denom: stuarch'
  if [ "$(echo $line | grep 'denom: uarch')" ]; then
    # Get the amount from the previous line
    amount="${prev_line#*: \"}"
    # Remove trailing characters starting from '"'
    amount="${amount%\"}"
    total_arch_balance_amount1=amount
    echo "Amount before claim: $amount"
    break
  fi
  # If line contains 'amount:'
  if [ "$(echo $line | grep 'amount:')" ]; then
    prev_line="$line"
  fi
done << EOF
$OUTPUT5
EOF


echo "list user redemption  record"
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 q records list-user-redemption-record -o=json"
echo "$CMD"
output=$(eval $CMD)
echo "$output"

sleep 4


echo "list epoch unbonding record"
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 q records list-epoch-unbonding-record -o=json"
echo "$CMD"
output=$(eval $CMD)
echo "$output"

# Delay for another 5 seconds
sleep 4


# Parse JSON from the variable
json=$(echo "$output" | jq -c '.')

claimed_records=0
stride_epoch=0
# Iterate over json array
echo "${json}" | jq -r '.epoch_unbonding_record[] | @base64' | while read epoch; do

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
                echo "Part 1: $part1"
                echo "Part 2: $part2"
                echo "Part 3: $part3"
                echo "claim undelegated tokens"
                CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/strided --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/stride1 tx stakeibc claim-undelegated-tokens ${part1} ${part2} ${part3} --from admin --keyring-backend test --chain-id STRIDE -y"
                echo "$CMD"
                output=$(eval $CMD)
                echo "$output"
                # Delay for another 5 seconds
                #echo "Number of successfully claimed records" >> /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/liquiditydistributionforDapps
                sleep 4
            done
        fi
    done
done

sleep 10

echo "archway redeemed uarch"
CMD="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/build/archwayd --home /media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/state/photo1 q bank balances --chain-id localnet archway1n3fvgm3ck5wylx6q4tsywglg82vxflj3h8e90m"
echo "$CMD"
OUTPUT5=$(eval $CMD)
echo "$OUTPUT5"
sleep 4

while IFS= read -r line
do
  # If line contains 'denom: stuarch'
  if [ "$(echo $line | grep 'denom: uarch')" ]; then
    # Get the amount from the previous line
    amount="${prev_line#*: \"}"
    # Remove trailing characters starting from '"'
    amount="${amount%\"}"
    total_arch_balance_amount2=amount
    echo "Amount after claim: $amount"
    break
  fi
  # If line contains 'amount:'
  if [ "$(echo $line | grep 'amount:')" ]; then
    prev_line="$line"
  fi
done << EOF
$OUTPUT5
EOF


FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/enableredeemstake"
OUTPUT_FILE="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redeemepoch"

if [ -f "$FILE" ]; then
    last_line=$(tail -n 1 "$FILE")
    echo "Last line of file is: $last_line"

    echo "$last_line" >> "$OUTPUT_FILE"
    echo "Last line written to $OUTPUT_FILE"
else
    echo "$FILE does not exist."
fi

rm "/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/enableredeemstake"
