package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ramiqadoumi/ggezchain/v2/testutil/sample"
	"github.com/ramiqadoumi/ggezchain/v2/x/acl/types"
	"github.com/stretchr/testify/require"
)

func TestMsgDeleteAdmin(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	superAdmin := sample.AccAddress()
	alice := sample.AccAddress()
	bob := sample.AccAddress()

	k.SetSuperAdmin(ctx, types.SuperAdmin{Admin: superAdmin})
	aclAdmins := types.ConvertStringsToAclAdmins([]string{alice, bob})
	k.SetAclAdmins(ctx, aclAdmins)

	testCases := []struct {
		name        string
		input       *types.MsgDeleteAdmin
		expectedLen int
		expErr      bool
		expErrMsg   string
	}{
		{
			name: "address unauthorized",
			input: &types.MsgDeleteAdmin{
				Creator: sample.AccAddress(),
			},
			expErr:    true,
			expErrMsg: "unauthorized account",
		},
		{
			name: "admin not exist",
			input: &types.MsgDeleteAdmin{
				Creator: superAdmin,
				Admins:  []string{sample.AccAddress()},
			},
			expErr:    true,
			expErrMsg: "admin(s) does not exist",
		},
		{
			name: "all good",
			input: &types.MsgDeleteAdmin{
				Creator: superAdmin,
				Admins:  []string{alice},
			},
			expectedLen: 1,
			expErr:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.DeleteAdmin(wctx, tc.input)

			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.Len(t, k.GetAllAclAdmin(ctx), tc.expectedLen)
				require.NoError(t, err)
			}
		})
	}
}
