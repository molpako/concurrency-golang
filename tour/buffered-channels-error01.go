// エラーコード. 実行に失敗します.
// バッファが詰まった時は、チャネルの送信をブロックします.
package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("バッファを詰まらせる")
	ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
