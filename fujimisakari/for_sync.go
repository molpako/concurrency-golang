// ループ毎の処理全てを並列に実行したい場合に利用したいパターン
// sync.WaitGroupで全てのgoroutineの終了を待つ
package main

import (
	"fmt"
	"sync"
	"time"
)

func printer(msg string) {
	time.Sleep(1 * time.Second)
	fmt.Println(msg)
}

func main() {
	var messages = []string{
		"test1",
		"test2",
		"test3",
		"test4",
		"test5",
	}

	var wg sync.WaitGroup

	for _, msg := range messages {
		wg.Add(1)
		go func(m string) {
			defer wg.Done()
			printer(m)
		}(msg)
	}
	wg.Wait()
}
