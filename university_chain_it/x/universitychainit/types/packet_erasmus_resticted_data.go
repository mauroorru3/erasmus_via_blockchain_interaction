package types

// ValidateBasic is used for validating the packet
func (p ErasmusRestictedDataPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p ErasmusRestictedDataPacketData) GetBytes() ([]byte, error) {
	var modulePacket UniversitychainitPacketData

	modulePacket.Packet = &UniversitychainitPacketData_ErasmusRestictedDataPacket{&p}

	return modulePacket.Marshal()
}