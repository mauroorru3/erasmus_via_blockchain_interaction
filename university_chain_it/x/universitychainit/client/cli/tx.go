package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"university_chain_it/x/universitychainit/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdConfigureChain())
	cmd.AddCommand(CmdRegisterNewStudent())
	cmd.AddCommand(CmdInsertStudentPersonalInfo())
	cmd.AddCommand(CmdInsertStudentContactInfo())
	cmd.AddCommand(CmdInsertStudentResidenceInfo())
	cmd.AddCommand(CmdInsertExamGrade())
	cmd.AddCommand(CmdPayTaxes())
	cmd.AddCommand(CmdInsertErasmusRequest())
	cmd.AddCommand(CmdInsertErasmusExam())
	cmd.AddCommand(CmdStartErasmus())
	cmd.AddCommand(CmdSendErasmusStudent())
	// this line is used by starport scaffolding # 1

	return cmd
}
