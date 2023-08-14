package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"hub/x/hub/types"
)

func CmdShowTranscriptOfRecords() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-transcript-of-records",
		Short: "shows transcript_of_records",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetTranscriptOfRecordsRequest{}

			res, err := queryClient.TranscriptOfRecords(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
