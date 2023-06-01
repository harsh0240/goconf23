package main

import (
	"bytes"
	"testing"
)

// START OMIT

func BenchmarkRead1(b *testing.B) { // HL
	buf := bytes.NewBufferString("gocon23")
	bt := make([]byte, 10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		read1(buf, bt)
	}
}

func BenchmarkRead2(b *testing.B) { // HL
	mbuf := &MyBuffer{buffer: bytes.NewBufferString("gocon23")}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		read2(mbuf)
	}
}

// END OMIT
