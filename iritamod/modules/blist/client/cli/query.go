package cli

import (
	"context"
	"github.com/bianjieai/iritamod/modules/blist/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	BListQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      " blist ",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	BListQueryCmd.AddCommand(
		GetBListByName(),
		GetBList(),
	)
	return BListQueryCmd
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
		Args:  cobra.ExactArgs(1),
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