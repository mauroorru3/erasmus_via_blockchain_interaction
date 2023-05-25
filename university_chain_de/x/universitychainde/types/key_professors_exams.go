package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ProfessorsExamsKeyPrefix is the prefix to retrieve all ProfessorsExams
	ProfessorsExamsKeyPrefix = "ProfessorsExams/value/"
)

// ProfessorsExamsKey returns the store key to retrieve a ProfessorsExams from the index fields
func ProfessorsExamsKey(
	examName string,
) []byte {
	var key []byte

	examNameBytes := []byte(examName)
	key = append(key, examNameBytes...)
	key = append(key, []byte("/")...)

	return key
}
