// select ステートメントは、goroutine を複数の通信操作で待たせる
// select は、複数ある case のいずれかが準備できるようになるまでブロックし、準備ができた case を実行する。
// もし、複数の case の準備ができている場合、 case はランダムに選択されます。
package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	fmt.Println("fibonnacci start")
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	// 2つチャネル生成する
	c := make(chan int)
	quit := make(chan int)

	go func() {
		fmt.Println("goroutine start")
		for i := 0; i < 10; i++ {
			// チャネルを受信するまでブロッキング
			fmt.Printf("受信したチャネル: %v\n", <-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
