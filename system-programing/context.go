//　コンテキストは深いネストの中でも正しい終了やキャンセル、タイムアウトが実装できるようにする仕組み
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")

	// 終了を受け取るための終了関数付きコンテキスト
	// 終了を受け取るコンテキストctx と そのコンテキストを終了させるcancel関数
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")

		// 終了を通知
		cancel()
	}()

	// 終了を待つ
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
