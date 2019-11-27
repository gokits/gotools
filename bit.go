package gotools

func GetBitR2L(d uint64, idx uint) uint64 {
	return d >> idx & 1
}

func GetBit32L2R(d uint32, idx uint) uint32 {
	return d >> (32 - idx - 1) & 1
}
