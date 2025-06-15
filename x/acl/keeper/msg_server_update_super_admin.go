package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ramiqadoumi/ggezchain/x/acl/types"
)

func (k msgServer) UpdateSuperAdmin(goCtx context.Context, msg *types.MsgUpdateSuperAdmin) (*types.MsgUpdateSuperAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsSuperAdmin(ctx, msg.Creator) {
		return nil, types.ErrUnauthorized
	}

	k.SetSuperAdmin(ctx, types.SuperAdmin{Admin: msg.NewSuperAdmin})

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeInit,
			sdk.NewAttribute(types.AttributeKeySuperAdmin, msg.NewSuperAdmin),
		),
	)
	return &types.MsgUpdateSuperAdminResponse{}, nil
}
