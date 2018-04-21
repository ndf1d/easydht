package id

import (
	"testing"
)

// no lint
func TestID_Hex(t *testing.T) {
	id, _ := NewFromHex("001e")
	if id.Hex() != "001e" {
		t.Fail()
	}
}

// no lint
func TestID_Less(t *testing.T) {
	id, _ := NewFromHex("00000001")
	id2, _ := NewFromHex("00000003")
	if id2.Less(id) {
		t.Fail()
	}
}

// no lint
func TestID_Equals(t *testing.T) {
	id, _ := NewFromHex("749a74bb2b72")
	id2, _ := NewFromHex("749a74bb2b72")
	if !id2.Equals(id) {
		t.Fail()
	}
}

// no lint
func TestID_Len(t *testing.T) {
	id, _ := NewFromHex("72aa371a5cd718175e81")
	if id.Len() != 10 {
		t.Fail()
	}
}

// no lint
func TestXOR(t *testing.T) {
	id, _ := NewFromHex("ab")  // 10101011
	id2, _ := NewFromHex("ca") // 11001010
	res := XOR(id, id2)
	if res.Hex() != "61" { // 01100001
		t.Fail()
	}
}

// no lint
func TestID_LeadingZerosLen(t *testing.T) {
	id, _ := NewFromHex("86f7e437faa5a7fce15d1ddcb9eaeaea377667b8")
	id2, _ := NewFromHex("86f516841ba77a5b4648de2cd0dfcb30ea46dbb4")
	if XOR(id, id2).LeadingZerosLen() != 14 {
		t.Fail()
	}
}
