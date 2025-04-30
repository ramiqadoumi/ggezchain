package types

import (
	"encoding/json"
	"slices"
	"strings"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ValidateModuleAccessList takes a JSON string of module accesses, validates it,
// and returns a structured slice of ModuleAccess or an error if invalid.
func ValidateModuleAccessList(moduleAccessListStr string) ([]*ModuleAccess, error) {
	var moduleAccessList []*ModuleAccess
	if err := json.Unmarshal([]byte(moduleAccessListStr), &moduleAccessList); err != nil {
		return nil, ErrInvalidModuleAccessList
	}
	seenModules := make(map[string]bool)

	for _, moduleAccess := range moduleAccessList {
		moduleAccess.Module = strings.ToLower(strings.TrimSpace(moduleAccess.Module))

		if moduleAccess.Module == "" {
			return nil, ErrInvalidModuleName.Wrapf("empty module not allowed")
		}

		if seenModules[moduleAccess.Module] {
			return nil, ErrInvalidModuleName.Wrapf("%s is duplicated module", moduleAccess.Module)
		}
		seenModules[moduleAccess.Module] = true

		// todo: check if that condition necessary
		if !moduleAccess.IsMaker && !moduleAccess.IsChecker {
			return nil, sdkerrors.ErrInvalidRequest.Wrapf("at least one of is_maker or is_checker must be true for module %s", moduleAccess.Module)
		}
	}

	return moduleAccessList, nil
}

// ValidateSingleModuleAccess takes a JSON string for one ModuleAccess object, validates it,
// and returns the structured object or an error.
func ValidateSingleModuleAccess(moduleAccessStr string) (*ModuleAccess, error) {
	var moduleAccess ModuleAccess
	if err := json.Unmarshal([]byte(moduleAccessStr), &moduleAccess); err != nil {
		return nil, ErrInvalidModuleAccessObject
	}

	moduleAccess.Module = strings.ToLower(strings.TrimSpace(moduleAccess.Module))

	if moduleAccess.Module == "" {
		return nil, ErrInvalidModuleName.Wrapf("empty module not allowed")
	}

	if !moduleAccess.IsMaker && !moduleAccess.IsChecker {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf(
			"at least one of is_maker or is_checker must be true for module %s",
			moduleAccess.Module)
	}

	return &moduleAccess, nil
}

// GetUpdatedModuleAccessList finds a module by name (case-insensitive) and updates its roles in-place within the current list.
func GetUpdatedModuleAccessList(currentList []*ModuleAccess, update *ModuleAccess) []*ModuleAccess {
	for _, m := range currentList {
		if strings.EqualFold(m.Module, update.Module) {
			m.IsMaker = update.IsMaker
			m.IsChecker = update.IsChecker
			break
		}
	}
	return currentList
}

// GetAuthorityModules retrieves the module names from a list of ModuleAccess.
func GetAuthorityModules(moduleAccessList []*ModuleAccess) []string {

	if len(moduleAccessList) == 0 {
		return nil
	}

	modules := make([]string, 0, len(moduleAccessList))
	for _, moduleAccess := range moduleAccessList {
		if moduleAccess.Module != "" {
			modules = append(modules, moduleAccess.Module)
		}
	}

	return modules
}

// HasModule check if retrieved module exist in currentModules
func HasModule(module string, currentModules []string) bool {
	if module == "" {
		return false
	}

	normalizedModule := strings.ToLower(strings.TrimSpace(module))

	return slices.Contains(currentModules, normalizedModule)
}

// ValidateUpdateModuleAccess validate updated module
func ValidateUpdateModuleAccess(aclAuthority AclAuthority, updatedModule string) error {
	modules := GetAuthorityModules(aclAuthority.ModuleAccess)
	if !HasModule(updatedModule, modules) {
		return ErrModuleNotExist.Wrapf("%s module not exist", updatedModule)
	}
	return nil
}

// ValidateAddModuleAccess validates adding new modules, checking for empty input and existing duplicates.
func ValidateAddModuleAccess(currentModules, newModules []string) error {
	if len(newModules) == 0 {
		return ErrEmptyModuleAccessList.Wrapf("new module list cannot be empty")
	}

	for _, module := range newModules {
		if HasModule(module, currentModules) {
			return ErrModuleExist.Wrapf("%s module already exists", module)
		}
	}

	return nil
}

// ValidateDeleteModuleAccess validates module names for removal and returns the updated access list if valid.
func ValidateDeleteModuleAccess(moduleNames []string, moduleAccessList []*ModuleAccess) ([]*ModuleAccess, error) {
	if len(moduleNames) == 0 {
		return nil, ErrInvalidModuleName.Wrapf("at least one module name must be provided")
	}

	if len(moduleAccessList) == 0 {
		return nil, ErrEmptyModuleAccessList
	}

	moduleNameSet := make(map[string]struct{}, len(moduleNames))
	for _, name := range moduleNames {
		normalizedName := strings.ToLower(strings.TrimSpace(name))
		if normalizedName == "" {
			return nil, ErrInvalidModuleName.Wrapf("module name cannot be empty")
		}
		moduleNameSet[normalizedName] = struct{}{}
	}

	currentModules := make(map[string]struct{}, len(moduleAccessList))
	for _, module := range moduleAccessList {
		currentModules[module.Module] = struct{}{}
	}

	// check if module not exist in current modules
	var missingModules []string
	for name := range moduleNameSet {
		if _, exists := currentModules[name]; !exists {
			missingModules = append(missingModules, name)
		}
	}

	if len(missingModules) > 0 {
		return nil, ErrModuleNotExist.Wrapf("module(s) not found: %s", strings.Join(missingModules, ", "))
	}

	updatedList := make([]*ModuleAccess, 0, len(moduleAccessList))
	for _, module := range moduleAccessList {
		if _, shouldRemove := moduleNameSet[module.Module]; shouldRemove {
			continue
		}
		updatedList = append(updatedList, module)
	}

	return updatedList, nil
}

// ValidateModuleOverlap check if there is common modules between added or updated modules and removed modules
func ValidateModuleOverlap(modules, removedModules []string) error {
	removedSet := make(map[string]struct{}, len(removedModules))
	for _, module := range removedModules {
		normalized := strings.ToLower(strings.TrimSpace(module))
		if normalized != "" {
			removedSet[normalized] = struct{}{}
		}
	}

	// check if any added/updated module exists in the removed set
	for _, module := range modules {
		normalized := strings.ToLower(strings.TrimSpace(module))
		if _, exists := removedSet[normalized]; exists {
			return ErrUpdateAndRemoveModule.Wrapf("%q module", normalized)
		}
	}

	return nil
}

// ValidateAndExtractModuleNames validates a ModuleAccess JSON string and extracts the list of module names.
func ValidateAndExtractModuleNames(addModuleAccess string) ([]string, error) {
	moduleAccessList, err := ValidateModuleAccessList(addModuleAccess)
	if err != nil {
		return nil, err
	}
	modules := GetAuthorityModules(moduleAccessList)
	return modules, nil
}

// ValidateConflictBetweenModuleAccess validate update add and remove flags
func ValidateConflictBetweenModuleAccess(updateModuleAccess string, addModuleAccess string, removeList []string) error {
	if updateModuleAccess != "" && len(removeList) > 0 {
		updateModuleAccess, err := ValidateSingleModuleAccess(updateModuleAccess)
		if err != nil {
			return err
		}

		if err = ValidateModuleOverlap([]string{updateModuleAccess.Module}, removeList); err != nil {
			return err
		}
	}

	if addModuleAccess != "" && len(removeList) > 0 {
		addedModules, err := ValidateAndExtractModuleNames(addModuleAccess)
		if err != nil {
			return err
		}

		if err = ValidateModuleOverlap(addedModules, removeList); err != nil {
			return err
		}
	}
	return nil
}

// ValidateJSONFormat validate JSON format
func ValidateJSONFormat(jsonStr string, fieldName string) error {
	if jsonStr != "" {
		if !json.Valid([]byte(jsonStr)) {
			return sdkerrors.ErrInvalidRequest.Wrapf("invalid JSON format for field '%s'", fieldName)
		}
	}
	return nil
}
