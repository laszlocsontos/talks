package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, s=%v\n", len(s), cap(s), s)
}

func main() {
	var s []int
	printSlice(s)

	s = []int {1,2,3,4}
	printSlice(s)

	s = append(s, 5)
	printSlice(s)
}

