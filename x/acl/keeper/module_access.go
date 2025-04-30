package keeper

import (
	"strings"

	"github.com/GGEZLabs/ggezchain/x/acl/types"
)

// UpdateAclAuthorityName update authority name
func (k Keeper) UpdateAclAuthorityName(aclAuthority types.AclAuthority, name string) types.AclAuthority {
	aclAuthority.Name = strings.TrimSpace(name)

	return aclAuthority
}

// OverwriteModuleAccessList completely replace the module access list
func (k Keeper) OverwriteModuleAccessList(aclAuthority types.AclAuthority, moduleAccessListStr string) (types.AclAuthority, error) {
	newModuleAccess, err := types.ValidateModuleAccessList(moduleAccessListStr)
	if err != nil {
		return types.AclAuthority{}, err
	}

	aclAuthority.ModuleAccess = newModuleAccess

	return aclAuthority, nil
}

// UpdateModuleAccess update specific module permission
func (k Keeper) UpdateModuleAccess(aclAuthority types.AclAuthority, singleModuleAccessStr string) (types.AclAuthority, error) {
	updatedModuleAccess, err := types.ValidateSingleModuleAccess(singleModuleAccessStr)
	if err != nil {
		return types.AclAuthority{}, err
	}

	err = types.ValidateUpdateModuleAccess(aclAuthority, updatedModuleAccess.Module)
	if err != nil {
		return types.AclAuthority{}, err
	}

	updatedModuleAccessList := types.GetUpdatedModuleAccessList(aclAuthority.ModuleAccess, updatedModuleAccess)
	aclAuthority.ModuleAccess = updatedModuleAccessList

	return aclAuthority, nil
}

// AddModuleAccess add one or more module access
func (k Keeper) AddModuleAccess(aclAuthority types.AclAuthority, moduleAccessListStr string) (types.AclAuthority, error) {
	moduleAccessList, err := types.ValidateModuleAccessList(moduleAccessListStr)
	if err != nil {
		return types.AclAuthority{}, err
	}

	newModules := types.GetAuthorityModules(moduleAccessList)
	currentModules := types.GetAuthorityModules(aclAuthority.ModuleAccess)

	err = types.ValidateAddModuleAccess(currentModules, newModules)
	if err != nil {
		return types.AclAuthority{}, err
	}

	aclAuthority.ModuleAccess = append(aclAuthority.ModuleAccess, moduleAccessList...)

	return aclAuthority, nil
}

// DeleteModuleAccess remove one or more module access
func (k Keeper) DeleteModuleAccess(aclAuthority types.AclAuthority, moduleNames []string) (types.AclAuthority, error) {

	newModuleAccessList, err := types.ValidateDeleteModuleAccess(moduleNames, aclAuthority.ModuleAccess)
	if err != nil {
		return types.AclAuthority{}, err
	}

	aclAuthority.ModuleAccess = newModuleAccessList

	return aclAuthority, nil
}

// ClearAllModuleAccess clear all module access
func (k Keeper) ClearAllModuleAccess(aclAuthority types.AclAuthority) types.AclAuthority {
	aclAuthority.ModuleAccess = make([]*types.ModuleAccess, 0)

	return aclAuthority
}
