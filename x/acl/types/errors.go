package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/acl module sentinel errors
var (
	ErrInvalidSigner             = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrEmptyName                 = sdkerrors.Register(ModuleName, 1101, "empty name not allowed")
	ErrInvalidModuleAccessList   = sdkerrors.Register(ModuleName, 1102, "invalid ModuleAccessList format")
	ErrInvalidModuleAccessObject = sdkerrors.Register(ModuleName, 1103, "invalid ModuleAccessObject format")
	ErrInvalidModuleName         = sdkerrors.Register(ModuleName, 1104, "invalid module name")
	ErrUnauthorized              = sdkerrors.Register(ModuleName, 1105, "unauthorized account")
	ErrAuthorityAddressExist     = sdkerrors.Register(ModuleName, 1106, "authority address already exist")
	ErrAuthorityAddressNotExist  = sdkerrors.Register(ModuleName, 1107, "authority address not exist")
	ErrModuleNotExist            = sdkerrors.Register(ModuleName, 1108, "module not exist")
	ErrModuleExist               = sdkerrors.Register(ModuleName, 1109, "module exist")
	ErrEmptyModuleAccessList     = sdkerrors.Register(ModuleName, 1110, "empty module access list")
	ErrNoUpdateFlags             = sdkerrors.Register(ModuleName, 1111, "at least one update flag must be provided")
	ErrUpdateAndRemoveModule     = sdkerrors.Register(ModuleName, 1113, "cannot be both modified (added/updated) and removed in the same request")
)
