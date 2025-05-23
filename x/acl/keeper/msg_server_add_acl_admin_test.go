package keeper_test

import (
	"testing"

	"github.com/GGEZLabs/ggezchain/testutil/sample"
	"github.com/GGEZLabs/ggezchain/x/acl/types"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestMsgAddAclAdmin(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	aclAdmin := sample.AccAddress()
	duplicateAdmin := sample.AccAddress()

	k.SetAclAdmin(ctx, types.AclAdmin{Address: aclAdmin})
	k.SetAclAdmin(ctx, types.AclAdmin{Address: duplicateAdmin})

	testCases := []struct {
		name        string
		input       *types.MsgAddAclAdmin
		expectedLen int
		expErr      bool
		expErrMsg   string
	}{
		{
			name: "address unauthorized",
			input: &types.MsgAddAclAdmin{
				Creator: sample.AccAddress(),
			},
			expErr:    true,
			expErrMsg: "unauthorized account",
		},
		{
			name: "duplicate admin",
			input: &types.MsgAddAclAdmin{
				Creator: aclAdmin,
				Admins:  []string{duplicateAdmin, sample.AccAddress()},
			},
			expErr:    true,
			expErrMsg: "admin(s) already exist",
		},
		{
			name: "all good",
			input: &types.MsgAddAclAdmin{
				Creator: aclAdmin,
				Admins:  []string{sample.AccAddress(), sample.AccAddress()},
			},
			// duplicateAdmin + aclAdmin + 2
			expectedLen: 4,
			expErr:      false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.AddAclAdmin(wctx, tc.input)

			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.Equal(t, len(k.GetAllAclAdmin(ctx)), tc.expectedLen)
				require.NoError(t, err)
			}
		})
	}
}
