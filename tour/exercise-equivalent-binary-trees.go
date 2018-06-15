package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// import した Tree の構造体
//type Tree struct {
//    Left  *Tree
//    Value int
//    Right *Tree
//}

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

// Same は t1, t2 が同じ値を含むか どうか調べる
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for v := range ch1 {
		if v != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)

	// tree.New(int) はソートされた二分木を生成する
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
