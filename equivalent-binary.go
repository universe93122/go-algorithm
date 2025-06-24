package main

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walk func(*tree.Tree)
	walk = func(node *tree.Tree) {
		if node == nil {
			return
		}
		walk(node.Left)
		ch <- node.Value
		walk(node.Right)
	}
	walk(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			// One channel closed before the other
			return false
		}
		if !ok1 {
			// Both channels closed = done
			break
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}
