package keeper_test

import (
	"testing"

	keepertest "github.com/ramiqadoumi/ggezchain/v2/testutil/keeper"
	"github.com/ramiqadoumi/ggezchain/v2/x/acl/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AclKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
