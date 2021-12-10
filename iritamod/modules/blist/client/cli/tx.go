package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/bianjieai/iritamod/modules/blist/types"
)

func NewTxCmd() *cobra.Command {
	BListTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "blist transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	BListTxCmd.AddCommand(
		AddToBList(),
		RemoveFromBList(),
	)
	return BListTxCmd
}

func AddToBList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add [name]",
		Args:  cobra.ExactArgs(1),
		Short: "add name to blist",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			name := args[0]
			msg := types.NewMsgAddToBList(
				name,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func RemoveFromBList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [name]",
		Args:  cobra.ExactArgs(1),
		Short: "remove name from blist",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			name := args[0]
			msg := types.NewMsgRemoveFromBList(
				name,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
