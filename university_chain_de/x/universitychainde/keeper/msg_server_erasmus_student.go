package keeper

import (
	"context"

	"university_chain_de/x/universitychainde/types"
	"university_chain_de/x/universitychainde/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

func (k msgServer) SendErasmusStudent(goCtx context.Context, msg *types.MsgSendErasmusStudent) (*types.MsgSendErasmusStudentResponse, error) {

	utilfunc.PrintLogs("SendErasmusStudent")

	ctx := sdk.UnwrapSDKContext(goCtx)

	var packet types.ErasmusStudentPacketData

	packet.Student = nil

	// Transmit the packet
	err := k.TransmitErasmusStudentPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		timeoutTimestamp,
	)
	if err != nil {
		return nil, err
	} else {
		return &types.MsgSendErasmusStudentResponse{
			Status: 0,
		}, nil
	}

	/*

		utilfunc.PrintLogs("SendErasmusStudent")

		return &types.MsgSendErasmusStudentResponse{
			Status: -1,
		}, types.ErrNonCallableFunction

	*/

	/*

		ctx := sdk.UnwrapSDKContext(goCtx)

		chainInfo, found := k.Keeper.GetChainInfo(ctx)
		if !found {
			panic("ChainInfo not found")
		} else {
			if !chainInfo.InitStatus {
				return &types.MsgSendErasmusStudentResponse{
					Status: -1,
				}, types.ErrChainConfigurationNotDone
			} else {

				_, found := k.Keeper.GetUniversityInfo(ctx, strings.Split(msg.Index, "_")[0])
				if !found {
					return &types.MsgSendErasmusStudentResponse{
						Status: -1,
					}, types.ErrWrongNameUniversity
				} else {

					searchedStudent, found := k.Keeper.GetStoredStudent(ctx, msg.Index)

					if !found {
						return &types.MsgSendErasmusStudentResponse{
							Status: -1,
						}, types.ErrStudentNotPresent
					} else {
						if searchedStudent.GetStudentData().GetStudentKey() != msg.Creator {
							return &types.MsgSendErasmusStudentResponse{
								Status: -1,
							}, types.ErrKeyEnteredMismatchStudent
						} else {

							err := utilfunc.CheckCompleteInformation(searchedStudent)

							if err != nil {
								return &types.MsgSendErasmusStudentResponse{
									Status: -1,
								}, types.ErrIncompleteStudentInformation
							} else {

								ok, err := utilfunc.CheckTaxPayment(searchedStudent)

								if err != nil {
									return &types.MsgSendErasmusStudentResponse{
										Status: -1,
									}, err
								} else {
									if !ok {
										return &types.MsgSendErasmusStudentResponse{
											Status: -1,
										}, types.ErrUnpaidTaxes
									} else {

										res, err := utilfunc.CheckErasmusStatus(searchedStudent)
										if err != nil {
											return &types.MsgSendErasmusStudentResponse{
												Status: -1,
											}, err
										} else {
											if res == "" {
												return &types.MsgSendErasmusStudentResponse{
													Status: -1,
												}, types.ErrNoErasmusRequest
											} else if res == "in progress" {
												return &types.MsgSendErasmusStudentResponse{
													Status: -1,
												}, types.ErrPreviousRequestInProgress
											} else if res == "terminated" {
												return &types.MsgSendErasmusStudentResponse{
													Status: -1,
												}, types.ErrPreviousRequestCompleted
											} else {

												// TODO: logic before transmitting the packet

												utilfunc.PrintLogs("SendErasmusStudent")

												var packet types.ErasmusStudentPacketData

												packet.Student = &searchedStudent

												// Transmit the packet
												err = k.TransmitErasmusStudentPacket(
													ctx,
													packet,
													msg.Port,
													msg.ChannelID,
													clienttypes.ZeroHeight(),
													msg.TimeoutTimestamp,
												)
												if err != nil {
													return nil, err
												} else {
													return &types.MsgSendErasmusStudentResponse{
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
	*/
}
