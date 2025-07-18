package keeper_test

import (
	"testing"

	keepertest "github.com/ramiqadoumi/ggezchain/v2/testutil/keeper"
	"github.com/ramiqadoumi/ggezchain/v2/testutil/nullify"
	"github.com/ramiqadoumi/ggezchain/v2/x/trade/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestTradeIndexQuery(t *testing.T) {
	keeper, ctx := keepertest.TradeKeeper(t)
	item := createTestTradeIndex(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetTradeIndexRequest
		response *types.QueryGetTradeIndexResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetTradeIndexRequest{},
			response: &types.QueryGetTradeIndexResponse{TradeIndex: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.TradeIndex(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
