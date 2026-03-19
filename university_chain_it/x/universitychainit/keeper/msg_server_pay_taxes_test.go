package keeper_test

import (
	"log"
	"os"
	"testing"
	"university_chain_it/x/universitychainit/testutil"
	"university_chain_it/x/universitychainit/types"

	"github.com/stretchr/testify/require"
)

func TestPayTaxesOk(t *testing.T) {
	msgServer, _, context, _, bank := setupMsgServerConfigureChain(t)
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

	bank.PayTaxes(context, testutil.Mario_Rossi, testutil.Unipi, 20000)

}
