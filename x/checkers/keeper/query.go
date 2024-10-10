package keeper

import (
	"github.com/ibondarev/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
