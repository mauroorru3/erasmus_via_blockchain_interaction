package keeper_test

import (
	"log"
	"os"
	"testing"
	"university_chain_it/x/universitychainit/testutil"
	"university_chain_it/x/universitychainit/types"

	"github.com/stretchr/testify/require"
)

func TestInsertStudentContactInfoOk(t *testing.T) {
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

	contactInfoResponse, err := msgServer.InsertStudentContactInfo(context, &types.MsgInsertStudentContactInfo{
		Creator:        testutil.Mario_Rossi,
		University:     "unipi",
		StudentIndex:   "1",
		ContactAddress: "via roma",
		Email:          "mario.rossi@example.it",
		MobilePhone:    "0000000000",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgInsertStudentContactInfoResponse{
		Status: 0,
	}, *contactInfoResponse)
}

func TestInsertStudentContactInfoWrongUniversity(t *testing.T) {
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

	contactInfoResponse, err := msgServer.InsertStudentContactInfo(context, &types.MsgInsertStudentContactInfo{
		Creator:        testutil.Mario_Rossi,
		University:     "aaaaa",
		StudentIndex:   "1",
		ContactAddress: "via roma",
		Email:          "mario.rossi@example.it",
		MobilePhone:    "0000000000",
	})

	require.EqualError(t,
		err,
		"the university name does not exists")
	require.EqualValues(t, types.MsgInsertStudentContactInfoResponse{
		Status: -1,
	}, *contactInfoResponse)
}

func TestInsertStudentContactInfoWrongStudentId(t *testing.T) {
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

	contactInfoResponse, err := msgServer.InsertStudentContactInfo(context, &types.MsgInsertStudentContactInfo{
		Creator:        testutil.Mario_Rossi,
		University:     "unipi",
		StudentIndex:   "2",
		ContactAddress: "via roma",
		Email:          "mario.rossi@example.it",
		MobilePhone:    "0000000000",
	})

	require.EqualError(t,
		err,
		"the student is not present")
	require.EqualValues(t, types.MsgInsertStudentContactInfoResponse{
		Status: -1,
	}, *contactInfoResponse)
}

func TestInsertStudentContactInfoWrongEmail(t *testing.T) {
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

	contactInfoResponse, err := msgServer.InsertStudentContactInfo(context, &types.MsgInsertStudentContactInfo{
		Creator:        testutil.Mario_Rossi,
		University:     "unipi",
		StudentIndex:   "1",
		ContactAddress: "via roma",
		Email:          "aaaaa",
		MobilePhone:    "0000000000",
	})

	require.EqualError(t,
		err,
		"mail: missing '@' or angle-addr")
	require.EqualValues(t, types.MsgInsertStudentContactInfoResponse{
		Status: -1,
	}, *contactInfoResponse)
}

func TestInsertStudentContactInfoWrongMobileNumber(t *testing.T) {
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

	contactInfoResponse, err := msgServer.InsertStudentContactInfo(context, &types.MsgInsertStudentContactInfo{
		Creator:        testutil.Mario_Rossi,
		University:     "unipi",
		StudentIndex:   "1",
		ContactAddress: "via roma",
		Email:          "mario.rossi@example.it",
		MobilePhone:    "000000000000000000",
	})

	require.EqualError(t,
		err,
		"the student mobile number is wrong")
	require.EqualValues(t, types.MsgInsertStudentContactInfoResponse{
		Status: -1,
	}, *contactInfoResponse)
}
