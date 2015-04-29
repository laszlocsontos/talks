package main

import "fmt"

const (
	N = 100
	K = 10
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}

	c <- sum // send sum to c
}

func main() {
	a := make([]int, N, N)
	for i := 0; i < N; i++ {
		a[i] = i + 1
	}

	c := make(chan int)

	for i := 0; i < N/K; i++ {
		start := i * K
		end := (i+1)*K

		go sum(a[start:end], c)
	}

	total := 0
	for i := 0; i < N/K; i++ {
		total += <- c
	}

	fmt.Println(total)
}

