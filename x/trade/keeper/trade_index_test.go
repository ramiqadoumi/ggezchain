package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/ramiqadoumi/ggezchain/testutil/keeper"
	"github.com/ramiqadoumi/ggezchain/testutil/nullify"
	"github.com/ramiqadoumi/ggezchain/x/trade/keeper"
	"github.com/ramiqadoumi/ggezchain/x/trade/types"
	"github.com/stretchr/testify/require"
)

func createTestTradeIndex(keeper keeper.Keeper, ctx context.Context) types.TradeIndex {
	item := types.TradeIndex{
		NextId: 1,
	}
	keeper.SetTradeIndex(ctx, item)
	return item
}

func TestTradeIndexGet(t *testing.T) {
	keeper, ctx := keepertest.TradeKeeper(t)
	item := createTestTradeIndex(keeper, ctx)
	rst, found := keeper.GetTradeIndex(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestTradeIndexRemove(t *testing.T) {
	keeper, ctx := keepertest.TradeKeeper(t)
	createTestTradeIndex(keeper, ctx)
	keeper.RemoveTradeIndex(ctx)
	_, found := keeper.GetTradeIndex(ctx)
	require.False(t, found)
}
