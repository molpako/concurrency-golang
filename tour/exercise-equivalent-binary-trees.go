package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk はツリーからチャネルに全ての値を送信する
func Walk(t *tree.Tree, ch chan int) {
	var walker func(*tree.Tree)

	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}
}
