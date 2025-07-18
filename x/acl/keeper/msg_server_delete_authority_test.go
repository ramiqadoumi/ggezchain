package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ramiqadoumi/ggezchain/v2/testutil/sample"
	"github.com/ramiqadoumi/ggezchain/v2/x/acl/types"
	"github.com/stretchr/testify/require"
)

func TestMsgDeleteAuthority(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	superAdmin := sample.AccAddress()
	admin := sample.AccAddress()
	alice := sample.AccAddress()
	bob := sample.AccAddress()
	aclAuthorityAlice := types.AclAuthority{
		Address:           alice,
		Name:              "Alice",
		AccessDefinitions: []*types.AccessDefinition{},
	}
	aclAuthorityBob := types.AclAuthority{
		Address:           bob,
		Name:              "Bob",
		AccessDefinitions: []*types.AccessDefinition{},
	}
	k.SetSuperAdmin(ctx, types.SuperAdmin{Admin: superAdmin})
	k.SetAclAuthority(ctx, aclAuthorityAlice)
	k.SetAclAuthority(ctx, aclAuthorityBob)
	k.SetAclAdmin(ctx, types.AclAdmin{Address: admin})
	wctx := sdk.UnwrapSDKContext(ctx)

	testCases := []struct {
		name      string
		input     *types.MsgDeleteAuthority
		expErr    bool
		expErrMsg string
	}{
		{
			name: "unauthorized account",
			input: &types.MsgDeleteAuthority{
				Creator:     sample.AccAddress(),
				AuthAddress: alice,
			},
			expErr:    true,
			expErrMsg: "unauthorized account",
		},
		{
			name: "authority not found",
			input: &types.MsgDeleteAuthority{
				Creator:     admin,
				AuthAddress: sample.AccAddress(),
			},
			expErr:    true,
			expErrMsg: "authority address does not exist",
		},
		{
			name: "delete authority by super admin",
			input: &types.MsgDeleteAuthority{
				Creator:     superAdmin,
				AuthAddress: alice,
			},
			expErr:    false,
			expErrMsg: "",
		},
		{
			name: "delete authority by admin",
			input: &types.MsgDeleteAuthority{
				Creator:     admin,
				AuthAddress: bob,
			},
			expErr:    false,
			expErrMsg: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.DeleteAuthority(wctx, tc.input)

			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
