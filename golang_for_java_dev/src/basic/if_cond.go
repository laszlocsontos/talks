package main

import "fmt"

func main() {
	s := []int{3, 42, 73, 1}

	max := -1

	for i := 0; i < len(s); i++ {
		if v := s[i]; v > max {
			max = v
		}
	}

	fmt.Println(max)
}

