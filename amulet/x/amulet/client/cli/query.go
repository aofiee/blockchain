package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/aofiee/amulet/x/amulet/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group amulet queries under a subcommand
	amuletQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	amuletQueryCmd.AddCommand(
		flags.GetCommands(
      // this line is used by starport scaffolding
			GetCmdListBuying(queryRoute, cdc),
			GetCmdListSelling(queryRoute, cdc),
			GetCmdListAmulet(queryRoute, cdc),
		)...,
	)

	return amuletQueryCmd
}
