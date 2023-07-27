package app

import (
	"fmt"
	wasmdKeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmdTypes "github.com/CosmWasm/wasmd/x/wasm/types"
	cosmwasm "github.com/CosmWasm/wasmvm"
	epochsmodule "github.com/Stride-Labs/stride/v4/x/epochs"
	epochsmodulekeeper "github.com/Stride-Labs/stride/v4/x/epochs/keeper"
	epochtypes "github.com/Stride-Labs/stride/v4/x/epochs/types"
	"github.com/Stride-Labs/stride/v4/x/interchainquery"
	"github.com/archway-network/archway/x/photosynthesis"
	"github.com/archway-network/archwa