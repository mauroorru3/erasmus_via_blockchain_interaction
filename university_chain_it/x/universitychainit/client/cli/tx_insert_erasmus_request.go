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

func CmdInsertErasmusRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insert-erasmus-request [university] [student-index] [duration-in-months] [foreign-university-name] [erasmus-type]",
		Short: "Broadcast message insert_erasmus_request",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUniversity := args[0]
			argStudentIndex := args[1]
			argDurationInMonths := args[2]
			argForeignUniversityName := args[3]
			argErasmusType := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInsertErasmusRequest(
				clientCtx.GetFromAddress().String(),
				argUniversity,
				argStudentIndex,
				argDurationInMonths,
				argForeignUniversityName,
				argErasmusType,
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
