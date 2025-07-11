package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/ramiqadoumi/ggezchain/v2/testutil/keeper"
	"github.com/ramiqadoumi/ggezchain/v2/x/acl/keeper"
	"github.com/ramiqadoumi/ggezchain/v2/x/acl/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(tb testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	tb.Helper()
	k, ctx := keepertest.AclKeeper(tb)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
