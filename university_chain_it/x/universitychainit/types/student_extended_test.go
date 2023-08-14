package types_test

import (
	"testing"
	"university_chain_it/x/universitychainit/testutil"
	"university_chain_it/x/universitychainit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func GetStoredStudent1() types.StoredStudent {
	return types.StoredStudent{
		Index: "1",
		StudentData: &types.StudentInfo{
			Name:                     "Mario",
			Surname:                  "Rossi",
			CourseType:               "0",
			CourseOfStudy:            "CS",
			Status:                   "1",
			CurrentYearOfStudy:       1,
			OutOfCourse:              false,
			NumberOfYearsOutOfCourse: 0,
			StudentKey:               testutil.Mario_Rossi,
		},
		TranscriptData: &types.TranscriptOfRecords{},
		PersonalData:   &types.PersonalInfo{},
		ResidenceData:  &types.ResidenceInfo{},
		ContactData:    &types.ContactInfo{},
		TaxesData:      &types.TaxesInfo{},
		ErasmusData:    &types.ErasmusInfo{},
	}
}

func TestCanGetAddress(t *testing.T) {
	aliceAddress, err1 := sdk.AccAddressFromBech32(testutil.Mario_Rossi)
	studentAddress, err2 := GetStoredStudent1().GetStudentAddress()
	require.Equal(t, aliceAddress, studentAddress)
	require.Nil(t, err2)
	require.Nil(t, err1)
}

func TestGetAddressWrong(t *testing.T) {
	storedStudent := GetStoredStudent1()
	storedStudent.StudentData.StudentKey = "cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4" // Bad last digit
	studentAddress, err := storedStudent.GetStudentAddress()
	require.Nil(t, studentAddress)
	require.EqualError(t,
		err,
		"student address is invalid: cosmos1jmjfq0tplp9tmx4v9uemw72y4d2wa5nr3xn9d4: decoding bech32 failed: invalid checksum (expected 3xn9d3 got 3xn9d4)")
	require.EqualError(t, storedStudent.Validate(), err.Error())
}

func TestStudentValidateOk(t *testing.T) {
	storedStudent := GetStoredStudent1()
	require.NoError(t, storedStudent.Validate())
}
