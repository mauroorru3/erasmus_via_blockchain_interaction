package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSendErasmusStudent{}, "hub/SendErasmusStudent", nil)
	cdc.RegisterConcrete(&MsgConfigureChain{}, "hub/ConfigureChain", nil)
	cdc.RegisterConcrete(&MsgSendErasmusIndex{}, "hub/SendErasmusIndex", nil)
	cdc.RegisterConcrete(&MsgSendEndErasmusPeriodRequest{}, "hub/SendEndErasmusPeriodRequest", nil)
	cdc.RegisterConcrete(&MsgSendFinalErasmusData{}, "hub/SendFinalErasmusData", nil)
	cdc.RegisterConcrete(&MsgSendExtendErasmusPeriod{}, "hub/SendExtendErasmusPeriod", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendErasmusStudent{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgConfigureChain{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendErasmusIndex{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendEndErasmusPeriodRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendFinalErasmusData{},
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
