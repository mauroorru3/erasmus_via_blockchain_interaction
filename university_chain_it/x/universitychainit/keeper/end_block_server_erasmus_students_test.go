package keeper_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"
	"university_chain_it/x/universitychainit/testutil"
	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestTerminateExpiredErasmusPeriods(t *testing.T) {
	msgServer, keeper, context, _, bank := setupMsgServerConfigureChain(t)
	//msgServer, _, context, _, _ := setupMsgServerConfigureChain(t)

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
	require.EqualValues(t, types.MsgInsertErasmusExamResponse{
		Status: 0,
	}, *insertExamGradeResponse)

	insertErasmusRequestResponse, err := msgServer.InsertErasmusRequest(context, &types.MsgInsertErasmusRequest{
		Creator:               testutil.Mario_Rossi,
		University:            "unipi",
		StudentIndex:          "1",
		DurationInMonths:      "6",
		ForeignUniversityName: "tum",
		ErasmusType:           "study",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgInsertErasmusRequestResponse{
		Status: 0,
	}, *insertErasmusRequestResponse)

	insertErasmusExamResponse, err := msgServer.InsertErasmusExam(context, &types.MsgInsertErasmusExam{
		Creator:      testutil.Mario_Rossi,
		University:   "unipi",
		StudentIndex: "1",
		ExamName:     "Advanced databases",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgInsertErasmusExamResponse{
		Status: 0,
	}, *insertErasmusExamResponse)

	// Commented code because of the bank module

	startErasmusResponse, err := msgServer.StartErasmus(context, &types.MsgStartErasmus{
		Creator:      testutil.Mario_Rossi,
		University:   "unipi",
		StudentIndex: "1",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgStartErasmusResponse{
		Status: 0,
	}, *startErasmusResponse)

	_, found = keeper.GetStoredStudent(ctx, "unipi_1")
	require.True(t, found)

	systemInfo, found := keeper.GetUniversityInfo(ctx, "unipi")
	require.True(t, found)
	require.EqualValues(t, types.UniversityInfo{
		UniversityName:  "unipi",
		NextStudentId:   2,
		SecretariatKey:  "cosmos1sf3zdq9q7tzfk9l83g05padqky4v9tld2dhsur",
		UniversityKey:   "cosmos1mhmthq4cx8ltfexahqrhutqxu9d26zltwhh39n",
		FifoHeadErasmus: "unipi_1",
		FifoTailErasmus: "unipi_1",
		DeadlineTaxes:   "2023-12-31 12:00:00",
		DeadlineErasmus: "2023-12-31 12:00:00",
		TaxesBrackets:   "{\"first\":[6500,0],\"second\":[28000,200],\"third\":[60000,400],\"fourth\":800}",
		MaxErasmusExams: 3,
		CaiKey:          "",
	}, systemInfo)

	keeper.TerminateExpiredErasmusPeriods(context)

	systemInfo, found = keeper.GetUniversityInfo(ctx, "unipi")
	require.True(t, found)
	require.EqualValues(t, types.UniversityInfo{

		UniversityName:  "unipi",
		NextStudentId:   2,
		SecretariatKey:  "cosmos1sf3zdq9q7tzfk9l83g05padqky4v9tld2dhsur",
		UniversityKey:   "cosmos1mhmthq4cx8ltfexahqrhutqxu9d26zltwhh39n",
		FifoHeadErasmus: "",
		FifoTailErasmus: "",
		DeadlineTaxes:   "2023-12-31 12:00:00",
		DeadlineErasmus: "2023-12-31 12:00:00",
		TaxesBrackets:   "{\"first\":[6500,0],\"second\":[28000,200],\"third\":[60000,400],\"fourth\":800}",
		MaxErasmusExams: 3,
		CaiKey:          "",
	}, systemInfo)

}
