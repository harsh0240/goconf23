package main

import "bytes"

type MyBuffer struct { // HL
	buffer *bytes.Buffer
}

func (mb *MyBuffer) Read() ([]byte, error) { // HL
	b := make([]byte, mb.buffer.Len())

	copy(b, mb.buffer.Bytes())
	return b, nil
}
