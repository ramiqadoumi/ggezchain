package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ramiqadoumi/ggezchain/testutil/sample"
	"github.com/stretchr/testify/require"
)

func TestMsgDeleteAuthority_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteAuthority
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgDeleteAuthority{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "invalid auth address",
			msg: MsgDeleteAuthority{
				Creator:     sample.AccAddress(),
				AuthAddress: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "all good",
			msg: MsgDeleteAuthority{
				Creator:     sample.AccAddress(),
				AuthAddress: sample.AccAddress(),
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
