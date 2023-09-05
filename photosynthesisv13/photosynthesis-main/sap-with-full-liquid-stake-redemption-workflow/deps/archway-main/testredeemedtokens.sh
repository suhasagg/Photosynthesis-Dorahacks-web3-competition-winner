#!/bin/bash

log_file="/media/usbHDD1/photov10/Photosynthesis-Dorahacks-web3-competition-winner/photosynthesisv5/photosynthesis-main/sap-with-full-liquid-stake-redemption-workflow/dockernet/logs/redeemstake.log"

# Check if the log file exists
if [ ! -f "$log_file" ]; then
    echo "Error: Log file $log_file does not exist."
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


