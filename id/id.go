package id

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
)

// ID is a peer id representation.
type ID []byte

// NewFromHex returns ID from hexadecimal string.
func NewFromHex(s string) (ID, error) {
	return hex.DecodeString(s)
}

// NewRandID returns ID from length.
func NewRandID(length int) (ID, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	return ID(b), err
}

// Hex return hexadecimal ID representation.
func (id ID) Hex() string {
	return hex.EncodeToString(id)
}

// Less checks if id less than id2
func (id ID) Less(id2 ID) bool {
	return bytes.Compare(id, id2) < 0
}

// Equals checks if id equal to id2
func (id ID) Equals(id2 ID) bool {
	return bytes.Equal(id, id2)
}

// XOR calculate the distance between any two keys in the same address space.
// distance(id, id2) = id ^ id2 where ^ represents the XOR operator.
// The result is obtained by taking the bytewise exclusive OR
// of each byte of the operands.
func XOR(id, id2 ID) ID {
	result := make([]byte, id.Len())
	for i := 0; i < id.Len(); i++ {
		result[i] = id[i] ^ id2[i]
	}
	return result
}

// LeadingZerosLen returns the number of consecutive zeroes in a byte slice.
func (id ID) LeadingZerosLen() int {
	for i := 0; i < len(id); i++ {
		for j := 0; j < 8; j++ {
			if (id[i]>>uint8(7-j))&0x1 != 0 {
				return i*8 + j
			}
		}
	}
	return len(id) * 8
}

// Len returns number of bytes.
func (id ID) Len() int {
	return len(id)
}
