// チャネルは生合成が壊れることがない安全なデータ構造になっている
// 同時に複数のgoroutineでチャネルに読み書きを行なっても、1つの
// goroutineだけがデータを投入できる. 取り出しも同時に1つだけ
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るためのチャネル
	done := make(chan bool)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
		// 終了を通知
		done <- true
	}()

	// 終了を待つ
	<-done
	fmt.Println("all tasks are finished")
}
