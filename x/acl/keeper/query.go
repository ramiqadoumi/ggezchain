package keeper

import (
	"github.com/ramiqadoumi/ggezchain/x/acl/types"
)

var _ types.QueryServer = Keeper{}
