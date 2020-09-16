package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/aofiee/amulet/x/amulet/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

func (k Keeper) CreateAmulet(ctx sdk.Context, amulet types.Amulet) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.AmuletPrefix + amulet.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(amulet)
	store.Set(key, value)
}

func listAmulet(ctx sdk.Context, k Keeper) ([]byte, error) {
  var amuletList []types.Amulet
  store := ctx.KVStore(k.storeKey)
  iterator := sdk.KVStorePrefixIterator(store, []byte(types.AmuletPrefix))
  for ; iterator.Valid(); iterator.Next() {
    var amulet types.Amulet
    k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &amulet)
    amuletList = append(amuletList, amulet)
  }
  res := codec.MustMarshalJSONIndent(k.cdc, amuletList)
  return res, nil
}