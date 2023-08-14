package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"university_chain_it/x/universitychainit/types"
)

func CmdListForeignUniversities() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-foreign-universities",
		Short: "list all foreign_universities",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllForeignUniversitiesRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.ForeignUniversitiesAll(context.Background(), params)
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

func CmdShowForeignUniversities() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-foreign-universities [university-name]",
		Short: "shows a foreign_universities",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argUniversityName := args[0]

			params := &types.QueryGetForeignUniversitiesRequest{
				UniversityName: argUniversityName,
			}

			res, err := queryClient.ForeignUniversities(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
