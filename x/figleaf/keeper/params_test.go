package keeper_test

import (
	"testing"

	testkeeper "github.com/marlin-mixon/figleaf/testutil/keeper"
	"github.com/marlin-mixon/figleaf/x/figleaf/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FigleafKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
