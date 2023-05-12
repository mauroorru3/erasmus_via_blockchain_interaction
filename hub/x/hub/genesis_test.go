package hub_test

import (
	"testing"

	keepertest "hub/testutil/keeper"
	"hub/testutil/nullify"
	"hub/x/hub"
	"hub/x/hub/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		StudentInfo: &types.StudentInfo{
			Name:                     "5",
			Surname:                  "42",
			CourseType:               36,
			CourseOfStudy:            "52",
			Status:                   32,
			CurrentYearOfStudy:       65,
			OutOfCourse:              false,
			NumberOfYearsOutOfCourse: 49,
			StudentKey:               "32",
		},
		TranscriptOfRecords: &types.TranscriptOfRecords{
			ExamsData:       "62",
			TotalExams:      32,
			ExamsPassed:     14,
			TotalCredits:    89,
			AchievedCredits: 92,
		},
		PersonalInfo: &types.PersonalInfo{
			Gender:             76,
			DateOfBirth:        "35",
			PrimaryNationality: "39",
			CountryOfBirth:     "67",
			ProvinceOfBirth:    "6",
			TownOfBirth:        "47",
			TaxCode:            "53",
		},
		ResidenceInfo: &types.ResidenceInfo{
			Country:     "94",
			Province:    "37",
			Town:        "36",
			PostCode:    "60",
			Address:     "78",
			HouseNumber: "45",
			HomePhone:   "27",
		},
		ContactInfo: &types.ContactInfo{
			ContactAddress: "46",
			Email:          "61",
			MobilePhone:    "8",
		},
		TaxesInfo: &types.TaxesInfo{
			Status:       true,
			TotalAmount:  36,
			TaxesHistory: "36",
		},
		ErasmusInfo: &types.ErasmusInfo{
			ErasmusStudent:      "11",
			NumberTimes:         19,
			NumberMonths:        89,
			TotalExams:          27,
			ExamsPassed:         24,
			TotalCredits:        90,
			AchievedCredits:     79,
			Career:              "17",
			PreviousStudentFifo: "58",
			NextStudentFifo:     "1",
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
			ChainKey:              "8",
			ChainAdministratorKey: "85",
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
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HubKeeper(t)
	hub.InitGenesis(ctx, *k, genesisState)
	got := hub.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.Equal(t, genesisState.StudentInfo, got.StudentInfo)
	require.Equal(t, genesisState.TranscriptOfRecords, got.TranscriptOfRecords)
	require.Equal(t, genesisState.PersonalInfo, got.PersonalInfo)
	require.Equal(t, genesisState.ResidenceInfo, got.ResidenceInfo)
	require.Equal(t, genesisState.ContactInfo, got.ContactInfo)
	require.Equal(t, genesisState.TaxesInfo, got.TaxesInfo)
	require.Equal(t, genesisState.ErasmusInfo, got.ErasmusInfo)
	require.ElementsMatch(t, genesisState.StoredStudentList, got.StoredStudentList)
	require.Equal(t, genesisState.ChainInfo, got.ChainInfo)
	require.ElementsMatch(t, genesisState.UniversitiesList, got.UniversitiesList)
	// this line is used by starport scaffolding # genesis/test/assert
}
