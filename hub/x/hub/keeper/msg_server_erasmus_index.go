package keeper

import (
	"context"
	"strings"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) SendErasmusIndex(goCtx context.Context, msg *types.MsgSendErasmusIndex) (*types.MsgSendErasmusIndexResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			return &types.MsgSendErasmusIndexResponse{},
				types.ErrChainConfigurationNotDone
		} else {

			_, found := k.Keeper.GetUniversities(ctx, strings.Split(msg.Index, "_")[0])
			if !found {
				return &types.MsgSendErasmusIndexResponse{},
					types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.Index)

				if !found {
					return &types.MsgSendErasmusIndexResponse{},
						types.ErrStudentNotPresent
				} else {
					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgSendErasmusIndexResponse{},
							types.ErrKeyEnteredMismatchStudent
					} else {

						err := utilfunc.CheckCompleteInformation(searchedStudent)

						if err != nil {
							return &types.MsgSendErasmusIndexResponse{},
								types.ErrIncompleteStudentInformation
						} else {

							ok, err := utilfunc.CheckTaxPayment(searchedStudent)

							if err != nil {
								return &types.MsgSendErasmusIndexResponse{},
									err
							} else {
								if !ok {
									return &types.MsgSendErasmusIndexResponse{},
										types.ErrUnpaidTaxes
								} else {

									res, err := utilfunc.CheckErasmusStatus(searchedStudent)
									if err != nil {
										return &types.MsgSendErasmusIndexResponse{},
											err
									} else {
										if res == "" {
											return &types.MsgSendErasmusIndexResponse{},
												types.ErrNoErasmusRequest
										} else if res == "to start" {
											return &types.MsgSendErasmusIndexResponse{},
												types.ErrPreviousRequestInProgress
										} else if res == "terminated" {
											return &types.MsgSendErasmusIndexResponse{},
												types.ErrPreviousRequestCompleted
										} else {

											// TODO: logic before transmitting the packet

											utilfunc.PrintLogs("SendErasmusIndex")

											var packet types.ErasmusIndexPacketData

											packet.Index = msg.Index
											packet.ForeignIndex = msg.ForeignIndex

											// Transmit the packet
											err := k.TransmitErasmusIndexPacket(
												ctx,
												packet,
												msg.Port,
												msg.ChannelID,
												clienttypes.ZeroHeight(),
												msg.TimeoutTimestamp,
											)
											if err != nil {
												return nil, err
											}

											return &types.MsgSendErasmusIndexResponse{}, nil
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
