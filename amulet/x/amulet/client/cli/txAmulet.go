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

func GetCmdCreateAmulet(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-amulet [Name] [Identity] [Address] [Telephon] [AmuletType] [AmuletName] [From] [Model] [Surface] [Year] [Province] [Description] [Remark]",
		Short: "Creates a new amulet",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
      argsName := string(args[0])
      argsIdentity := string(args[1])
      argsAddress := string(args[2])
      argsTelephon := string(args[3])
      argsAmuletType := string(args[4])
      argsAmuletName := string(args[5])
      argsFrom := string(args[6])
      argsModel := string(args[7])
      argsSurface := string(args[8])
      argsYear := string(args[9])
      argsProvince := string(args[10])
      argsDescription := string(args[11])
      argsRemark := string(args[12])
      
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateAmulet(cliCtx.GetFromAddress(), argsName, argsIdentity, argsAddress, argsTelephon, argsAmuletType, argsAmuletName, argsFrom, argsModel, argsSurface, argsYear, argsProvince, argsDescription, argsRemark)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
