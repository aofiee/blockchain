package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/aofiee/amulet/x/amulet/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateBuying(ctx sdk.Context, buying types.Buying) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.BuyingPrefix + buying.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(buying)
	store.Set(key, value)
}

func listBuying(ctx sdk.Context, k Keeper) ([]byte, error) {
  var buyingList []types.Buying
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.BuyingPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var buying types.Buying
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &buying)
    buyingList = append(buyingList, buying)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, buyingList)
  return res, nil
}