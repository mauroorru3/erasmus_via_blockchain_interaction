package universitychainit

import (
	"university_chain_it/x/universitychainit/keeper"
	"university_chain_it/x/universitychainit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.StudentInfo != nil {
		k.SetStudentInfo(ctx, *genState.StudentInfo)
	}
	// Set if defined
	if genState.TranscriptOfRecords != nil {
		k.SetTranscriptOfRecords(ctx, *genState.TranscriptOfRecords)
	}
	// Set if defined
	if genState.PersonalInfo != nil {
		k.SetPersonalInfo(ctx, *genState.PersonalInfo)
	}
	// Set if defined
	if genState.ResidenceInfo != nil {
		k.SetResidenceInfo(ctx, *genState.ResidenceInfo)
	}
	// Set if defined
	if genState.ContactInfo != nil {
		k.SetContactInfo(ctx, *genState.ContactInfo)
	}
	// Set if defined
	if genState.TaxesInfo != nil {
		k.SetTaxesInfo(ctx, *genState.TaxesInfo)
	}
	// Set if defined
	if genState.ErasmusInfo != nil {
		k.SetErasmusInfo(ctx, *genState.ErasmusInfo)
	}
	// Set if defined
	k.SetChainInfo(ctx, genState.ChainInfo)
	// Set all the foreignUniversities
	for _, elem := range genState.ForeignUniversitiesList {
		k.SetForeignUniversities(ctx, elem)
	}
	// Set all the universityInfo
	for _, elem := range genState.UniversityInfoList {
		k.SetUniversityInfo(ctx, elem)
	}
	// Set all the professorsExams
	for _, elem := range genState.ProfessorsExamsList {
		k.SetProfessorsExams(ctx, elem)
	}
	// Set all the storedStudent
	for _, elem := range genState.StoredStudentList {
		k.SetStoredStudent(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	// Get all studentInfo
	studentInfo, found := k.GetStudentInfo(ctx)
	if found {
		genesis.StudentInfo = &studentInfo
	}
	// Get all transcriptOfRecords
	transcriptOfRecords, found := k.GetTranscriptOfRecords(ctx)
	if found {
		genesis.TranscriptOfRecords = &transcriptOfRecords
	}
	// Get all personalInfo
	personalInfo, found := k.GetPersonalInfo(ctx)
	if found {
		genesis.PersonalInfo = &personalInfo
	}
	// Get all residenceInfo
	residenceInfo, found := k.GetResidenceInfo(ctx)
	if found {
		genesis.ResidenceInfo = &residenceInfo
	}
	// Get all contactInfo
	contactInfo, found := k.GetContactInfo(ctx)
	if found {
		genesis.ContactInfo = &contactInfo
	}
	// Get all taxesInfo
	taxesInfo, found := k.GetTaxesInfo(ctx)
	if found {
		genesis.TaxesInfo = &taxesInfo
	}
	// Get all erasmusInfo
	erasmusInfo, found := k.GetErasmusInfo(ctx)
	if found {
		genesis.ErasmusInfo = &erasmusInfo
	}
	// Get all chainInfo
	chainInfo, found := k.GetChainInfo(ctx)
	if found {
		genesis.ChainInfo = chainInfo
	}
	genesis.ForeignUniversitiesList = k.GetAllForeignUniversities(ctx)
	genesis.UniversityInfoList = k.GetAllUniversityInfo(ctx)
	genesis.ProfessorsExamsList = k.GetAllProfessorsExams(ctx)
	genesis.StoredStudentList = k.GetAllStoredStudent(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
