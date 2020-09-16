package amulet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/aofiee/amulet/x/amulet/types"
  "github.com/aofiee/amulet/x/amulet/keeper"
  	"github.com/tendermint/tendermint/crypto"
)

func handleMsgCreateAmulet(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateAmulet) (*sdk.Result, error) {
	var amulet = types.Amulet{
		Creator: msg.Creator,
		ID:      msg.ID,
    Name: msg.Name,
    Identity: msg.Identity,
    Address: msg.Address,
    Telephon: msg.Telephon,
    AmuletType: msg.AmuletType,
    AmuletName: msg.AmuletName,
    From: msg.From,
    Model: msg.Model,
    Surface: msg.Surface,
    Year: msg.Year,
    Province: msg.Province,
    Description: msg.Description,
    Remark: msg.Remark,
  }
  moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	payment, _ := sdk.ParseCoins("1rune")
	if err := k.CoinKeeper.SendCoins(ctx, amulet.Creator, moduleAcct, payment); err != nil {
		return nil, err
	}
	k.CreateAmulet(ctx, amulet)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
