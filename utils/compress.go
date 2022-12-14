package utils

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Compress(data []byte) []byte {
	if len(data) < 1 {
		return nil
	}

	var byBuf bytes.Buffer

	w := gzip.NewWriter(&byBuf)
	w.Write(data)
	w.Close()

	result := byBuf.Bytes()
	// fmt.Println("compress: original:", len(data), "--result:", len(result))

	return result
}

func DeCompress(data []byte) []byte {
	if len(data) < 1 {
		return nil
	}

	inBuf := bytes.NewReader(data)

	r, err := gzip.NewReader(inBuf)
	if err != nil {
		return nil
	}
	defer r.Close()

	var outBuf bytes.Buffer
	io.Copy(&outBuf, r)

	// fmt.Println("de-compress: original:", len(data), "---result:", len(outBuf.Bytes()))

	return outBuf.Bytes()
}

// func Compress(data []byte) []byte {
// 	if len(data) < 1 {
// 		return nil
// 	}

// 	var byBuf bytes.Buffer

// 	w := zlib.NewWriter(&byBuf)
// 	w.Write(data)
// 	w.Close()

// 	result := byBuf.Bytes()
// 	// fmt.Println("compress: original:", len(data), "--result:", len(result))

// 	return result
// }

// func DeCompress(data []byte) []byte {
// 	if len(data) < 1 {
// 		return nil
// 	}

// 	inBuf := bytes.NewReader(data)

// 	r, err := zlib.NewReader(inBuf)
// 	if err != nil {
// 		return nil
// 	}
// 	defer r.Close()

// 	var outBuf bytes.Buffer
// 	io.Copy(&outBuf, r)

// 	// fmt.Println("de-compress: original:", len(data), "---result:", len(outBuf.Bytes()))

// 	return outBuf.Bytes()
// }
