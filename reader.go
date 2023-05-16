package main

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

import _ "net/http/pprof"

type MyBuffer struct {
	buffer *bytes.Buffer
}

func (mb *MyBuffer) Read() ([]byte, error) {
	b := make([]byte, 1000000)

	copy(b, mb.buffer.Bytes())
	return b, nil
}

const str = "proposaltitle:\"Escape Analysis in Go: Understanding and Optimizing Memory Allocation\"Abstract:Go is a popular programming language known for its performance and efficiency. However, in order to achieve maximum performance, it's important to understand how Go manages memory. One of the key concepts in this area is escape analysis, a technique used by the Go runtime to determine when and where to allocate memory on the heap. In this talk, we will cover the basics of escape analysis in Go, including how it works, common pitfalls to watch out for, and techniques for optimizing memory allocation in your Go programs.Outline: [20 minutes talk]Introduction to escape analysis in Go [8 minutes]Recap of memory allocation (stack/heap) and pointers Experiments to prove why escape analysis is importantExplanation of what escape analysis is and how it works in GoHow does escape analysis affect memory allocationCommon pitfalls to watch out for [4 minutes]How to avoid unnecessary heap allocationSome special cases for escape analysisTechniques for optimizing memory allocation in Go programs [5 minutes]How to use the \"go build\" flag to check for escape analysisBest practices for variable declarationMeasuring the performance improvementsConclusion [3 minutes]Summary of key takeawaysAdditional resources for further learningOriginality:The Go documentation provides some information about what escape analysis is and a general overview of how that works. This talk will extend those ideas and discuss why it is important to know about this optimization and the cases where it works/not works. The talk will also include some major guidelines about how to use it for the general audience to consider. Target Audience:This talk is aimed at Go developers of all levels who are interested in improving the performance of their applications. Whether you are a beginner just starting out with Go or an experienced developer looking to optimize your code, this talk will provide valuable insights and information. The following groups will particularly benefit from this talk:Beginner Go developers: This talk will provide a comprehensive introduction to escape analysis and its benefits, making it an excellent starting point for those new to the language.Experienced Go developers: For those familiar with Go, this talk will provide a deeper understanding of escape analysis and its implementation, as well as best practices for applying it in real-world projects.Performance-focused developers: This talk will provide valuable information for developers who are focused on optimizing the performance of their applications. Attendees will learn about escape analysis and how it can be used to reduce the number of heap allocations and improve the performance of their applications.Background:I am a software engineer with a year of experience in developing Go-based applications. I have experience in Go memory management and optimization, and I am passionate about helping developers write more efficient Go code.I got introduced to this topic while finding ways to optimize the Go code. I was curious about how the memory allocations in Go work and found something related to Escape Analysis for the first time. As I learned more, I found this optimization beneficial and could be utilized more.I am confident that this talk would be of great interest to the attendees of Go Conference 2023 as I will be sharing my knowledge on a topic that is crucial for developers who want to achieve maximum performance in their Go-based applications."

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	b := make([]byte, 1000000)

	buf := bytes.NewBufferString(str)
	mbuf := &MyBuffer{
		buffer: bytes.NewBufferString(str),
	}

	read1(buf, b)
	read2(mbuf)
	time.Sleep(10 * time.Minute)
}

//go:noinline
func read1(buf *bytes.Buffer, b []byte) {
	read, err := buf.Read(b)
	if err != nil {
		return
	}
	println(b, read)
}

//go:noinline
func read2(buf *MyBuffer) {
	myRead, err := buf.Read()
	if err != nil {
		return
	}
	time.Sleep(10 * time.Minute)
	println(myRead)
}
