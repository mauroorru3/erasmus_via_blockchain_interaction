package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {

	/*
		uniInfo, err := utilfunc.ReadUniversityInfo()
		if err != nil {
			os.Exit(-1)
		}
	*/

	return &GenesisState{
		PortId:              PortID,
		StudentInfo:         nil,
		TranscriptOfRecords: nil,
		PersonalInfo:        nil,
		ResidenceInfo:       nil,
		ContactInfo:         nil,
		TaxesInfo:           nil,
		ErasmusInfo:         nil,
		ChainInfo: ChainInfo{
			HubKey:                "cosmos1nqg8gn5kdvs3na2psm9fp2sad7yka8tmh90dpd",
			ChainAdministratorKey: "cosmos1cc90wwv6k78yrp6x59et2jhj03ryxy8jq54cxs",
			ChainKey:              "cosmos1lw4azenln0hdr09q6ckmyg4vac46vdxpw9yrcc",
			Country:               "Italy",
			InitStatus:            false,
		},
		ForeignUniversitiesList: []ForeignUniversities{},
		UniversityInfoList:      []UniversityInfo{},
		ProfessorsExamsList:     []ProfessorsExams{},
		StoredStudentList:       []StoredStudent{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated index in foreignUniversities
	foreignUniversitiesIndexMap := make(map[string]struct{})

	for _, elem := range gs.ForeignUniversitiesList {
		index := string(ForeignUniversitiesKey(elem.UniversityName))
		if _, ok := foreignUniversitiesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for foreignUniversities")
		}
		foreignUniversitiesIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in universityInfo
	universityInfoIndexMap := make(map[string]struct{})

	for _, elem := range gs.UniversityInfoList {
		index := string(UniversityInfoKey(elem.UniversityName))
		if _, ok := universityInfoIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for universityInfo")
		}
		universityInfoIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in professorsExams
	professorsExamsIndexMap := make(map[string]struct{})

	for _, elem := range gs.ProfessorsExamsList {
		index := string(ProfessorsExamsKey(elem.ExamName))
		if _, ok := professorsExamsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for professorsExams")
		}
		professorsExamsIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in storedStudent
	storedStudentIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredStudentList {
		index := string(StoredStudentKey(elem.Index))
		if _, ok := storedStudentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedStudent")
		}
		storedStudentIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
