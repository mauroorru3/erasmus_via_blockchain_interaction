package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"university_chain_it/x/universitychainit/types"
)

var _ = strconv.Itoa(0)

func CmdStartErasmus() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start-erasmus [university] [student-index]",
		Short: "Broadcast message start_erasmus",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUniversity := args[0]
			argStudentIndex := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgStartErasmus(
				clientCtx.GetFromAddress().String(),
				argUniversity,
				argStudentIndex,
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
