package main

import "fmt"

func main() {
	a := 10
	a1 := nextInt(a)

	fmt.Println(a1)
}

func nextInt(i int) int {
	return i + 1
}
