#!/bin/bash

find . -type f -name "*.go" -exec sed -i \
    -e 's|github.com/cosmos/cosmos-sdk/x/auth/client/rest|github.com/cosmos/cosmos-sdk/x/auth/client/rest|g' \
    -e 's|google.golang.org/grpc/credentials/insecure|google.golang.org/grpc/v1.33.2/credentials/insecure|g' \
    -e 's|github.com/cosmos/cosmos-sdk/types/rest|github.com/cosmos/cosmos-sdk/types/rest|g' \
    -e 's|github.com/cosmos/cosmos-sdk/x/gov/client/rest|github.com/cosmos/cosmos-sdk/x/gov/client/rest|g' \
    -e 's|github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx|github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx|g' \
    {} +

