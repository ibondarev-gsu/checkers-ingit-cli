package test_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/ibondarev/checkers/testutil/keeper"
	"github.com/ibondarev/checkers/x/checkers/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CheckersKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
