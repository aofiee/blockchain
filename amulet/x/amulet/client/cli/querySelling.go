package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
  "github.com/aofiee/amulet/x/amulet/types"
)

func GetCmdListSelling(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-selling",
		Short: "list all selling",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListSelling, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Selling\n%s\n", err.Error())
				return nil
			}
			var out []types.Selling
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
