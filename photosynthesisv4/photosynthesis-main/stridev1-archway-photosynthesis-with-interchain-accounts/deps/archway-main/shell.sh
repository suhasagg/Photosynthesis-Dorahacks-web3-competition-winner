#!/bin/bash

# Set up repositories and required versions
stride_repo="https://github.com/Stride-Labs/stride.git"
stride_version="v5.0.0" # replace with the actual version you're using
stride_cosmos_version="v0.48.0"

archway_repo="https://github.com/archway-network/archway.git"
archway_version="v1.0.0" # replace with the actual version you're using
archway_cosmos_version="v0.45.11"

# Clone, switch to required versions, and update go.mod.bk2
git clone $stride_repo stride-v5-with-cosmos-sdk-v0.47.1
cd stride-v5-with-cosmos-sdk-v0.48.0
git checkout $stride_version
sed -i "s|github.com/cosmos/cosmos-sdk .*|github.com/cosmos/cosmos-sdk $stride_cosmos_version|" go.mod
cd ..

git clone $archway_repo archway-with-cosmos-sdk-v0.45.11
cd archway-with-cosmos-sdk-v0.45.11
git checkout $archway_version
sed -i "s|github.com/cosmos/cosmos-sdk .*|github.com/cosmos/cosmos-sdk $archway_cosmos_version|" go.mod
cd ..

# Update your project's go.mod.bk2
sed -i "s|github.com/Stride-Labs/stride/v4 .*|replace github.com/Stride-Labs/stride/v4 => ./stride-v5-with-cosmos-sdk-v0.48.0|" go.mod
sed -i "s|github.com/archway-network/archway .*|replace github.com/archway-network/archway => ./archway-with-cosmos-sdk-v0.45.11|" go.mod

# Add the older versions of the missing packages
echo "replace github.com/cosmos/cosmos-sdk/x/auth/client/rest => github.com/cosmos/cosmos-sdk/x/auth/client/rest" >> go.mod
echo "replace google.golang.org/grpc/credentials/insecure => google.golang.org/grpc/v1.33.2/credentials/insecure" >> go.mod
echo "replace github.com/cosmos/cosmos-sdk/types/rest => github.com/cosmos/cosmos-sdk/types/rest" >> go.mod
echo "replace github.com/cosmos/cosmos-sdk/x/gov/client/rest => github.com/cosmos/cosmos-sdk/x/gov/client/rest" >> go.mod
echo "replace github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx => github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx" >> go.mod

