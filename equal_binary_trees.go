package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	walkRecurse(t, ch)
}

func walkRecurse(t *tree.Tree, ch chan int) {

	// check the left side
	if t.Left != nil {
		walkRecurse(t.Left, ch)
	}

	// save the node value
	ch <- t.Value

	// check the right side
	if t.Right != nil {
		walkRecurse(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {

	// Create the channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// kick off the Go routines
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	// Loop until there are no more values from ch1 or ch2
	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2

		if n1 != n2 {
			return false
		}
		if !ok1 || !ok2 {
			// no more values
			break
		}
	}
	return true
}

func main() {
	same1 := Same(tree.New(1), tree.New(1))
	fmt.Printf("tree.New(1) the same as tree.New(1)? %v\n", same1)

	same2 := Same(tree.New(1), tree.New(2))
	fmt.Printf("tree.New(1) the same as tree.New(2)? %v\n", same2)
}
