package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ramiqadoumi/ggezchain/v2/x/acl/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SuperAdmin(goCtx context.Context, req *types.QueryGetSuperAdminRequest) (*types.QueryGetSuperAdminResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetSuperAdmin(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetSuperAdminResponse{SuperAdmin: val}, nil
}
