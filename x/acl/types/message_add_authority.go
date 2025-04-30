package types

import (
	"encoding/json"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAddAuthority{}

func NewMsgAddAuthority(creator string, authAddress string, name string, moduleAccess string) *MsgAddAuthority {
	return &MsgAddAuthority{
		Creator:      creator,
		AuthAddress:  authAddress,
		Name:         name,
		ModuleAccess: moduleAccess,
	}
}

func (msg *MsgAddAuthority) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(msg.AuthAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid AuthAddress (%s)", err)
	}

	if msg.Name == "" {
		return ErrEmptyName
	}

	isValid := json.Valid([]byte(msg.ModuleAccess))
	if !isValid {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid ModuleAccess JSON format")
	}

	return nil
}
