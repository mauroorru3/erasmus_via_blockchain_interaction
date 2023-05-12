package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"hub/x/hub/types"
)

func CmdListUniversities() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-universities",
		Short: "list all universities",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllUniversitiesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.UniversitiesAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowUniversities() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-universities [university-name]",
		Short: "shows a universities",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argUniversityName := args[0]

			params := &types.QueryGetUniversitiesRequest{
				UniversityName: argUniversityName,
			}

			res, err := queryClient.Universities(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
