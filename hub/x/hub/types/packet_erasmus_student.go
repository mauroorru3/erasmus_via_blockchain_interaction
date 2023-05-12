package types

// ValidateBasic is used for validating the packet
func (p ErasmusStudentPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p ErasmusStudentPacketData) GetBytes() ([]byte, error) {
	var modulePacket HubPacketData

	modulePacket.Packet = &HubPacketData_ErasmusStudentPacket{&p}

	return modulePacket.Marshal()
}
