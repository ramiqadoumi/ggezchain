package keeper

import (
	"github.com/ramiqadoumi/ggezchain/v2/x/acl/types"
)

var _ types.QueryServer = Keeper{}
