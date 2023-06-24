package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/universitychainit module sentinel errors
var (
	ErrSample                          = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout            = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion                  = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrInvalidStudentAddress           = sdkerrors.Register(ModuleName, 1101, "student address is invalid: %s")
	ErrChainConfigurationAlreadyDone   = sdkerrors.Register(ModuleName, 1102, "the initial configuration of the chain has already been performed")
	ErrChainConfigurationNotDone       = sdkerrors.Register(ModuleName, 1103, "the initial configuration of the chain has not yet been performed.")
	ErrKeyEnteredMismatchAdminChain    = sdkerrors.Register(ModuleName, 1104, "the key entered does not match the chain administrator's key")
	ErrWrongNameUniversity             = sdkerrors.Register(ModuleName, 1105, "the university name does not exists")
	ErrWrongCourseType                 = sdkerrors.Register(ModuleName, 1106, "the course type does not exists")
	ErrWrongCourseOfStudy              = sdkerrors.Register(ModuleName, 1107, "the course of study does not exists")
	ErrWrongDepartment                 = sdkerrors.Register(ModuleName, 1108, "the department does not exists")
	ErrStudentAlreadyPresent           = sdkerrors.Register(ModuleName, 1109, "the student is already present")
	ErrStudentNotPresent               = sdkerrors.Register(ModuleName, 1110, "the student is not present")
	ErrGenderNotFound                  = sdkerrors.Register(ModuleName, 1111, "the student gender is not found")
	ErrWrongDate                       = sdkerrors.Register(ModuleName, 1112, "the student birth date is wrong")
	ErrWrongTaxCode                    = sdkerrors.Register(ModuleName, 1113, "the student tax code is wrong")
	ErrWrongMobileNumber               = sdkerrors.Register(ModuleName, 1114, "the student mobile number is wrong")
	ErrWrongExamName                   = sdkerrors.Register(ModuleName, 1115, "the exam does not exists")
	ErrUnauthorisedUser                = sdkerrors.Register(ModuleName, 1116, "the user is not authorised to enter the grade for the exam under consideration")
	ErrWrongExamGrade                  = sdkerrors.Register(ModuleName, 1117, "the exam grade is wrong")
	ErrGradeAlreadyAssigned            = sdkerrors.Register(ModuleName, 1118, "the exam grade was already assigned")
	ErrIncompleteStudentInformation    = sdkerrors.Register(ModuleName, 1119, "the student must first enter all information about him/herself")
	ErrStudentCannotPay                = sdkerrors.Register(ModuleName, 1120, "the student cannot pay the taxes")
	ErrCannotPayTaxes                  = sdkerrors.Register(ModuleName, 1121, "the bank module cannot transfer the taxes")
	ErrNoTaxesToPay                    = sdkerrors.Register(ModuleName, 1122, "the student does not have to pay taxes")
	ErrTaxesAlreadyPayed               = sdkerrors.Register(ModuleName, 1123, "the student has already paid taxes")
	ErrUnpaidTaxes                     = sdkerrors.Register(ModuleName, 1124, "the student has not yet paid university taxes")
	ErrKeyEnteredMismatchStudent       = sdkerrors.Register(ModuleName, 1125, "the key entered does not match the student key")
	ErrWrongErasmusDuration            = sdkerrors.Register(ModuleName, 1126, "the Erasmus duration entered is wrong")
	ErrLimitErasmusMonthsExceeded      = sdkerrors.Register(ModuleName, 1127, "the Erasmus duration exceeds the total of 12 months of the study cycle")
	ErrInvalidErasmusType              = sdkerrors.Register(ModuleName, 1128, "the Erasmus type is invalid")
	ErrWrongForeignUniversity          = sdkerrors.Register(ModuleName, 1129, "wrong foreign university")
	ErrTooFewCFUToAttend               = sdkerrors.Register(ModuleName, 1130, "the number of CFUs are too few to participate")
	ErrErasmusConfFileCourseType       = sdkerrors.Register(ModuleName, 1131, "error in the configuration file: course type is wrong")
	ErrErasmusConfFileCourseOfStudy    = sdkerrors.Register(ModuleName, 1132, "error in the configuration file: course of study is wrong")
	ErrErasmusConfFileErasmusType      = sdkerrors.Register(ModuleName, 1133, "error in the configuration file: Erasmus type is wrong")
	ErrPreviousRequestStartup          = sdkerrors.Register(ModuleName, 1134, "previous Erasmus request at start-up")
	ErrPreviousRequestInProgress       = sdkerrors.Register(ModuleName, 1135, "previous Erasmus request in progress")
	ErrNoErasmusRequest                = sdkerrors.Register(ModuleName, 1136, "the Erasmus request is absent")
	ErrPreviousRequestCompleted        = sdkerrors.Register(ModuleName, 1137, "previous Erasmus request completed")
	ErrErasmusExamNotFound             = sdkerrors.Register(ModuleName, 1138, "Erasmus exam not found")
	ErrExamAlreadyHeld                 = sdkerrors.Register(ModuleName, 1139, "the exam has already been held")
	ErrErasmusExamAlreadyInserted      = sdkerrors.Register(ModuleName, 1140, "the erasmus exam has already been inserted in the plan")
	ErrMaxNumberErasmusExamReached     = sdkerrors.Register(ModuleName, 1141, "maximum number of erasmus exam reached")
	ErrErasmusDeadlineExceeded         = sdkerrors.Register(ModuleName, 1142, "erasmus deadline exceeded")
	ErrUniversityCannotPay             = sdkerrors.Register(ModuleName, 1143, "the university cannot pay the Erasmus contribution")
	ErrCannotPayErasmusContribution    = sdkerrors.Register(ModuleName, 1144, "the bank module cannot transfer the Erasmus contribution")
	ErrNoErasmusContributionToPay      = sdkerrors.Register(ModuleName, 1145, "the university does not have to pay Erasmus contribution")
	ErrErasmusContributionAlreadyPayed = sdkerrors.Register(ModuleName, 1146, "the university has already paid the Erasmus contribution")
	ErrNonCallableFunction             = sdkerrors.Register(ModuleName, 1147, "this function cannot be called; the logic of the function is implemented within other functions")
	ErrOutgoingPeriod                  = sdkerrors.Register(ModuleName, 1148, "the student is currently doing the period abroad")
	ErrCompletedIncomingPeriod         = sdkerrors.Register(ModuleName, 1149, "the incoming Erasmus period is completed")
	ErrCompletedOutgoingPeriod         = sdkerrors.Register(ModuleName, 1150, "the outgoing Erasmus period is completed")
	ErrIncomingPeriod                  = sdkerrors.Register(ModuleName, 1151, "it is not possible to perform this operation in the destination university")
)
