package codec

import (
	"encoding/binary"
	"io"
)

func EncodeString(s string) []byte {
	b := make([]byte, 0)
	b = binary.BigEndian.AppendUint16(b, uint16(len(s)))
	b = append(b, []byte(s)...)
	return b
}

func DecodeString(r io.Reader) string {
	uintbuf := make([]byte, 2)
	n, err := r.Read(uintbuf)
	if err != nil {
		panic(err)
	}
	if n != 2 {
		panic("Not enought bytes")
	}
	len := binary.BigEndian.Uint16(uintbuf)
	strbuf := make([]byte, len)
	n, err = r.Read(strbuf)
	if err != nil {
		panic(err)
	}
	if n != int(len) {
		panic("Not enought bytes")
	}
	return string(strbuf)
}
