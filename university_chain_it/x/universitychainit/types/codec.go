package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgConfigureChain{}, "universitychainit/ConfigureChain", nil)
	cdc.RegisterConcrete(&MsgRegisterNewStudent{}, "universitychainit/RegisterNewStudent", nil)
	cdc.RegisterConcrete(&MsgInsertStudentPersonalInfo{}, "universitychainit/InsertStudentPersonalInfo", nil)
	cdc.RegisterConcrete(&MsgInsertStudentContactInfo{}, "universitychainit/InsertStudentContactInfo", nil)
	cdc.RegisterConcrete(&MsgInsertStudentResidenceInfo{}, "universitychainit/InsertStudentResidenceInfo", nil)
	cdc.RegisterConcrete(&MsgInsertExamGrade{}, "universitychainit/InsertExamGrade", nil)
	cdc.RegisterConcrete(&MsgPayTaxes{}, "universitychainit/PayTaxes", nil)
	cdc.RegisterConcrete(&MsgInsertErasmusRequest{}, "universitychainit/InsertErasmusRequest", nil)
	cdc.RegisterConcrete(&MsgInsertErasmusExam{}, "universitychainit/InsertErasmusExam", nil)
	cdc.RegisterConcrete(&MsgStartErasmus{}, "universitychainit/StartErasmus", nil)
	cdc.RegisterConcrete(&MsgSendErasmusStudent{}, "universitychainit/SendErasmusStudent", nil)
	cdc.RegisterConcrete(&MsgSendEndErasmusPeriodRequest{}, "universitychainit/SendEndErasmusPeriodRequest", nil)
	cdc.RegisterConcrete(&MsgEndErasmusBeforeDeadline{}, "universitychainit/EndErasmusBeforeDeadline", nil)
	cdc.RegisterConcrete(&MsgExtendErasmus{}, "universitychainit/ExtendErasmus", nil)
	cdc.RegisterConcrete(&MsgSendExtendErasmusPeriod{}, "universitychainit/SendExtendErasmusPeriod", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConfigureChain{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterNewStudent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInsertStudentPersonalInfo{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInsertStudentContactInfo{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInsertStudentResidenceInfo{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInsertExamGrade{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPayTaxes{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInsertErasmusRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInsertErasmusExam{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStartErasmus{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendErasmusStudent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendEndErasmusPeriodRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEndErasmusBeforeDeadline{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgExtendErasmus{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendExtendErasmusPeriod{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
