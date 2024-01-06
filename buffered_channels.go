package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	// Since the channel has a buffer of size 2, we can't send to
	// the channel until something is received from the channel.
	// Therefore, this will cause an error.
	ch <- 3
	ch <- 4

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
