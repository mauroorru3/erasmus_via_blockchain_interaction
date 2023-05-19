package keeper

import (
	"context"
	"strings"

	"university_chain_it/x/universitychainit/types"
	"university_chain_it/x/universitychainit/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) SendEndErasmusPeriodRequest(goCtx context.Context, msg *types.MsgSendEndErasmusPeriodRequest) (*types.MsgSendEndErasmusPeriodRequestResponse, error) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	chainInfo, found := k.Keeper.GetChainInfo(ctx)
	if !found {
		panic("ChainInfo not found")
	} else {
		if !chainInfo.InitStatus {
			return &types.MsgSendEndErasmusPeriodRequestResponse{
				Status: -1,
			}, types.ErrChainConfigurationNotDone
		} else {

			_, found := k.Keeper.GetUniversityInfo(ctx, strings.Split(msg.Index, "_")[0])
			if !found {
				return &types.MsgSendEndErasmusPeriodRequestResponse{
					Status: -1,
				}, types.ErrWrongNameUniversity
			} else {

				searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.Index)

				if !found {
					return &types.MsgSendEndErasmusPeriodRequestResponse{
						Status: -1,
					}, types.ErrStudentNotPresent
				} else {
					if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
						return &types.MsgSendEndErasmusPeriodRequestResponse{
							Status: -1,
						}, types.ErrKeyEnteredMismatchStudent
					} else {

						err := utilfunc.CheckCompleteInformation(searchedStudent)

						if err != nil {
							return &types.MsgSendEndErasmusPeriodRequestResponse{
								Status: -1,
							}, types.ErrIncompleteStudentInformation
						} else {

							ok, err := utilfunc.CheckTaxPayment(searchedStudent)

							if err != nil {
								return &types.MsgSendEndErasmusPeriodRequestResponse{
									Status: -1,
								}, err
							} else {
								if !ok {
									return &types.MsgSendEndErasmusPeriodRequestResponse{
										Status: -1,
									}, types.ErrUnpaidTaxes
								} else {

									res, err := utilfunc.CheckErasmusStatus(searchedStudent)
									if err != nil {
										return &types.MsgSendEndErasmusPeriodRequestResponse{
											Status: -1,
										}, err
									} else {
										if res == "" {
											return &types.MsgSendEndErasmusPeriodRequestResponse{
												Status: -1,
											}, types.ErrNoErasmusRequest
										} else if res == "in progress" {
											return &types.MsgSendEndErasmusPeriodRequestResponse{
												Status: -1,
											}, types.ErrPreviousRequestInProgress
										} else if res == "to start" {
											return &types.MsgSendEndErasmusPeriodRequestResponse{
												Status: -1,
											}, types.ErrPreviousRequestStartup
										} else {

											// TODO: logic before transmitting the packet

											utilfunc.PrintLogs("SendEndErasmusPeriodRequest")

											var packet types.EndErasmusPeriodRequestPacketData

											packet.StartingUniversityName = msg.StartingUniversityName
											packet.DestinationUniversityName = msg.DestinationUniversityName
											packet.Index = msg.Index
											packet.ForeignIndex = msg.ForeignIndex

											// Transmit the packet
											err = k.TransmitEndErasmusPeriodRequestPacket(
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

											return &types.MsgSendEndErasmusPeriodRequestResponse{
												Status: 0,
											}, nil
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
