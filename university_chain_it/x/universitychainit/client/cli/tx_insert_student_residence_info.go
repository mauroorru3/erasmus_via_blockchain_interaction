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

func CmdInsertStudentResidenceInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insert-student-residence-info [university] [student-index] [country] [province] [town] [post-code] [address] [house-number] [home-phone]",
		Short: "Broadcast message insert_student_residence_info",
		Args:  cobra.ExactArgs(9),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUniversity := args[0]
			argStudentIndex := args[1]
			argCountry := args[2]
			argProvince := args[3]
			argTown := args[4]
			argPostCode := args[5]
			argAddress := args[6]
			argHouseNumber := args[7]
			argHomePhone := args[8]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInsertStudentResidenceInfo(
				clientCtx.GetFromAddress().String(),
				argUniversity,
				argStudentIndex,
				argCountry,
				argProvince,
				argTown,
				argPostCode,
				argAddress,
				argHouseNumber,
				argHomePhone,
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
