package cli

import (
	"bufio"
  
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/aofiee/amulet/x/amulet/types"
)

func GetCmdCreateBuying(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-buying [AmuletHash] [BuyingPrice] [SellingBy] [BuyingBy]",
		Short: "Creates a new buying",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsAmuletHash := string(args[0])
      argsBuyingPrice := string(args[1])
      argsSellingBy := string(args[2])
      argsBuyingBy := string(args[3])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateBuying(cliCtx.GetFromAddress(), argsAmuletHash, argsBuyingPrice, argsSellingBy, argsBuyingBy)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
