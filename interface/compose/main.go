package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
)

type HashReader interface {
	io.Reader
	hash() string
}

type hashReader struct {
	*bytes.Reader
	buf *bytes.Buffer
}

func NewHashReader(b []byte) *hashReader {
	return &hashReader{
		Reader: bytes.NewReader(b),
		buf:    bytes.NewBuffer(b),
	}
}

func (h *hashReader) hash() string {
	return hex.EncodeToString(h.buf.Bytes())
}

func main() {
	payload := []byte("Hello world - catfact getcatfact putcatfact")
	hashAndBroadcast(NewHashReader(payload))
}

func hashAndBroadcast(r HashReader) error {
	// b, err := io.ReadAll(r)

	// if err != nil {
	// 	return err
	// }

	// hash := sha1.Sum(b)
	hash := r.hash()

	fmt.Println(hash)

	return broadcast(r)
}

func broadcast(r io.Reader) error {
	b, err := io.ReadAll(r)

	if err != nil {
		return err
	}

	fmt.Println("String of the bytes: ", string(b))

	return nil
}
