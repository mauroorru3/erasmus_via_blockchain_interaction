package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"university_chain_de/x/universitychainde/types"
)

var _ = strconv.Itoa(0)

func CmdExtendErasmus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "extend-erasmus [university] [student-index] [duration-in-months]",
		Short: "Broadcast message extend_erasmus",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUniversity := args[0]
			argStudentIndex := args[1]
			argDurationInMonths := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgExtendErasmus(
				clientCtx.GetFromAddress().String(),
				argUniversity,
				argStudentIndex,
				argDurationInMonths,
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
