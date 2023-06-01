package main

import "fmt"

func main() {
	a := 10
	a1 := nextIntWithPtr(&a)

	fmt.Println(a1)
}

func nextIntWithPtr(i *int) int {
	return *i + 1
}
