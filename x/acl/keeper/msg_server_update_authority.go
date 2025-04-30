package keeper

import (
	"context"
	"encoding/json"

	"github.com/GGEZLabs/ggezchain/x/acl/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateAuthority(goCtx context.Context, msg *types.MsgUpdateAuthority) (*types.MsgUpdateAuthorityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.IsAdmin(ctx, msg.Creator) {
		return nil, types.ErrUnauthorized
	}

	aclAuthority, found := k.GetAclAuthority(ctx, msg.AuthAddress)
	if !found {
		return nil, types.ErrAuthorityAddressNotExist
	}

	if msg.NewName != "" {
		aclAuthority = k.UpdateAclAuthorityName(aclAuthority, msg.NewName)
	}

	var err error
	// if NewModuleAccess passed ignore another flags
	if msg.NewModuleAccess != "" {
		aclAuthority, err = k.OverwriteModuleAccessList(aclAuthority, msg.NewModuleAccess)
		if err != nil {
			return nil, err
		}
	} else if msg.ClearAllModuleAccess {
		// if ClearAllModuleAccess passed ignore another flags
		aclAuthority = k.ClearAllModuleAccess(aclAuthority)
	} else {
		if msg.UpdateModuleAccess != "" {
			aclAuthority, err = k.UpdateModuleAccess(aclAuthority, msg.UpdateModuleAccess)
			if err != nil {
				return nil, err
			}
		}

		if msg.AddModuleAccess != "" {
			aclAuthority, err = k.AddModuleAccess(aclAuthority, msg.AddModuleAccess)
			if err != nil {
				return nil, err
			}
		}

		if len(msg.DeleteModuleAccess) != 0 {
			aclAuthority, err = k.DeleteModuleAccess(aclAuthority, msg.DeleteModuleAccess)
			if err != nil {
				return nil, err
			}
		}
	}
	// apply updated aclAuthority
	k.SetAclAuthority(ctx, aclAuthority)

	var moduleAccessJSON []byte = []byte("[]")
	if aclAuthority.ModuleAccess != nil {
		moduleAccessJSON, err = json.Marshal(aclAuthority.ModuleAccess)
		if err != nil {
			moduleAccessJSON = []byte("[]")
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdateAuthority,
			sdk.NewAttribute(types.AttributeKeyAuthorityAddress, msg.AuthAddress),
			sdk.NewAttribute(types.AttributeKeyName, aclAuthority.Name),
			sdk.NewAttribute(types.AttributeKeyModuleAccess, string(moduleAccessJSON)),
		),
	)

	return &types.MsgUpdateAuthorityResponse{}, nil
}
