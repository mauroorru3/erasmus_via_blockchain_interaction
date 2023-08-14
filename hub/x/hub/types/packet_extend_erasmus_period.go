package types

// ValidateBasic is used for validating the packet
func (p ExtendErasmusPeriodPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p ExtendErasmusPeriodPacketData) GetBytes() ([]byte, error) {
	var modulePacket HubPacketData

	modulePacket.Packet = &HubPacketData_ExtendErasmusPeriodPacket{&p}

	return modulePacket.Marshal()
}
