package cli

import (
	"strconv"

	"university_chain_it/x/universitychainit/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdInsertStudentPersonalInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "insert-student-personal-info [university] [student-index] [gender] [date-of-birth] [primary-nationality] [country-of-birth] [province-of-birth] [town-of-birth] [tax-code] [incomeBracket]",
		Short: "Broadcast message insert_student_personal_info",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUniversity := args[0]
			argStudentIndex := args[1]
			argGender := args[2]
			argDateOfBirth := args[3]
			argPrimaryNationality := args[4]
			argCountryOfBirth := args[5]
			argProvinceOfBirth := args[6]
			argTownOfBirth := args[7]
			argTaxCode := args[8]
			argIncomeBracket, err := cast.ToUint32E(args[9])
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgInsertStudentPersonalInfo(
				clientCtx.GetFromAddress().String(),
				argUniversity,
				argStudentIndex,
				argGender,
				argDateOfBirth,
				argPrimaryNationality,
				argCountryOfBirth,
				argProvinceOfBirth,
				argTownOfBirth,
				argTaxCode,
				argIncomeBracket,
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
