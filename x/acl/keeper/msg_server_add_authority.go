package keeper

import (
	"context"
	"strings"

	"github.com/GGEZLabs/ggezchain/x/acl/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddAuthority(goCtx context.Context, msg *types.MsgAddAuthority) (*types.MsgAddAuthorityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsAdmin(ctx, msg.Creator) {
		return nil, types.ErrUnauthorized
	}

	_, found := k.GetAclAuthority(ctx, msg.AuthAddress)
	if found {
		return nil, types.ErrAuthorityAddressExist
	}

	moduleAccess, err := types.ValidateModuleAccessList(msg.ModuleAccess)
	if err != nil {
		return nil, err
	}

	aclAuthority := types.AclAuthority{
		Address:      msg.AuthAddress,
		Name:         strings.TrimSpace(msg.Name),
		ModuleAccess: moduleAccess,
	}
	k.SetAclAuthority(ctx, aclAuthority)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeAddAuthority,
			sdk.NewAttribute(types.AttributeKeyAuthorityAddress, msg.AuthAddress),
			sdk.NewAttribute(types.AttributeKeyName, msg.Name),
			sdk.NewAttribute(types.AttributeKeyModuleAccess, msg.ModuleAccess),
		),
	)
	return &types.MsgAddAuthorityResponse{}, nil
}
