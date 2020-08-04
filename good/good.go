package main

import "fmt"

func main() {
	s := make([]func(), 4)

	for i := 0; i < 4; i++ {
		s[i] = func() {
			j := i
			fmt.Printf("%d %p\n", j, &j)
		}
	}

	for i := 0; i < 4; i++ {
		s[i]()
	}
}
