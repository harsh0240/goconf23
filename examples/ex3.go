package main

import "fmt"

func getOne() *int {
	i := 1
	return &i
}

func main() {
	a := getOne()

	fmt.Println(*a)
}
