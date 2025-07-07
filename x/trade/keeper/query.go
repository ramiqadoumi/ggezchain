package keeper

import (
	"github.com/ramiqadoumi/ggezchain/v2/x/trade/types"
)

var _ types.QueryServer = Keeper{}
