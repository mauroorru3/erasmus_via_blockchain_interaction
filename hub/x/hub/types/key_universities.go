package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UniversitiesKeyPrefix is the prefix to retrieve all Universities
	UniversitiesKeyPrefix = "Universities/value/"
)

// UniversitiesKey returns the store key to retrieve a Universities from the index fields
func UniversitiesKey(
	universityName string,
) []byte {
	var key []byte

	universityNameBytes := []byte(universityName)
	key = append(key, universityNameBytes...)
	key = append(key, []byte("/")...)

	return key
}
