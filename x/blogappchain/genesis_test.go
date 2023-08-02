package blogappchain_test

import (
	"testing"

	keepertest "blog-appchain/testutil/keeper"
	"blog-appchain/testutil/nullify"
	"blog-appchain/x/blogappchain"
	"blog-appchain/x/blogappchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogappchainKeeper(t)
	blogappchain.InitGenesis(ctx, *k, genesisState)
	got := blogappchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
