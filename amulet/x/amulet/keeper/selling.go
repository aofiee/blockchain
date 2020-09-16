package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/aofiee/amulet/x/amulet/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateSelling(ctx sdk.Context, selling types.Selling) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.SellingPrefix + selling.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(selling)
	store.Set(key, value)
}

func listSelling(ctx sdk.Context, k Keeper) ([]byte, error) {
  var sellingList []types.Selling
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.SellingPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var selling types.Selling
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &selling)
    sellingList = append(sellingList, selling)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, sellingList)
  return res, nil
}