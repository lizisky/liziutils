package utils

import "encoding/binary"

const (
	Length_4 = 4
	Length_8 = 8
)

func Uint32ToBytesBuffer(buffer []byte, dat uint32) {
	binary.BigEndian.PutUint32(buffer[:Length_4], dat)
}

func Uint32ToBytes(dat uint32) []byte {
	// var buf = make([]byte, 4)
	buf := [Length_4]byte{}
	binary.BigEndian.PutUint32(buf[:], dat)
	return buf[:]
}

func BytesToUint32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

func Uint64ToBytes(dat uint64) []byte {
	// var buf = make([]byte, 8)
	buf := [Length_8]byte{}
	binary.BigEndian.PutUint64(buf[:], dat)
	return buf[:]
}

func BytesToUint64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}
