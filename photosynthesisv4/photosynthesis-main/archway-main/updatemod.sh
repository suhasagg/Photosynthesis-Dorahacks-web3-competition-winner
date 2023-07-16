#!/bin/bash

# Backup your go.mod.bk2 file before running the script
cp go.mod go.mod.bak

# Update the go.mod.bk2 file with the correct replacement paths
sed -i \
    -e 's|github.com/cosmos/cosmos-sdk/x/auth/client/rest|github.com/cosmos/cosmos-sdk@v0.45.11/x/auth/client/rest|g' \
    -e 's|google.golang.org/grpc/credentials/insecure|google.golang.org/grpc@v1.33.2/credentials/insecure|g' \
    -e 's|github.com/cosmos/cosmos-sdk/types/rest|github.com/cosmos/cosmos-sdk@v0.45.11/types/rest|g' \
    -e 's|github.com/cosmos/cosmos-sdk/x/gov/client/rest|github.com/cosmos/cosmos-sdk@v0.45.11/x/gov/client/rest|g' \
    -e 's|github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx|github.com/cosmos/cosmos-sdk@v0.45.11/x/auth/legacy/legacytx|g' \
    go.mod

find . -name '*.go' -type f -exec sed -i 's/github.com\/cosmos\/cosmos-sdk\/simapp/cosmossdk.io\/simapp/g' {} +
