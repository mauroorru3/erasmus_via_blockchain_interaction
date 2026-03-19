package keeper

import (
	"hub/x/hub/types"
)

var _ types.QueryServer = Keeper{}
