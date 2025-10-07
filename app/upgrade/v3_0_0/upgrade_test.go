package v3_0_0_test

import (
	"fmt"
	"testing"

	upgradetypes "cosmossdk.io/x/upgrade/types"
	"github.com/stretchr/testify/require"

	"github.com/ramiqadoumi/ggezchain/v2/app"
	v3_0_0 "github.com/ramiqadoumi/ggezchain/v2/app/upgrade/v3_0_0"
	"github.com/cosmos/cosmos-sdk/types/module"
	erc20types "github.com/cosmos/evm/x/erc20/types"
	feemarkettypes "github.com/cosmos/evm/x/feemarket/types"
	precisebanktypes "github.com/cosmos/evm/x/precisebank/types"
	evmtypes "github.com/cosmos/evm/x/vm/types"
)

func TestUpgradeToV3_0_0(t *testing.T) {
	// Initialize app in-memory (no disk I/O)
	gApp := app.Setup(t, false)
	ctx := gApp.BaseApp.NewContext(false)

	upgradeKeeper := gApp.UpgradeKeeper

	// Define the upgrade plan
	plan := upgradetypes.Plan{
		Name:   v3_0_0.UpgradeName,
		Height: 15,
		Info:   "test upgrade to v3.0.0",
	}

	// Set upgrade plan
	err := upgradeKeeper.ScheduleUpgrade(ctx, plan)
	require.NoError(t, err)

	// Verify upgrade plan stored
	storedPlan, err := upgradeKeeper.GetUpgradePlan(ctx)
	require.NoError(t, err)
	require.Equal(t, v3_0_0.UpgradeName, storedPlan.Name)

	// Run until upgrade height
	for ctx.BlockHeight() < plan.Height {
		ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	}

	// Simulate upgrade execution
	handler := v3_0_0.CreateUpgradeHandler(gApp.ModuleManager, gApp.Configurator())
	vMap, err := handler(ctx, plan, module.VersionMap{})
	require.NoError(t, err)
	fmt.Println(vMap)
	// Verify new stores are registered
	// storeKeys := gApp.GetStoreKeys()
	newStores := []string{
		erc20types.StoreKey,
		feemarkettypes.StoreKey,
		precisebanktypes.StoreKey,
		evmtypes.StoreKey,
	}

	for _, store := range newStores {
		_, exists := vMap[store]
		require.Truef(t, exists, "store %s should exist after upgrade", store)
	}

	t.Log("✅ Upgrade to v3.0.0 executed successfully and new stores are added.")
}
