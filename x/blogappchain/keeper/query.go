package keeper

import (
	"blog-appchain/x/blogappchain/types"
)

var _ types.QueryServer = Keeper{}
