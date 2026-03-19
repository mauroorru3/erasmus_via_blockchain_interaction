package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ExamsInfoKeyPrefix is the prefix to retrieve all ExamsInfo
	ExamsInfoKeyPrefix = "ExamsInfo/value/"
)

// ExamsInfoKey returns the store key to retrieve a ExamsInfo from the index fields
func ExamsInfoKey(
	examName string,
) []byte {
	var key []byte

	examNameBytes := []byte(examName)
	key = append(key, examNameBytes...)
	key = append(key, []byte("/")...)

	return key
}
