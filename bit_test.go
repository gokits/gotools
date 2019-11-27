package gotools

import "testing"

func TestGetBitR2L(t *testing.T) {
	idx := []uint{2, 3, 5, 8, 9, 11}
	value := []uint64{0, 1, 0, 1, 1, 0}
	for i := range idx {
		actual := GetBitR2L(5896, idx[i])
		if actual != value[i] {
			t.Fatalf("GetBit32(5896, %d) wrong, expected %d, actual %d", idx[i], value[i], actual)
		}
	}
}

func TestGetBit32L2R(t *testing.T) {
	idx := []uint{2, 3, 5, 8, 9, 10}
	value := []uint32{1, 0, 0, 0, 0, 1}
	for i := range idx {
		actual := GetBit32L2R(589685412, idx[i])
		if actual != value[i] {
			t.Fatalf("GetBit32(5896, %d) wrong, expected %d, actual %d", idx[i], value[i], actual)
		}
	}
}
