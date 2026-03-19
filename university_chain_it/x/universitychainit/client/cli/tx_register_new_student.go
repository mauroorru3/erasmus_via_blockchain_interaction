package cli

import (
	"strconv"

	"university_chain_it/x/universitychainit/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdRegisterNewStudent() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-new-student [university] [name] [surname] [course-type] [course-of-study] [department-name]",
		Short: "Broadcast message register_new_student",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argUniversity := args[0]
			argName := args[1]
			argSurname := args[2]
			argCourseType := args[3]
			argCourseOfStudy := args[4]
			argDepartmentName := args[5]
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegisterNewStudent(
				clientCtx.GetFromAddress().String(),
				argUniversity,
				argName,
				argSurname,
				argCourseType,
				argCourseOfStudy,
				argDepartmentName,
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
