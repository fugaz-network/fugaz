package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/fugaz-network/fugaz/testutil/keeper"
	"github.com/fugaz-network/fugaz/x/fugaz/keeper"
	"github.com/fugaz-network/fugaz/x/fugaz/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FugazKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
