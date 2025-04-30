package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateAuthority{}

func NewMsgUpdateAuthority(creator string, authAddress string, newName string, newModuleAccess string, addModuleAccess string, updateModuleAccess string, deleteModuleAccess []string, clearAllModuleAccess bool) *MsgUpdateAuthority {
	return &MsgUpdateAuthority{
		Creator:              creator,
		AuthAddress:          authAddress,
		NewName:              newName,
		NewModuleAccess:      newModuleAccess,
		AddModuleAccess:      addModuleAccess,
		UpdateModuleAccess:   updateModuleAccess,
		DeleteModuleAccess:   deleteModuleAccess,
		ClearAllModuleAccess: clearAllModuleAccess,
	}
}

func (msg *MsgUpdateAuthority) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.AuthAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid AuthAddress (%s)", err)
	}

	// check if none of the flags provided
	hasUpdate := msg.NewName != "" ||
		msg.NewModuleAccess != "" ||
		msg.AddModuleAccess != "" ||
		msg.UpdateModuleAccess != "" ||
		len(msg.DeleteModuleAccess) > 0 ||
		msg.ClearAllModuleAccess

	if !hasUpdate {
		return ErrNoUpdateFlags
	}

	// if NewModuleAccess passed ignores other module access flags
	if msg.NewModuleAccess != "" {
		if msg.ClearAllModuleAccess || msg.UpdateModuleAccess != "" || msg.AddModuleAccess != "" || len(msg.DeleteModuleAccess) > 0 {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "NewModuleAccess cannot be combined with other module access flags")
		}
		return ValidateJSONFormat(msg.NewModuleAccess, "NewModuleAccess")
	}

	// if ClearAllModuleAccess is true ignores other module access flags
	if msg.ClearAllModuleAccess {
		if msg.UpdateModuleAccess != "" || msg.AddModuleAccess != "" || len(msg.DeleteModuleAccess) > 0 {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "ClearAllModuleAccess cannot be combined with other module access flags")
		}
		return nil
	}

	if msg.UpdateModuleAccess != "" {
		if err := ValidateJSONFormat(msg.UpdateModuleAccess, "UpdateModuleAccess"); err != nil {
			return err
		}
	}

	if msg.AddModuleAccess != "" {
		if err := ValidateJSONFormat(msg.AddModuleAccess, "AddModuleAccess"); err != nil {
			return err
		}
	}

	if (msg.UpdateModuleAccess != "" || msg.AddModuleAccess != "") && len(msg.DeleteModuleAccess) > 0 {
		if err := ValidateConflictBetweenModuleAccess(msg.UpdateModuleAccess, msg.AddModuleAccess, msg.DeleteModuleAccess); err != nil {
			return err
		}
	}

	return nil
}
