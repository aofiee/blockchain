package amulet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/aofiee/amulet/x/amulet/types"
	"github.com/aofiee/amulet/x/amulet/keeper"
	  	"github.com/tendermint/tendermint/crypto"
)

func handleMsgCreateBuying(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateBuying) (*sdk.Result, error) {
	var buying = types.Buying{
		Creator: msg.Creator,
		ID:      msg.ID,
    AmuletHash: msg.AmuletHash,
    BuyingPrice: msg.BuyingPrice,
    SellingBy: msg.SellingBy,
    BuyingBy: msg.BuyingBy,
	}
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	payment, _ := sdk.ParseCoins("100gold")
	if err := k.CoinKeeper.SendCoins(ctx, buying.Creator, moduleAcct, payment); err != nil {
		return nil, err
	}
	k.CreateBuying(ctx, buying)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
