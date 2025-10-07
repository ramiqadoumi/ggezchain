package v3_0_0

import (
	"context"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	// erc20types "github.com/cosmos/evm/x/erc20/types"
	// erc20 "github.com/cosmos/evm/x/erc20"
	// feemarkettypes "github.com/cosmos/evm/x/feemarket/types"
	// feemarket "github.com/cosmos/evm/x/feemarket"
	// precisebanktypes "github.com/cosmos/evm/x/precisebank/types"
	// precisebank "github.com/cosmos/evm/x/precisebank"
	// evmtypes "github.com/cosmos/evm/x/vm/types"
	// evm "github.com/cosmos/evm/x/vm"

)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(context context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(context)
		// fromVM[erc20types.ModuleName] = erc20.AppModule{}.ConsensusVersion()
		// fromVM[feemarkettypes.ModuleName] = feemarket.AppModule{}.ConsensusVersion()
		// fromVM[precisebanktypes.ModuleName] = precisebank.AppModule{}.ConsensusVersion()
		// fromVM[evmtypes.ModuleName] = evm.AppModule{}.ConsensusVersion()

		logger := ctx.Logger().With("upgrade", UpgradeName)

		logger.Debug("running module migrations ...")
		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}