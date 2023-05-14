package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/hub module sentinel errors
var (
	ErrSample                        = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidPacketTimeout          = sdkerrors.Register(ModuleName, 1500, "invalid packet timeout")
	ErrInvalidVersion                = sdkerrors.Register(ModuleName, 1501, "invalid version")
	ErrInvalidStudentAddress         = sdkerrors.Register(ModuleName, 1101, "student address is invalid: %s")
	ErrChainConfigurationAlreadyDone = sdkerrors.Register(ModuleName, 1102, "the initial configuration of the chain has already been performed")
	ErrChainConfigurationNotDone     = sdkerrors.Register(ModuleName, 1103, "the initial configuration of the chain has not yet been performed.")
	ErrKeyEnteredMismatchAdminChain  = sdkerrors.Register(ModuleName, 1104, "the key entered does not match the chain administrator's key")
	ErrWrongNameUniversity           = sdkerrors.Register(ModuleName, 1105, "the university name does not exists")
	ErrStudentNotPresent             = sdkerrors.Register(ModuleName, 1106, "the student is not present")
	ErrKeyEnteredMismatchStudent     = sdkerrors.Register(ModuleName, 1107, "the key entered does not match the student key")
	ErrIncompleteStudentInformation  = sdkerrors.Register(ModuleName, 1108, "the student must first enter all information about him/herself")
	ErrUnpaidTaxes                   = sdkerrors.Register(ModuleName, 1109, "the student has not yet paid university taxes")
	ErrNoErasmusRequest              = sdkerrors.Register(ModuleName, 1110, "the Erasmus request is absent")
	ErrPreviousRequestInProgress     = sdkerrors.Register(ModuleName, 1111, "previous Erasmus request in progress")
	ErrPreviousRequestCompleted      = sdkerrors.Register(ModuleName, 1112, "previous Erasmus request completed")
)
