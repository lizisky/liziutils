package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	mathrand "math/rand"
	"time"
)

func GenerateRandomStr(inseed []byte) string {

	lenBuf := GenerateRandomNumInRange(32, 64)
	buf := make([]byte, lenBuf)
	io.ReadFull(rand.Reader, buf)

	curTime := time.Now().UnixNano()
	tmpseed := []byte(fmt.Sprintf("%s%d%s", buf, curTime, inseed))

	s256 := Sha256(tmpseed)
	// fmt.Println("s256", len(s256), s256)
	return s256

	// tmpHash := sha256.Sum256(tmpseed)
	// b64 := Base64Encode(tmpseed)
	// fmt.Println("b64", len(b64), b64)

	// return hex.EncodeToString(tmpHash[:])
}

// func GenerateCryptRandomStr() string {
// 	buf := make([]byte, 32)
// 	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
// 		return ""
// 	}
// 	return base64.URLEncoding.EncodeToString(b)
// }

func GenerateRandomNumInRange(min, max int) uint32 {
	if min < 1 {
		min = 1
	}
	if max <= min {
		max = min + 16
	}

	for {
		seed2 := time.Now().UnixNano()
		mathrand.New(mathrand.NewSource(seed2))
		random_num := mathrand.Uint32() % uint32(max)
		if random_num >= uint32(min) {
			return (random_num)
		}
	}
}
