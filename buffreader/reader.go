package main

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

import _ "net/http/pprof"

const str = "hello"

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	b := make([]byte, 10)

	buf := bytes.NewBufferString(str)
	mbuf := &MyBuffer{
		buffer: bytes.NewBufferString(str),
	}

	read1(buf, b)
	read2(mbuf)
	time.Sleep(10 * time.Minute)
}

// START OMIT

//go:noinline
func read1(buf *bytes.Buffer, b []byte) int { // HL
	read, err := buf.Read(b)
	if err != nil {
		return 0
	}
	return read
}

//go:noinline
func read2(buf *MyBuffer) []byte { // HL
	myRead, err := buf.Read()
	if err != nil {
		return nil
	}
	return myRead
}

// END OMIT
