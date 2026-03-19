package types_test

import (
	"testing"

	"university_chain_it/x/universitychainit/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				PortId: types.PortID,
				StudentInfo: &types.StudentInfo{
					Name:                     "69",
					Surname:                  "71",
					CourseType:               "7",
					CourseOfStudy:            "10",
					Status:                   "14",
					CurrentYearOfStudy:       67,
					OutOfCourse:              true,
					NumberOfYearsOutOfCourse: 25,
					StudentKey:               "3",
				},
				TranscriptOfRecords: &types.TranscriptOfRecords{
					ExamsData:       "69",
					TotalExams:      55,
					ExamsPassed:     53,
					TotalCredits:    39,
					AchievedCredits: 32,
				},
				PersonalInfo: &types.PersonalInfo{
					Gender:             "18",
					DateOfBirth:        "58",
					PrimaryNationality: "73",
					CountryOfBirth:     "97",
					ProvinceOfBirth:    "53",
					TownOfBirth:        "7",
					TaxCode:            "83",
				},
				ResidenceInfo: &types.ResidenceInfo{
					Country:     "84",
					Province:    "11",
					Town:        "61",
					PostCode:    "30",
					Address:     "13",
					HouseNumber: "77",
					HomePhone:   "77",
				},
				ContactInfo: &types.ContactInfo{
					ContactAddress: "89",
					Email:          "6",
					MobilePhone:    "63",
				},
				TaxesInfo: &types.TaxesInfo{
					Status:       false,
					TotalAmount:  63,
					TaxesHistory: "85",
				},
				ErasmusInfo: &types.ErasmusInfo{
					ErasmusStudent:      "57",
					NumberTimes:         69,
					NumberMonths:        57,
					TotalExams:          87,
					ExamsPassed:         51,
					TotalCredits:        7,
					AchievedCredits:     36,
					Career:              "55",
					PreviousStudentFifo: "42",
					NextStudentFifo:     "21",
				},
				ChainInfo: types.ChainInfo{
					HubKey:   "96",
					ChainKey: "54",
					Country:  "27",
				},
				ForeignUniversitiesList: []types.ForeignUniversities{
					{
						UniversityName: "0",
					},
					{
						UniversityName: "1",
					},
				},
				UniversityInfoList: []types.UniversityInfo{
					{
						UniversityName: "0",
					},
					{
						UniversityName: "1",
					},
				},
				ProfessorsExamsList: []types.ProfessorsExams{
					{
						ExamName: "0",
					},
					{
						ExamName: "1",
					},
				},
				StoredStudentList: []types.StoredStudent{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated foreignUniversities",
			genState: &types.GenesisState{
				ForeignUniversitiesList: []types.ForeignUniversities{
					{
						UniversityName: "0",
					},
					{
						UniversityName: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated universityInfo",
			genState: &types.GenesisState{
				UniversityInfoList: []types.UniversityInfo{
					{
						UniversityName: "0",
					},
					{
						UniversityName: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated professorsExams",
			genState: &types.GenesisState{
				ProfessorsExamsList: []types.ProfessorsExams{
					{
						ExamName: "0",
					},
					{
						ExamName: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated storedStudent",
			genState: &types.GenesisState{
				StoredStudentList: []types.StoredStudent{
					{
						Index: "0",
					},
					{
						Index: "0",
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestDefaultGenesisState_ExpectedInitialValues(t *testing.T) {
	require.EqualValues(t,
		&types.GenesisState{
			PortId:              "universitychainit",
			StudentInfo:         nil,
			TranscriptOfRecords: nil,
			PersonalInfo:        nil,
			ResidenceInfo:       nil,
			ContactInfo:         nil,
			TaxesInfo:           nil,
			ErasmusInfo:         nil,
			ChainInfo: types.ChainInfo{
				HubKey:                "cosmos1nqg8gn5kdvs3na2psm9fp2sad7yka8tmh90dpd",
				ChainKey:              "cosmos1lw4azenln0hdr09q6ckmyg4vac46vdxpw9yrcc",
				ChainAdministratorKey: "cosmos1cc90wwv6k78yrp6x59et2jhj03ryxy8jq54cxs",
				Country:               "Italy",
				InitStatus:            false,
			},
			ForeignUniversitiesList: []types.ForeignUniversities{},
			UniversityInfoList:      []types.UniversityInfo{},
			ProfessorsExamsList:     []types.ProfessorsExams{},
			StoredStudentList:       []types.StoredStudent{},
			Params:                  types.DefaultParams(),
		},
		types.DefaultGenesis())
}
