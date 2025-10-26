package v3_0_0

import (
	"context"
	"fmt"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	evmkeeper "github.com/cosmos/evm/x/vm/keeper"
	// evm "github.com/cosmos/evm/x/vm"
	evmtypes "github.com/cosmos/evm/x/vm/types"
)

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	evmkeeper *evmkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(context context.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		ctx := sdk.UnwrapSDKContext(context)

		logger := ctx.Logger().With("upgrade", UpgradeName)

		evmParams := evmkeeper.GetParams(ctx)
		fmt.Println(evmParams)
		// evmParams.EvmDenom = BaseDenom
		evmParams.ExtendedDenomOptions = &evmtypes.ExtendedDenomOptions{ExtendedDenom: BaseDenom}
		// evmParams.AllowUnprotectedTxs = true // TODO:
		if err := evmkeeper.SetParams(ctx, evmParams); err != nil {
			return fromVM, err
		}

		// Initialize EvmCoinInfo in the module store
		if err := evmkeeper.InitEvmCoinInfo(ctx); err != nil {
			return nil, err
		}

		fmt.Println(evmParams)

		logger.Debug("running module migrations ...")
		// fromVM[evmtypes.ModuleName] = evm.AppModule{}.ConsensusVersion()

		return mm.RunMigrations(ctx, configurator, fromVM)
	}
}
