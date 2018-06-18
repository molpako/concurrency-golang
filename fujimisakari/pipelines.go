// タスク分割したgoroutineを順序つけて処理したい場合に利用するパターン
package main

import (
	"fmt"
	"log"
)

func double(naturals <-chan int, doubles chan<- int) {
	for {
		x, ok := <-naturals
		log.Println("naturals:", x, ok)
		if !ok {
			close(doubles)
			break
		}
		doubles <- x * 2
	}
}

func squarer(doubles <-chan int, squares chan<- int) {
	for {
		x, ok := <-doubles
		log.Println("doubles", x, ok)
		if !ok {
			close(squares)
			break
		}
		squares <- x * x
	}
}

func main() {
	naturals := make(chan int)
	doubles := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 1; i < 10; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go double(naturals, doubles)
	go squarer(doubles, squares)

	for result := range squares {
		fmt.Println(result)
	}
}
