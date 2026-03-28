package keeper

import (
	"context"
	"time"
	"university_chain_de/x/universitychainde/types"
	"university_chain_de/x/universitychainde/utilfunc"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

// Function that checks whether the operation timer has expired and, if so,
// resends the packets of the various operations considered

func (k Keeper) TerminateExpiredOperations(goCtx context.Context) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	uniList := k.GetAllUniversityInfo(ctx)
	lenUniList := len(uniList)

	for i := 0; i < lenUniList; i++ {

		studentIndex := uniList[i].FifoHeadOperation
		finish := false
		count := 0
		for !finish {
			// Finished moving along
			if studentIndex == "" {
				finish = true
			} else {

				storedStudent, found := k.GetStoredStudent(ctx, studentIndex)

				if !found {
					panic("Fifo head game not found " + uniList[i].FifoHeadOperation)
				}
				deadline, err := GetOperationDeadline(storedStudent)
				if err != nil {
					panic(err)
				}

				s := utilfunc.FormatDeadline(ctx.BlockTime())
				formattedStartDate, _ := time.Parse(utilfunc.DeadlineLayout, s)

				if deadline.Before(formattedStartDate) {

					k.RemoveOperationFromFifo(ctx, &storedStudent, &uniList[i])

					k.SetStoredStudent(ctx, storedStudent)

					utilfunc.PrintLogs("TerminateExpiredOperations dentro deadline "+deadline.String()+" orario attuale "+formattedStartDate.String(), ctx)
					utilfunc.PrintLogs("TerminateExpiredOperations dentro deadline storedStudent.OperationInfo.StudentOperationDetails "+storedStudent.OperationInfo.StudentOperationDetails, ctx)

					opID, err := GetOperationID(storedStudent)
					if err != nil {
						panic(err)
					}

					switch opID {

					case "1":

						utilfunc.PrintLogs("TerminateExpiredOperations caso 1 - start erasmus", ctx)

						if storedStudent.Counters.RetryNumberOperations < storedStudent.Counters.MaximumNumberRetries {

							data, err := utilfunc.CreateHomeIndexJSONPacketFromStudentData(storedStudent)
							if err != nil {
								panic(err)
							}

							var packet types.ErasmusRestictedDataPacketData
							packet.ErasmusRestrictedInfo = data

							err = k.TransmitErasmusRestictedDataPacket(
								ctx,
								packet,
								"universitychainde",
								"channel-0",
								clienttypes.ZeroHeight(),
								timeoutTimestamp,
								" StartErasmus",
							)
							if err != nil {
								panic(err)
							}

							storedStudent.Counters.RetryNumberOperations++
							k.SetStoredStudent(ctx, storedStudent)

							// The timer for the start erasmus operation is created

							err = k.AddOperationQueue(ctx, &storedStudent, &uniList[i], "1", int(storedStudent.Counters.RetryNumberOperations))
							if err != nil {
								panic(err)
							}
							utilfunc.PrintLogs("TerminateExpiredOperations caso 1 - start erasmus - packet sent", ctx)
						} else {

							// revert state of the start erasmus operation

							err = k.ClearErasmusCareer(ctx, storedStudent.Index)
							if err != nil {
								panic(err)
							}

							storedStudent.Counters.RetryNumberOperations = 0
							k.SetStoredStudent(ctx, storedStudent)
						}

					case "2":

						utilfunc.PrintLogs("TerminateExpiredOperations caso 2 - end erasmus", ctx)

						var packet types.EndErasmusPeriodRequestPacketData

						packet.StartingUniversityName = storedStudent.StudentData.UniversityName
						packet.Index = storedStudent.Index

						foreignUniversityName, err := utilfunc.GetForeignUniversityName(storedStudent)
						if err != nil {
							panic(err)
						} else {

							packet.DestinationUniversityName = foreignUniversityName
							foreignIndex, err := utilfunc.GetForeignIndex(storedStudent)

							if err != nil {
								panic(err)
							} else {

								if storedStudent.Counters.RetryNumberOperations < storedStudent.Counters.MaximumNumberRetries {

									packet.ForeignIndex = foreignIndex

									err = k.TransmitEndErasmusPeriodRequestPacket(
										ctx,
										packet,
										"universitychainde",
										"channel-0",
										clienttypes.ZeroHeight(),
										timeoutTimestamp,
										"EndErasmusBeforeDeadline",
									)
									if err != nil {
										panic(err)
									}

									storedStudent.Counters.RetryNumberOperations++
									k.SetStoredStudent(ctx, storedStudent)

									// The timer for the end erasmus operation is created

									err = k.AddOperationQueue(ctx, &storedStudent, &uniList[i], "2", int(storedStudent.Counters.RetryNumberOperations))
									if err != nil {
										panic(err)
									}
									utilfunc.PrintLogs("TerminateExpiredOperations caso 2 - end erasmus - packet sent", ctx)
								} else {

									// revert state of the end erasmus operation

									err = k.RevertEndErasmus(ctx, storedStudent.Index)
									if err != nil {
										panic(err)
									}

									storedStudent.Counters.RetryNumberOperations = 0
									k.SetStoredStudent(ctx, storedStudent)
								}
							}
						}

					case "3":

						utilfunc.PrintLogs("TerminateExpiredOperations caso 3 - extend Erasmus", ctx)

						var packet types.ExtendErasmusPeriodPacketData

						foreignUni, err := utilfunc.GetForeignUniversityName(storedStudent)
						if err != nil {
							panic(err)
						} else {

							foreignIndex, err := utilfunc.GetForeignIndex(storedStudent)
							if err != nil {
								panic(err)
							} else {

								if storedStudent.Counters.RetryNumberOperations < storedStudent.Counters.MaximumNumberRetries {

									packet.DestinationUniversityName = foreignUni
									packet.ForeignIndex = foreignIndex
									packet.DurationInMonths = 6
									packet.FinalDate, _ = utilfunc.GetFinalDateErasmus(storedStudent)

									// Transmit the packet
									err = k.TransmitExtendErasmusPeriodPacket(
										ctx,
										packet,
										"universitychainde",
										"channel-0",
										clienttypes.ZeroHeight(),
										timeoutTimestamp,
										"ExtendErasmus",
									)
									if err != nil {
										panic(err)
									}

									storedStudent.Counters.RetryNumberOperations++
									k.SetStoredStudent(ctx, storedStudent)

									// The timer for the extend erasmus operation is created

									err = k.AddOperationQueue(ctx, &storedStudent, &uniList[i], "3", int(storedStudent.Counters.RetryNumberOperations))
									if err != nil {
										panic(err)
									}
									utilfunc.PrintLogs("TerminateExpiredOperations caso 3 - extend Erasmus - packet sent", ctx)
								} else {

									// revert state of the extend erasmus operation

									err := k.RevertExtendErasmus(ctx, storedStudent.Index)
									if err != nil {
										panic(err)
									}

									storedStudent.Counters.RetryNumberOperations = 0
									k.SetStoredStudent(ctx, storedStudent)
								}
							}
						}

					case "4":

						utilfunc.PrintLogs("TerminateExpiredOperations caso 4 - other 10 packets start Erasmus", ctx)

						if storedStudent.Counters.RetryNumberOperations < storedStudent.Counters.MaximumNumberRetries {

							data, err := utilfunc.CreateNameSurnameJSONPacketFromStudentData(storedStudent)
							if err != nil {
								panic(err)
							}

							var packet types.ErasmusRestictedDataPacketData
							packet.ErasmusRestrictedInfo = data

							err = k.TransmitErasmusRestictedDataPacket(
								ctx,
								packet,
								"universitychainde",
								"channel-0",
								clienttypes.ZeroHeight(),
								timeoutTimestamp,
								" TerminateExpiredOperations caso 4",
							)

							if err != nil {
								utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
								panic(err)
							} else {

								utilfunc.PrintLogs("TerminateExpiredOperations caso 4 CreateNameSurnameJSONPacketFromStudentData sent", ctx)
								data, err := utilfunc.CreateStudentKeyPart1JSONPacketFromStudentData(storedStudent)
								if err != nil {
									panic(err)
								}

								var packet types.ErasmusRestictedDataPacketData
								packet.ErasmusRestrictedInfo = data

								err = k.TransmitErasmusRestictedDataPacket(
									ctx,
									packet,
									"universitychainde",
									"channel-0",
									clienttypes.ZeroHeight(),
									timeoutTimestamp,
									" TerminateExpiredOperations caso 4",
								)

								if err != nil {
									utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
									panic(err)
								} else {

									utilfunc.PrintLogs("TerminateExpiredOperations caso 4 CreateStudentKeyPart1JSONPacketFromStudentData sent", ctx)

									data, err := utilfunc.CreateStudentKeyPart2JSONPacketFromStudentData(storedStudent)
									if err != nil {
										panic(err)
									}

									var packet types.ErasmusRestictedDataPacketData
									packet.ErasmusRestrictedInfo = data

									err = k.TransmitErasmusRestictedDataPacket(
										ctx,
										packet,
										"universitychainde",
										"channel-0",
										clienttypes.ZeroHeight(),
										timeoutTimestamp,
										" TerminateExpiredOperations caso 4",
									)

									if err != nil {
										utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
										panic(err)
									} else {
										utilfunc.PrintLogs("TerminateExpiredOperations caso 4 CreateStudentKeyPart2JSONPacketFromStudentData sent", ctx)

										data, err := utilfunc.CreateStartDateJSONPacketFromStudentData(storedStudent)
										if err != nil {
											panic(err)
										}

										var packet types.ErasmusRestictedDataPacketData
										packet.ErasmusRestrictedInfo = data

										err = k.TransmitErasmusRestictedDataPacket(
											ctx,
											packet,
											"universitychainde",
											"channel-0",
											clienttypes.ZeroHeight(),
											timeoutTimestamp,
											" TerminateExpiredOperations caso 4",
										)

										if err != nil {
											utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
											panic(err)
										} else {
											utilfunc.PrintLogs("TerminateExpiredOperations caso 4 CreateStartDateJSONPacketFromStudentData sent", ctx)
											data, err := utilfunc.CreateEndDateJSONPacketFromStudentData(storedStudent)
											if err != nil {
												panic(err)
											}

											var packet types.ErasmusRestictedDataPacketData
											packet.ErasmusRestrictedInfo = data

											err = k.TransmitErasmusRestictedDataPacket(
												ctx,
												packet,
												"universitychainde",
												"channel-0",
												clienttypes.ZeroHeight(),
												timeoutTimestamp,
												" TerminateExpiredOperations caso 4",
											)

											if err != nil {
												utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
												panic(err)
											} else {
												utilfunc.PrintLogs("TerminateExpiredOperations caso 4 CreateEndDateJSONPacketFromStudentData sent", ctx)
												data, err := utilfunc.CreateDurationJSONPacketFromStudentData(storedStudent)
												if err != nil {
													panic(err)
												}

												var packet types.ErasmusRestictedDataPacketData
												packet.ErasmusRestrictedInfo = data

												err = k.TransmitErasmusRestictedDataPacket(
													ctx,
													packet,
													"universitychainde",
													"channel-0",
													clienttypes.ZeroHeight(),
													timeoutTimestamp,
													" TerminateExpiredOperations caso 4",
												)

												if err != nil {
													utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
													panic(err)
												} else {
													utilfunc.PrintLogs("TerminateExpiredOperations caso 4 CreateDurationJSONPacketFromStudentData sent", ctx)
													data, err := utilfunc.CreateCourseDetailsJSONPacketFromStudentData(storedStudent)
													if err != nil {
														utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
														panic(err)
													}

													var packet types.ErasmusRestictedDataPacketData
													packet.ErasmusRestrictedInfo = data

													err = k.TransmitErasmusRestictedDataPacket(
														ctx,
														packet,
														"universitychainde",
														"channel-0",
														clienttypes.ZeroHeight(),
														timeoutTimestamp,
														" TerminateExpiredOperations caso 4",
													)

													if err != nil {
														utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
														panic(err)
													} else {
														utilfunc.PrintLogs("TerminateExpiredOperations caso 4 CreateCourseDetailsJSONPacketFromStudentData sent", ctx)
														data, err := utilfunc.CreateDepartmentJSONPacketFromStudentData(storedStudent)
														if err != nil {
															panic(err)
														}

														var packet types.ErasmusRestictedDataPacketData
														packet.ErasmusRestrictedInfo = data

														err = k.TransmitErasmusRestictedDataPacket(
															ctx,
															packet,
															"universitychainde",
															"channel-0",
															clienttypes.ZeroHeight(),
															timeoutTimestamp,
															" TerminateExpiredOperations caso 4 - case 9",
														)

														if err != nil {
															utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
															panic(err)
														} else {
															utilfunc.PrintLogs("TerminateExpiredOperations caso 4 CreateDepartmentJSONPacketFromStudentData sent", ctx)
															data, err := utilfunc.CreateErasmusTypeJSONPacketFromStudentData(storedStudent)
															if err != nil {
																utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
																panic(err)
															}

															var packet types.ErasmusRestictedDataPacketData
															packet.ErasmusRestrictedInfo = data

															err = k.TransmitErasmusRestictedDataPacket(
																ctx,
																packet,
																"universitychainde",
																"channel-0",
																clienttypes.ZeroHeight(),
																timeoutTimestamp,
																" TerminateExpiredOperations caso 4 - case 10",
															)

															if err != nil {
																utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
																panic(err)
															} else {
																utilfunc.PrintLogs("TerminateExpiredOperations caso 4 CreateErasmusTypeJSONPacketFromStudentData sent", ctx)
																data, err := utilfunc.CreateExamsJSONPacketFromStudentData(storedStudent)
																if err != nil {
																	utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
																	panic(err)
																}

																var packet types.ErasmusRestictedDataPacketData
																packet.ErasmusRestrictedInfo = data

																err = k.TransmitErasmusRestictedDataPacket(
																	ctx,
																	packet,
																	"universitychainde",
																	"channel-0",
																	clienttypes.ZeroHeight(),
																	timeoutTimestamp,
																	" TerminateExpiredOperations caso 4",
																)

																if err != nil {
																	utilfunc.PrintLogs("TerminateExpiredOperations caso 4 "+err.Error(), ctx)
																	panic(err)
																} else {

																	// The timer for the 10 packets of the start erasmus operation is created

																	err = k.AddOperationQueue(ctx, &storedStudent, &uniList[i], "4", int(storedStudent.Counters.RetryNumberOperations))
																	if err != nil {
																		panic(err)
																	}
																	utilfunc.PrintLogs("TerminateExpiredOperations caso 4 - start Erasmus - 10 packets sent", ctx)

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
						} else {

							// revert state of the start erasmus operation

							err = k.ClearErasmusCareer(ctx, storedStudent.Index)
							if err != nil {
								panic(err)
							}
							k.SetStoredStudent(ctx, storedStudent)
						}
					}

					// Move along FIFO
					studentIndex = uniList[i].FifoHeadOperation

					count++

				} else {

					finish = true
				}

			}
		}
		if count > 0 {
			k.SetUniversityInfo(ctx, uniList[i])
		}

	}

}
