package main

import "bytes"

const SIZE = 8

type MyBuffer struct { // HL
	buffer *bytes.Buffer
}

func (mb *MyBuffer) Read() ([]byte, error) { // HL

	b := make([]byte, mb.buffer.Len(), SIZE)

	copy(b, mb.buffer.Bytes())
	return b, nil
}
