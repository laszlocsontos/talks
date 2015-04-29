package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}

	for i := 0; i < len(s); i++ {
		v := s[i]
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	for i, v := range s {
		fmt.Printf("index: %d, value: %v\n", i, v)
	}

	m := map[int]string{1: "a", 2: "b", 3: "c"}

	for k, v := range m {
		fmt.Printf("key: %d, value: %v\n", k, v)
	}
}

