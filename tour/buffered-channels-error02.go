// エラーコード. 実行に失敗します.
// バッファが空の時には、チャネルの受信をブロックする
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	fmt.Println("バッファが空の状態で、チャネルを受信する")
	fmt.Println(<-ch)
}
