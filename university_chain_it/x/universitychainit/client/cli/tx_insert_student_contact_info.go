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

func CmdInsertStudentContactInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insert-student-contact-info [university] [student-index] [contact-address] [email] [mobile-phone]",
		Short: "Broadcast message insert_student_contact_info",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUniversity := args[0]
			argStudentIndex := args[1]
			argContactAddress := args[2]
			argEmail := args[3]
			argMobilePhone := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInsertStudentContactInfo(
				clientCtx.GetFromAddress().String(),
				argUniversity,
				argStudentIndex,
				argContactAddress,
				argEmail,
				argMobilePhone,
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
