package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/marlin-mixon/figleaf/testutil/keeper"
	"github.com/marlin-mixon/figleaf/x/figleaf/keeper"
	"github.com/marlin-mixon/figleaf/x/figleaf/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FigleafKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
