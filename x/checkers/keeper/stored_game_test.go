package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "checkers-app/testutil/keeper"
	"checkers-app/testutil/nullify"
	"checkers-app/x/checkers/keeper"
	"checkers-app/x/checkers/types"

	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredGame(keeper keeper.Keeper, ctx context.Context, n int) []types.IndexedGame {
	items := make([]types.IndexedGame, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredGame(ctx, items[i])
	}
	return items
}

func TestStoredGameGet(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredGame(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredGameRemove(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredGame(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredGame(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredGameGetAll(t *testing.T) {
	keeper, ctx := keepertest.CheckersKeeper(t)
	items := createNStoredGame(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredGame(ctx)),
	)
}
