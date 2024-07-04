package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch) // <- closes the channel when this function returns
	
	var walk func(t *tree.Tree)
	
	walk = func(t *tree.Tree) {
			if t == nil {
					return
			}
			
			walk(t.Left)
			ch <- t.Value
			walk(t.Right)
	}
	
	walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	values1 := make([]int, 10)
	values2 := make([]int, 10)

	for i := 0; i < 10; i++ {
		val1 := <-ch1
		val2 := <-ch2
		
		

		values1[i] = val1
		values2[i] = val2
	}

	for {
		if len(values1) == 0 {
			break
		}
		
		index := Index(values2, values1[0])

		if index == -1 {
			return false
		}

		values1 = values1[1:]
	}

	return true
}

func main() {

	fmt.Printf("Same: %v", Same(tree.New(10), tree.New(10)))
}