package main

import (
	"fmt"
	"time"
)

// This is not a part of the Go Tour. Rather it is from https://www.geeksforgeeks.org/buffered-channel-in-golang/
// I'm including it here because it provides better understanding of how the blocks work with Go routines and channels.

func write(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}
func main() {

	// creates capacity of 2
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	}
}
