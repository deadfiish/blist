package cli

import (
	"context"
	"github.com/bianjieai/iritamod/modules/blist/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	blistQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      " blist ",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	blacklistQueryCmd := &cobra.Command{
		Use:                        types.BListName,
		Short:                      " blacklist ",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	blacklistQueryCmd.AddCommand(
		GetBListByName(),
		GetBList(),
	)
	blistQueryCmd.AddCommand(blacklistQueryCmd)
	return blistQueryCmd
}

func GetBListByName() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blist [name]",
		Short: "query a name",
		Long:  "query a name",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			req := &types.QueryBListByNameRequest{}
			res, err := queryClient.QueryBListByName(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}
func GetBList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all",
		Short: "query blist",
		Long:  "query blist",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			req := &types.QueryBListRequest{}
			res, err := queryClient.QueryBList(context.Background(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}
	return cmd
}