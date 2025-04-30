package types

import (
	"testing"

	"github.com/GGEZLabs/ggezchain/testutil/sample"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateAuthority_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateAuthority
		err  error
	}{
		// {
		// 	name: "invalid address",
		// 	msg: MsgUpdateAuthority{
		// 		Creator: "invalid_address",
		// 	},
		// 	err: sdkerrors.ErrInvalidAddress,
		// }, 
		// {
		// 	name: "valid address",
		// 	msg: MsgUpdateAuthority{
		// 		Creator: sample.AccAddress(),
		// 	},
		// },
		{
			name: "valid address",
			msg: MsgUpdateAuthority{
				Creator: sample.AccAddress(),
				AuthAddress: sample.AccAddress() ,
				DeleteModuleAccess: []string{"trade1"},
				UpdateModuleAccess: "{\"module\":\"trade1\",\"is_maker\":true,\"is_checker\":true}",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
