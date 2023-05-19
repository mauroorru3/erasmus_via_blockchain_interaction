package types

// ValidateBasic is used for validating the packet
func (p EndErasmusPeriodRequestPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p EndErasmusPeriodRequestPacketData) GetBytes() ([]byte, error) {
	var modulePacket UniversitychainitPacketData

	modulePacket.Packet = &UniversitychainitPacketData_EndErasmusPeriodRequestPacket{&p}

	return modulePacket.Marshal()
}
