package types_test

import (
	"testing"

	"hub/x/hub/types"

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
					Name:                     "56",
					Surname:                  "60",
					CourseType:               35,
					CourseOfStudy:            "16",
					Status:                   74,
					CurrentYearOfStudy:       85,
					OutOfCourse:              true,
					NumberOfYearsOutOfCourse: 89,
					StudentKey:               "80",
				},
				TranscriptOfRecords: &types.TranscriptOfRecords{
					ExamsData:       "17",
					TotalExams:      18,
					ExamsPassed:     52,
					TotalCredits:    83,
					AchievedCredits: 87,
				},
				PersonalInfo: &types.PersonalInfo{
					Gender:             31,
					DateOfBirth:        "63",
					PrimaryNationality: "77",
					CountryOfBirth:     "28",
					ProvinceOfBirth:    "55",
					TownOfBirth:        "42",
					TaxCode:            "40",
				},
				ResidenceInfo: &types.ResidenceInfo{
					Country:     "32",
					Province:    "39",
					Town:        "61",
					PostCode:    "80",
					Address:     "58",
					HouseNumber: "71",
					HomePhone:   "81",
				},
				ContactInfo: &types.ContactInfo{
					ContactAddress: "90",
					Email:          "60",
					MobilePhone:    "27",
				},
				TaxesInfo: &types.TaxesInfo{
					Status:       false,
					TotalAmount:  73,
					TaxesHistory: "17",
				},
				ErasmusInfo: &types.ErasmusInfo{
					ErasmusStudent:      "51",
					NumberTimes:         16,
					NumberMonths:        14,
					TotalExams:          48,
					ExamsPassed:         80,
					TotalCredits:        48,
					AchievedCredits:     69,
					Career:              "82",
					PreviousStudentFifo: "75",
					NextStudentFifo:     "93",
				},
				StoredStudentList: []types.StoredStudent{
					{
						Index: "0",
					},
					{
						Index: "1",
					},
				},
				ChainInfo: types.ChainInfo{
					ChainKey:              "38",
					ChainAdministratorKey: "81",
					InitStatus:            false,
				},
				UniversitiesList: []types.Universities{
					{
						UniversityName: "0",
					},
					{
						UniversityName: "1",
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
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
		{
			desc: "duplicated universities",
			genState: &types.GenesisState{
				UniversitiesList: []types.Universities{
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
