package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgConfigureChain{}, "universitychainde/ConfigureChain", nil)
	cdc.RegisterConcrete(&MsgRegisterNewStudent{}, "universitychainde/RegisterNewStudent", nil)
	cdc.RegisterConcrete(&MsgInsertStudentPersonalInfo{}, "universitychainde/InsertStudentPersonalInfo", nil)
	cdc.RegisterConcrete(&MsgInsertStudentContactInfo{}, "universitychainde/InsertStudentContactInfo", nil)
	cdc.RegisterConcrete(&MsgInsertStudentResidenceInfo{}, "universitychainde/InsertStudentResidenceInfo", nil)
	cdc.RegisterConcrete(&MsgInsertExamGrade{}, "universitychainde/InsertExamGrade", nil)
	cdc.RegisterConcrete(&MsgPayTaxes{}, "universitychainde/PayTaxes", nil)
	cdc.RegisterConcrete(&MsgInsertErasmusRequest{}, "universitychainde/InsertErasmusRequest", nil)
	cdc.RegisterConcrete(&MsgInsertErasmusExam{}, "universitychainde/InsertErasmusExam", nil)
	cdc.RegisterConcrete(&MsgStartErasmus{}, "universitychainde/StartErasmus", nil)
	cdc.RegisterConcrete(&MsgSendErasmusStudent{}, "universitychainde/SendErasmusStudent", nil)
	cdc.RegisterConcrete(&MsgSendEndErasmusPeriodRequest{}, "universitychainde/SendEndErasmusPeriodRequest", nil)
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
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
