package keeper

import (
	"github.com/fugaz-network/fugaz/x/fugaz/types"
)

var _ types.QueryServer = Keeper{}
