package keeper_test

import (
	"log"
	"os"
	"testing"
	"university_chain_it/x/universitychainit/testutil"
	"university_chain_it/x/universitychainit/types"

	"github.com/stretchr/testify/require"
)

func TestInsertStudentResidenceInfoOk(t *testing.T) {
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

	residenceInfoResponse, err := msgServer.InsertStudentResidenceInfo(context, &types.MsgInsertStudentResidenceInfo{
		Creator:      testutil.Mario_Rossi,
		University:   "unipi",
		StudentIndex: "1",
		Country:      "italy",
		Province:     "PI",
		Town:         "Pisa",
		PostCode:     "56100",
		Address:      "via roma",
		HouseNumber:  "3",
		HomePhone:    "0000000000",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgInsertStudentResidenceInfoResponse{
		Status: 0,
	}, *residenceInfoResponse)

}

func TestInsertStudentResidenceInfoWrongUniversityName(t *testing.T) {
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

	residenceInfoResponse, err := msgServer.InsertStudentResidenceInfo(context, &types.MsgInsertStudentResidenceInfo{
		Creator:      testutil.Mario_Rossi,
		University:   "aaaa",
		StudentIndex: "1",
		Country:      "italy",
		Province:     "PI",
		Town:         "Pisa",
		PostCode:     "56100",
		Address:      "via roma",
		HouseNumber:  "3",
		HomePhone:    "0000000000",
	})

	require.EqualError(t,
		err,
		"the university name does not exists")
	require.EqualValues(t, types.MsgInsertStudentResidenceInfoResponse{
		Status: -1,
	}, *residenceInfoResponse)

}

func TestInsertStudentResidenceInfoWrongStudentId(t *testing.T) {
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

	residenceInfoResponse, err := msgServer.InsertStudentResidenceInfo(context, &types.MsgInsertStudentResidenceInfo{
		Creator:      testutil.Mario_Rossi,
		University:   "unipi",
		StudentIndex: "2",
		Country:      "italy",
		Province:     "PI",
		Town:         "Pisa",
		PostCode:     "56100",
		Address:      "via roma",
		HouseNumber:  "3",
		HomePhone:    "0000000000",
	})

	require.EqualError(t,
		err,
		"the student is not present")
	require.EqualValues(t, types.MsgInsertStudentResidenceInfoResponse{
		Status: -1,
	}, *residenceInfoResponse)

}

func TestInsertStudentResidenceInfoWrongMobileNumber(t *testing.T) {
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

	residenceInfoResponse, err := msgServer.InsertStudentResidenceInfo(context, &types.MsgInsertStudentResidenceInfo{
		Creator:      testutil.Mario_Rossi,
		University:   "unipi",
		StudentIndex: "1",
		Country:      "italy",
		Province:     "PI",
		Town:         "Pisa",
		PostCode:     "56100",
		Address:      "via roma",
		HouseNumber:  "3",
		HomePhone:    "0000000000000000",
	})

	require.EqualError(t,
		err,
		"the student mobile number is wrong")
	require.EqualValues(t, types.MsgInsertStudentResidenceInfoResponse{
		Status: -1,
	}, *residenceInfoResponse)

}
