package keeper

import (
	"book/x/book/types"
)

var _ types.QueryServer = Keeper{}
