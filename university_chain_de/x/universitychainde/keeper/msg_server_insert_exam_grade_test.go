package keeper_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"
	"university_chain_de/x/universitychainde/testutil"
	"university_chain_de/x/universitychainde/types"
	"university_chain_de/x/universitychainde/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestInsertExamGradeOk(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)

	err := os.Chdir("/university_chain_de")
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

	/*
		payTaxesResponse, err := msgServer.PayTaxes(context, &types.MsgPayTaxes{
			Creator:      testutil.Prof_ae,
			University:   "unipi",
			StudentIndex: "1",
		})

		require.Nil(t, err)
		require.EqualValues(t, types.MsgPayTaxesResponse{
			Status: 0,
		}, *payTaxesResponse)
	*/

	bank.PayTaxes(context, testutil.Mario_Rossi, testutil.Unipi, 20000)

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "25",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: 0,
	}, *insertExamGradeResponse)
}

func TestInsertExamGradeOk_2(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)

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

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "30L",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: 0,
	}, *insertExamGradeResponse)
}

func TestInsertExamGradeUnauthorisedUser(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)
	err := os.Chdir("/university_chain_de")
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

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Mario_Rossi,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "25",
	})

	require.EqualError(t,
		err,
		"the user is not authorised to enter the grade for the exam under consideration")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)

}

func TestInsertExamGradeNoTaxes(t *testing.T) {
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

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "25",
	})

	require.EqualError(t,
		err,
		"the student has not yet paid university taxes")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)

}

func TestInsertExamGradeUniversityNotFound(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)

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

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "25",
	})

	require.EqualError(t,
		err,
		"the university name does not exists")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)
}

func TestInsertExamGradeStudentNotFound(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)

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

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "2",
		ExamName:     "Algorithm engineering",
		Grade:        "25",
	})

	require.EqualError(t,
		err,
		"the student is not present")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)
}

func TestInsertExamGradeExamNotFound(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)

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

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "",
		Grade:        "25",
	})

	require.EqualError(t,
		err,
		"the exam does not exists")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)
}

func TestInsertExamGradeExamAlreadyDone(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)

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

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "25",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: 0,
	}, *insertExamGradeResponse)

	insertExamGradeResponse, err = msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "25",
	})

	require.EqualError(t,
		err,
		"the exam grade was already assigned")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)
}

func TestInsertExamGradeWrongGrade_1(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)

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

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "17",
	})

	require.EqualError(t,
		err,
		"the exam grade is wrong")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)
}

func TestInsertExamGradeWrongGrade_2(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)

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

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "31",
	})

	require.EqualError(t,
		err,
		"the exam grade is wrong")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)
}

func TestInsertExamGradeWrongGrade_3(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)

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

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "18.5",
	})

	require.EqualError(t,
		err,
		"the exam grade is wrong")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)
}

func TestInsertExamGradeNoConfig(t *testing.T) {
	msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "25",
	})

	require.EqualError(t,
		err,
		"the initial configuration of the chain has not yet been performed.")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)
}

func TestInsertExamGradeNoCompleteInfo(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)

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

	bank.PayTaxes(context, testutil.Mario_Rossi, testutil.Unipi, 20000)

	ctx := sdk.UnwrapSDKContext(context)

	student, found := keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []utilfunc.TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	require.EqualValues(t, nil, err)
	taxesData[0].Payment_made = true
	resultByteJSON, err := json.Marshal(taxesData)
	require.EqualValues(t, nil, err)
	student.TaxesData.TaxesHistory = string(resultByteJSON)
	keeper.SetStoredStudent(ctx, student)

	insertExamGradeResponse, err := msgServer.InsertExamGrade(context, &types.MsgInsertExamGrade{
		Creator:      testutil.Prof_ae,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Algorithm engineering",
		Grade:        "25",
	})

	require.EqualError(t,
		err,
		"the student must first enter all information about him/herself")
	require.EqualValues(t, types.MsgInsertExamGradeResponse{
		Status: -1,
	}, *insertExamGradeResponse)
}
