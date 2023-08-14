package keeper_test

import (
	"log"
	"os"
	"testing"
	"university_chain_it/x/universitychainit/testutil"
	"university_chain_it/x/universitychainit/types"

	"github.com/stretchr/testify/require"
)

func TestInsertStudentPersonalInfoOk(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)
	err := os.Chdir("/university_chain_it")
	if err != nil {
		log.Println(err)
	}

	configureChainResponse, err := msgServer.ConfigureChain(context, &types.MsgConfigureChain{
		Creator: testutil.Chain_admin,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgConfigureChainResponse{
		Status: 0,
	}, *configureChainResponse)

	newStudentResponse, err := msgServer.RegisterNewStudent(context, &types.MsgRegisterNewStudent{
		Creator:        testutil.Mario_Rossi,
		University:     "unipi",
		Name:           "Mario",
		Surname:        "Rossi",
		CourseType:     "master",
		CourseOfStudy:  "cs",
		DepartmentName: "Computer Science",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgRegisterNewStudentResponse{
		StudentIndex: "1",
	}, *newStudentResponse)

	personalInfoResponse, err := msgServer.InsertStudentPersonalInfo(context, &types.MsgInsertStudentPersonalInfo{
		Creator:            testutil.Mario_Rossi,
		University:         "unipi",
		StudentIndex:       "1",
		Gender:             "male",
		DateOfBirth:        "1994-06-06",
		PrimaryNationality: "italian",
		CountryOfBirth:     "italy",
		ProvinceOfBirth:    "Rome",
		TownOfBirth:        "Rome",
		TaxCode:            "1111111111111111",
		IncomeBracket:      20000,
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgInsertStudentPersonalInfoResponse{
		Status: 0,
	}, *personalInfoResponse)

}

func TestInsertStudentPersonalInfoWrongUniversityName(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)

	configureChainResponse, err := msgServer.ConfigureChain(context, &types.MsgConfigureChain{
		Creator: testutil.Chain_admin,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgConfigureChainResponse{
		Status: 0,
	}, *configureChainResponse)

	newStudentResponse, err := msgServer.RegisterNewStudent(context, &types.MsgRegisterNewStudent{
		Creator:        testutil.Mario_Rossi,
		University:     "unipi",
		Name:           "Mario",
		Surname:        "Rossi",
		CourseType:     "master",
		CourseOfStudy:  "cs",
		DepartmentName: "Computer Science",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgRegisterNewStudentResponse{
		StudentIndex: "1",
	}, *newStudentResponse)

	personalInfoResponse, err := msgServer.InsertStudentPersonalInfo(context, &types.MsgInsertStudentPersonalInfo{
		Creator:            testutil.Mario_Rossi,
		University:         "aaaa",
		StudentIndex:       "1",
		Gender:             "male",
		DateOfBirth:        "1994-06-06",
		PrimaryNationality: "italian",
		CountryOfBirth:     "italy",
		ProvinceOfBirth:    "Rome",
		TownOfBirth:        "Rome",
		TaxCode:            "1111111111111111",
		IncomeBracket:      20000,
	})

	require.EqualError(t,
		err,
		"the university name does not exists")
	require.EqualValues(t, types.MsgInsertStudentPersonalInfoResponse{
		Status: -1,
	}, *personalInfoResponse)

}

func TestInsertStudentPersonalInfoWrongStudentId(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)

	configureChainResponse, err := msgServer.ConfigureChain(context, &types.MsgConfigureChain{
		Creator: testutil.Chain_admin,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgConfigureChainResponse{
		Status: 0,
	}, *configureChainResponse)

	newStudentResponse, err := msgServer.RegisterNewStudent(context, &types.MsgRegisterNewStudent{
		Creator:        testutil.Mario_Rossi,
		University:     "unipi",
		Name:           "Mario",
		Surname:        "Rossi",
		CourseType:     "master",
		CourseOfStudy:  "cs",
		DepartmentName: "Computer Science",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgRegisterNewStudentResponse{
		StudentIndex: "1",
	}, *newStudentResponse)

	personalInfoResponse, err := msgServer.InsertStudentPersonalInfo(context, &types.MsgInsertStudentPersonalInfo{
		Creator:            testutil.Mario_Rossi,
		University:         "unipi",
		StudentIndex:       "2",
		Gender:             "male",
		DateOfBirth:        "1994-06-06",
		PrimaryNationality: "italian",
		CountryOfBirth:     "italy",
		ProvinceOfBirth:    "Rome",
		TownOfBirth:        "Rome",
		TaxCode:            "1111111111111111",
		IncomeBracket:      20000,
	})

	require.EqualError(t,
		err,
		"the student is not present")
	require.EqualValues(t, types.MsgInsertStudentPersonalInfoResponse{
		Status: -1,
	}, *personalInfoResponse)

}

func TestInsertStudentPersonalInfoWrongGender(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)

	configureChainResponse, err := msgServer.ConfigureChain(context, &types.MsgConfigureChain{
		Creator: testutil.Chain_admin,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgConfigureChainResponse{
		Status: 0,
	}, *configureChainResponse)

	newStudentResponse, err := msgServer.RegisterNewStudent(context, &types.MsgRegisterNewStudent{
		Creator:        testutil.Mario_Rossi,
		University:     "unipi",
		Name:           "Mario",
		Surname:        "Rossi",
		CourseType:     "master",
		CourseOfStudy:  "cs",
		DepartmentName: "Computer Science",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgRegisterNewStudentResponse{
		StudentIndex: "1",
	}, *newStudentResponse)

	personalInfoResponse, err := msgServer.InsertStudentPersonalInfo(context, &types.MsgInsertStudentPersonalInfo{
		Creator:            testutil.Mario_Rossi,
		University:         "unipi",
		StudentIndex:       "1",
		Gender:             "aaaaa",
		DateOfBirth:        "1994-06-06",
		PrimaryNationality: "italian",
		CountryOfBirth:     "italy",
		ProvinceOfBirth:    "Rome",
		TownOfBirth:        "Rome",
		TaxCode:            "1111111111111111",
		IncomeBracket:      20000,
	})

	require.EqualError(t,
		err,
		"the student gender is not found")
	require.EqualValues(t, types.MsgInsertStudentPersonalInfoResponse{
		Status: -1,
	}, *personalInfoResponse)

}

func TestInsertStudentPersonalInfoWrongDate(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)

	configureChainResponse, err := msgServer.ConfigureChain(context, &types.MsgConfigureChain{
		Creator: testutil.Chain_admin,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgConfigureChainResponse{
		Status: 0,
	}, *configureChainResponse)

	newStudentResponse, err := msgServer.RegisterNewStudent(context, &types.MsgRegisterNewStudent{
		Creator:        testutil.Mario_Rossi,
		University:     "unipi",
		Name:           "Mario",
		Surname:        "Rossi",
		CourseType:     "master",
		CourseOfStudy:  "cs",
		DepartmentName: "Computer Science",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgRegisterNewStudentResponse{
		StudentIndex: "1",
	}, *newStudentResponse)

	personalInfoResponse, err := msgServer.InsertStudentPersonalInfo(context, &types.MsgInsertStudentPersonalInfo{
		Creator:            testutil.Mario_Rossi,
		University:         "unipi",
		StudentIndex:       "1",
		Gender:             "male",
		DateOfBirth:        "2022-06-06",
		PrimaryNationality: "italian",
		CountryOfBirth:     "italy",
		ProvinceOfBirth:    "Rome",
		TownOfBirth:        "Rome",
		TaxCode:            "1111111111111111",
		IncomeBracket:      20000,
	})

	require.EqualError(t,
		err,
		"the student birth date is wrong")
	require.EqualValues(t, types.MsgInsertStudentPersonalInfoResponse{
		Status: -1,
	}, *personalInfoResponse)

}

func TestInsertStudentPersonalInfoWrongTaxCode(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)

	configureChainResponse, err := msgServer.ConfigureChain(context, &types.MsgConfigureChain{
		Creator: testutil.Chain_admin,
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgConfigureChainResponse{
		Status: 0,
	}, *configureChainResponse)

	newStudentResponse, err := msgServer.RegisterNewStudent(context, &types.MsgRegisterNewStudent{
		Creator:        testutil.Mario_Rossi,
		University:     "unipi",
		Name:           "Mario",
		Surname:        "Rossi",
		CourseType:     "master",
		CourseOfStudy:  "cs",
		DepartmentName: "Computer Science",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgRegisterNewStudentResponse{
		StudentIndex: "1",
	}, *newStudentResponse)

	personalInfoResponse, err := msgServer.InsertStudentPersonalInfo(context, &types.MsgInsertStudentPersonalInfo{
		Creator:            testutil.Mario_Rossi,
		University:         "unipi",
		StudentIndex:       "1",
		Gender:             "male",
		DateOfBirth:        "1994-06-06",
		PrimaryNationality: "italian",
		CountryOfBirth:     "italy",
		ProvinceOfBirth:    "Rome",
		TownOfBirth:        "Rome",
		TaxCode:            "1111111111111111111",
		IncomeBracket:      20000,
	})

	require.EqualError(t,
		err,
		"the student tax code is wrong")
	require.EqualValues(t, types.MsgInsertStudentPersonalInfoResponse{
		Status: -1,
	}, *personalInfoResponse)

}
