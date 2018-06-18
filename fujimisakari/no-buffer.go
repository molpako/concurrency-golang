// 重い処理を非同期で実行したい場合に利用されるパターン
package main

import (
	"log"
	"time"
)

func main() {
	done := make(chan struct{})
	log.Println("start")
	go func() {
		time.Sleep(time.Second)
		log.Println("done")

		// チャネルの送信
		done <- struct{}{}
	}()
	log.Println("between")

	// チャネルの受信
	<-done
}
