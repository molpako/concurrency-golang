// どの case も準備ができていなければ、 select の中の default が実行される
// ブロックせずに送受信を行うなら default を使う
package main

import (
	"fmt"
	"time"
)

func main() {
	// 指定した間隔でチャネルを返す
	tick := time.Tick(100 * time.Millisecond)

	// 指定した時間後にチャネルを返す
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

}
