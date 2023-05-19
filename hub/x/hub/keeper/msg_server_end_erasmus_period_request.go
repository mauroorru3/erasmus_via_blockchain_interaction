package keeper

import (
	"context"
	"strings"

	"hub/x/hub/types"
	"hub/x/hub/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) SendEndErasmusPeriodRequest(goCtx context.Context, msg *types.MsgSendEndErasmusPeriodRequest) (*types.MsgSendEndErasmusPeriodRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			return &types.MsgSendEndErasmusPeriodRequestResponse{},
				types.ErrChainConfigurationNotDone
		} else {

			_, found := k.Keeper.GetUniversities(ctx, strings.Split(msg.Index, "_")[0])
			if !found {
				return &types.MsgSendEndErasmusPeriodRequestResponse{},
					types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.Index)

				if !found {
					return &types.MsgSendEndErasmusPeriodRequestResponse{},
						types.ErrStudentNotPresent
				} else {
					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgSendEndErasmusPeriodRequestResponse{},
							types.ErrKeyEnteredMismatchStudent
					} else {

						err := utilfunc.CheckCompleteInformation(searchedStudent)

						if err != nil {
							return &types.MsgSendEndErasmusPeriodRequestResponse{},
								types.ErrIncompleteStudentInformation
						} else {

							ok, err := utilfunc.CheckTaxPayment(searchedStudent)

							if err != nil {
								return &types.MsgSendEndErasmusPeriodRequestResponse{},
									err
							} else {
								if !ok {
									return &types.MsgSendEndErasmusPeriodRequestResponse{},
										types.ErrUnpaidTaxes
								} else {

									res, err := utilfunc.CheckErasmusStatus(searchedStudent)
									if err != nil {
										return &types.MsgSendEndErasmusPeriodRequestResponse{},
											err
									} else {
										if res == "" {
											return &types.MsgSendEndErasmusPeriodRequestResponse{},
												types.ErrNoErasmusRequest
										} else if res == "to start" {
											return &types.MsgSendEndErasmusPeriodRequestResponse{},
												types.ErrPreviousRequestInProgress
										} else if res == "terminated" {
											return &types.MsgSendEndErasmusPeriodRequestResponse{},
												types.ErrPreviousRequestCompleted
										} else {

											// TODO: logic before transmitting the packet

											utilfunc.PrintLogs("SendEndErasmusPeriodRequest")

											// Construct the packet
											var packet types.EndErasmusPeriodRequestPacketData

											packet.StartingUniversityName = msg.StartingUniversityName
											packet.DestinationUniversityName = msg.DestinationUniversityName
											packet.Index = msg.Index
											packet.ForeignIndex = msg.ForeignIndex

											// Transmit the packet
											err := k.TransmitEndErasmusPeriodRequestPacket(
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

											return &types.MsgSendEndErasmusPeriodRequestResponse{}, nil
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
