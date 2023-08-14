package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ForeignUniversitiesKeyPrefix is the prefix to retrieve all ForeignUniversities
	ForeignUniversitiesKeyPrefix = "ForeignUniversities/value/"
)

// ForeignUniversitiesKey returns the store key to retrieve a ForeignUniversities from the index fields
func ForeignUniversitiesKey(
	universityName string,
) []byte {
	var key []byte

	universityNameBytes := []byte(universityName)
	key = append(key, universityNameBytes...)
	key = append(key, []byte("/")...)

	return key
}
