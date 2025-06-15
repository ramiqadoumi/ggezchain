package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ramiqadoumi/ggezchain/x/acl/types"
)

func (k msgServer) DeleteAdmin(goCtx context.Context, msg *types.MsgDeleteAdmin) (*types.MsgDeleteAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsSuperAdmin(ctx, msg.Creator) {
		return nil, types.ErrUnauthorized
	}

	err := types.ValidateDeleteAdmin(k.GetAllAclAdmin(ctx), msg.Admins)
	if err != nil {
		return nil, err
	}

	k.RemoveAclAdmins(ctx, msg.Admins)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeDeleteAdmin,
			sdk.NewAttribute(types.AttributeKeyAdmins, strings.Join(msg.Admins, ",")),
		),
	)
	return &types.MsgDeleteAdminResponse{}, nil
}
