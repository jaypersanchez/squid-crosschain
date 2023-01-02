package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "squid-crosschain/testutil/keeper"
	"squid-crosschain/x/squidcrosschain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SquidcrosschainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
