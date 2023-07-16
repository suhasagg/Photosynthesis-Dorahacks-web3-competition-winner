package v11

import (
	"github.com/CosmosContracts/juno/v14/app/upgrades"
	store "github.com/cosmos/cosmos-sdk/store/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"
)

// UpgradeName defines the on-chain upgrade name for the Juno v11 upgrade.
const UpgradeName = "v11" // maybe multiverse?

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV11UpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		Added: []string{icacontrollertypes.StoreKey, icahosttypes.StoreKey},
	},
}
