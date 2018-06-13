// エラーコード. 実行に失敗します.
package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	fmt.Println("バッファが空の状態で、チャネルを受信する")
	fmt.Println(<-ch)
}
