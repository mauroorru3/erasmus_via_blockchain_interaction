package types

// ValidateBasic is used for validating the packet
func (p FinalErasmusDataPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p FinalErasmusDataPacketData) GetBytes() ([]byte, error) {
	var modulePacket HubPacketData

	modulePacket.Packet = &HubPacketData_FinalErasmusDataPacket{&p}

	return modulePacket.Marshal()
}
