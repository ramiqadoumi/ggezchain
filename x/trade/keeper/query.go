package keeper

import (
	"github.com/ramiqadoumi/ggezchain/x/trade/types"
)

var _ types.QueryServer = Keeper{}
