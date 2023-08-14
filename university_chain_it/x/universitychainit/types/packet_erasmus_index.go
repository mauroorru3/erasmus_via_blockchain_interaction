package types

// ValidateBasic is used for validating the packet
func (p ErasmusIndexPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p ErasmusIndexPacketData) GetBytes() ([]byte, error) {
	var modulePacket UniversitychainitPacketData

	modulePacket.Packet = &UniversitychainitPacketData_ErasmusIndexPacket{&p}

	return modulePacket.Marshal()
}
