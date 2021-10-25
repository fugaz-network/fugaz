package fugaz_test

import (
	"testing"

	keepertest "github.com/fugaz-network/fugaz/testutil/keeper"
	"github.com/fugaz-network/fugaz/x/fugaz"
	"github.com/fugaz-network/fugaz/x/fugaz/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FugazKeeper(t)
	fugaz.InitGenesis(ctx, *k, genesisState)
	got := fugaz.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
