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

func CmdInsertErasmusExam() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insert-erasmus-exam [university] [student-index] [exam-name]",
		Short: "Broadcast message insert_erasmus_exam",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUniversity := args[0]
			argStudentIndex := args[1]
			argExamName := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInsertErasmusExam(
				clientCtx.GetFromAddress().String(),
				argUniversity,
				argStudentIndex,
				argExamName,
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
