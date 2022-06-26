package book_test

import (
	"testing"

	keepertest "book/testutil/keeper"
	"book/testutil/nullify"
	"book/x/book"
	"book/x/book/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BookKeeper(t)
	book.InitGenesis(ctx, *k, genesisState)
	got := book.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
