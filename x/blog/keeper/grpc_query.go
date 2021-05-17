package keeper

import (
	"github.com/fadeev/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
