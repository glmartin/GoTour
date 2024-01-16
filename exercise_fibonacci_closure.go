package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
	f0 := 0
	f1 := 1
	return func(i int) int {
		if i < 2 {
			return i
		} else {
			sum := f0 + f1
			f0 = f1
			f1 = sum
			return sum
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
