package main

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	justWalk(t, ch)
	close(ch)
}

func justWalk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		justWalk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		justWalk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for v1 := range ch1 {
		v2 := <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		println(i)
	}

	trueExpected := Same(tree.New(1), tree.New(1))
	falseExpected := Same(tree.New(1), tree.New(2))
	println(trueExpected, falseExpected)
}
