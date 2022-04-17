package figleaf_test

import (
	"testing"

	keepertest "github.com/marlin-mixon/figleaf/testutil/keeper"
	"github.com/marlin-mixon/figleaf/testutil/nullify"
	"github.com/marlin-mixon/figleaf/x/figleaf"
	"github.com/marlin-mixon/figleaf/x/figleaf/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FigleafKeeper(t)
	figleaf.InitGenesis(ctx, *k, genesisState)
	got := figleaf.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
