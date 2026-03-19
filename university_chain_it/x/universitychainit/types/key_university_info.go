package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// UniversityInfoKeyPrefix is the prefix to retrieve all UniversityInfo
	UniversityInfoKeyPrefix = "UniversityInfo/value/"
)

// UniversityInfoKey returns the store key to retrieve a UniversityInfo from the index fields
func UniversityInfoKey(
	universityName string,
) []byte {
	var key []byte

	universityNameBytes := []byte(universityName)
	key = append(key, universityNameBytes...)
	key = append(key, []byte("/")...)

	return key
}
