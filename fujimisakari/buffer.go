// 同時に実行できるタスク数（非同期I/O）を制限したい場合に利用するパターン
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetch(sem chan struct{}, url string) {
	// チャネルへ値を送信（エンキュー）
	sem <- struct{}{}

	// 終了時に受信（デキュー ）
	defer func() { <-sem }()

	http.Get(url)
	fmt.Println("fetched", url)
}

func main() {
	// 指定した数分、生成したチャネルをキューとして扱う
	sem := make(chan struct{}, 10)

	// 複数のgoroutineの完了を待つための値
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			url := "https://golang.org"
			fetch(sem, url)
		}()
	}
	wg.Wait()

}
