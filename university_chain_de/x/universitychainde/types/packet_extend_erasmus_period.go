package types

// ValidateBasic is used for validating the packet
func (p ExtendErasmusPeriodPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p ExtendErasmusPeriodPacketData) GetBytes() ([]byte, error) {
	var modulePacket UniversitychaindePacketData

	modulePacket.Packet = &UniversitychaindePacketData_ExtendErasmusPeriodPacket{&p}

	return modulePacket.Marshal()
}
