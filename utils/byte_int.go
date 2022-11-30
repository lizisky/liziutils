package utils

import "encoding/binary"

const (
	length_4 = 4
	length_8 = 8
)

func Uint32ToBytes(dat uint32) []byte {
	// var buf = make([]byte, 4)
	buf := [length_4]byte{}
	binary.BigEndian.PutUint32(buf[:], dat)
	return buf[:]
}

func BytesToUint32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

func Uint64ToBytes(dat uint64) []byte {
	// var buf = make([]byte, 8)
	buf := [length_8]byte{}
	binary.BigEndian.PutUint64(buf[:], dat)
	return buf[:]
}

func BytesToUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}
