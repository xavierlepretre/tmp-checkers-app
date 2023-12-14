package keeper

import (
	"checkers-app/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
