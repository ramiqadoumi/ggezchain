package acl

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/GGEZLabs/ggezchain/api/ggezchain/acl"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "AclAuthorityAll",
					Use:       "list-acl-authority",
					Short:     "List all aclAuthority",
				},
				{
					RpcMethod:      "AclAuthority",
					Use:            "show-acl-authority [id]",
					Short:          "Shows a aclAuthority",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "address"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "AddAuthority",
					Use:            "add-authority [auth-address] [name] [module-access]",
					Short:          "Add a new authority with specific module access permissions. Must have admin authority to do so.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auth_address"}, {ProtoField: "name"}, {ProtoField: "module_access"}},
				},
				{
					RpcMethod:      "DeleteAuthority",
					Use:            "delete-authority [auth-address]",
					Short:          "Delete an existing authority. Must have admin authority to do so.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auth_address"}},
				},
				{
					RpcMethod:      "UpdateAuthority",
					Use:            "update-authority [auth-address]",
					Short:          "Modify the name or module access permissions of an existing authority. Must have admin authority to do so.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "auth_address"}},
					FlagOptions: map[string]*autocliv1.FlagOptions{
						"new_name": {
							Name:         "new-name",
							Usage:        "Set a new name for the authority.",
							DefaultValue: "",
						},
						"new_module_access": {
							Name:         "new-module-access",
							Usage:        "Overwrite the entire module access list with this JSON string. Ignores other module access flags.",
							DefaultValue: "",
						},
						"add_module_access": {
							Name:         "add-module-access",
							Usage:        "Add one or more new module access.",
							DefaultValue: "",
						},
						"update_module_access": {
							Name:         "update-module-access",
							Usage:        "Update module access values for an existing module. (matched by module name)",
							DefaultValue: "",
						},
						"delete_module_access": {
							Name:         "delete-module-access",
							Usage:        "Delete one or more specific module access (by module name).",
							DefaultValue: "",
						},
						"clear_all_module_access": {
							Name:         "clear-all-module-access",
							Usage:        "Clear all module access. Default is false.",
							DefaultValue: "false",
						},
					},
					Example: `Overwrite the entire module access list with this JSON string. Ignores other module access flags:
ggezchaind tx acl update-authority ggezauthaddress... --new-module-access '[{"module":"module1","is_maker":true,"is_checker":false}]' --from ggezauthaddress...

Add one or more new module access:
ggezchaind tx acl update-authority ggezauthaddress... --add-module-access '[{"module":"module2","is_maker":true,"is_checker":true}]' --from ggezauthaddress...

Update module access values for an existing module (by module name):
ggezchaind tx acl update-authority ggezauthaddress... --update-module-access '{"module":"module2","is_maker":false,"is_checker":true}' --from ggezauthaddress...

Delete one or more specific module access (by module name):
ggezchaind tx acl update-authority ggezauthaddress... --delete-module-access 'module2,module1' --from ggezauthaddress...

Clear all module access. Default is false ()
`,
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
