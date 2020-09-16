package keeper

import (
  // this line is used by starport scaffolding
	"github.com/aofiee/amulet/x/amulet/types"
		
	
		
	
		
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for amulet clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
    // this line is used by starport scaffolding # 2
		case types.QueryListBuying:
			return listBuying(ctx, k)
		case types.QueryListSelling:
			return listSelling(ctx, k)
		case types.QueryListAmulet:
			return listAmulet(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown amulet query endpoint")
		}
	}
}