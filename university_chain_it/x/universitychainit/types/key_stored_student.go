package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StoredStudentKeyPrefix is the prefix to retrieve all StoredStudent
	StoredStudentKeyPrefix = "StoredStudent/value/"
)

// StoredStudentKey returns the store key to retrieve a StoredStudent from the index fields
func StoredStudentKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
