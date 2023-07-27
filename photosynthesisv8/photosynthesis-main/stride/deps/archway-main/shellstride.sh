#!/bin/bash

PARENT_DIR="."
DIRECTORIES=("x/icacallbacks" "x/interchainquery" "x/epochs")

for dir in "${DIRECTORIES[@]}"; do
  find "$PARENT_DIR/$dir" -type f -name "*.go" -exec sed -i 's|cosmos-sdk-v0.47.0|cosmos/cosmos-sdk/v0.47.0|g' {} +
done
