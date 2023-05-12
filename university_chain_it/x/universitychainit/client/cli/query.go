package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"university_chain_it/x/universitychainit/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group universitychainit queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdShowStudentInfo())
	cmd.AddCommand(CmdShowTranscriptOfRecords())
	cmd.AddCommand(CmdShowPersonalInfo())
	cmd.AddCommand(CmdShowResidenceInfo())
	cmd.AddCommand(CmdShowContactInfo())
	cmd.AddCommand(CmdShowTaxesInfo())
	cmd.AddCommand(CmdShowErasmusInfo())
	cmd.AddCommand(CmdShowChainInfo())
	cmd.AddCommand(CmdListForeignUniversities())
	cmd.AddCommand(CmdShowForeignUniversities())
	cmd.AddCommand(CmdListUniversityInfo())
	cmd.AddCommand(CmdShowUniversityInfo())
	cmd.AddCommand(CmdListProfessorsExams())
	cmd.AddCommand(CmdShowProfessorsExams())
	cmd.AddCommand(CmdListStoredStudent())
	cmd.AddCommand(CmdShowStoredStudent())
	// this line is used by starport scaffolding # 1

	return cmd
}
