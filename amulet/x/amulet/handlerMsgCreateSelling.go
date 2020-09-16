package amulet

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/aofiee/amulet/x/amulet/types"
	"github.com/aofiee/amulet/x/amulet/keeper"
)

func handleMsgCreateSelling(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateSelling) (*sdk.Result, error) {
	var selling = types.Selling{
		Creator: msg.Creator,
		ID:      msg.ID,
    AmuletHash: msg.AmuletHash,
    SellingPrice: msg.SellingPrice,
    SellingBy: msg.SellingBy,
	}
	k.CreateSelling(ctx, selling)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
