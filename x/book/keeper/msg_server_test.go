package keeper_test

import (
	"context"
	"testing"

	keepertest "book/testutil/keeper"
	"book/x/book/keeper"
	"book/x/book/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BookKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
