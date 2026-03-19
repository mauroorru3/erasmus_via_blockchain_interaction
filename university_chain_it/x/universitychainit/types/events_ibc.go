package types

// IBC events
const (
	EventTypeTimeout                       = "timeout"
	EventTypeErasmusStudentPacket          = "erasmusStudent_packet"
	EventTypeErasmusIndexPacket            = "erasmusIndex_packet"
	EventTypeEndErasmusPeriodRequestPacket = "endErasmusPeriodRequest_packet"
	EventTypeFinalErasmusDataPacket        = "finalErasmusData_packet"
	EventTypeExtendErasmusPeriodPacket     = "extendErasmusPeriod_packet"
	EventTypeErasmusRestictedDataPacket    = "erasmusRestictedData_packet"
	// this line is used by starport scaffolding # ibc/packet/event

	AttributeKeyAckSuccess = "success"
	AttributeKeyAck        = "acknowledgement"
	AttributeKeyAckError   = "error"
)
