package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ramiqadoumi/ggezchain/v2/x/acl/types"
)

func (k msgServer) Init(goCtx context.Context, msg *types.MsgInit) (*types.MsgInitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetSuperAdmin(ctx)
	if found {
		return nil, types.ErrSuperAdminInitialized
	}

	// Set super admin
	k.SetSuperAdmin(ctx, types.SuperAdmin{Admin: msg.SuperAdmin})

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeInit,
			sdk.NewAttribute(types.AttributeKeySuperAdmin, msg.SuperAdmin),
		),
	)
	return &types.MsgInitResponse{}, nil
}
