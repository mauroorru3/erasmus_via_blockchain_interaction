package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ErasmusExamsKeyPrefix is the prefix to retrieve all ErasmusExams
	ErasmusExamsKeyPrefix = "ErasmusExams/value/"
)

// ErasmusExamsKey returns the store key to retrieve a ErasmusExams from the index fields
func ErasmusExamsKey(
	examName string,
) []byte {
	var key []byte

	examNameBytes := []byte(examName)
	key = append(key, examNameBytes...)
	key = append(key, []byte("/")...)

	return key
}
