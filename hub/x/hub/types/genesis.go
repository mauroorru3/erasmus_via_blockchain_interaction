package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:              PortID,
		StudentInfo:         nil,
		TranscriptOfRecords: nil,
		PersonalInfo:        nil,
		ResidenceInfo:       nil,
		ContactInfo:         nil,
		TaxesInfo:           nil,
		ErasmusInfo:         nil,
		StoredStudentList:   []StoredStudent{},
		ChainInfo: ChainInfo{
			ChainAdministratorKey: "cosmos1w68naxagtnzfs7jrldjaalgcgfrctrk06y2rep",
			ChainKey:              "cosmos1nqg8gn5kdvs3na2psm9fp2sad7yka8tmh90dpd",
			InitStatus:            false,
		},
		UniversitiesList: []Universities{},
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
	// Check for duplicated index in storedStudent
	storedStudentIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredStudentList {
		index := string(StoredStudentKey(elem.Index))
		if _, ok := storedStudentIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedStudent")
		}
		storedStudentIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in universities
	universitiesIndexMap := make(map[string]struct{})

	for _, elem := range gs.UniversitiesList {
		index := string(UniversitiesKey(elem.UniversityName))
		if _, ok := universitiesIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for universities")
		}
		universitiesIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
