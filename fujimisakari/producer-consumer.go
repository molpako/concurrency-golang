// 同時に複数のタスクを実行した場合に利用するパターン
// `buffer.go` との違いはタスク数の制限方法がバッファチャネルではなくワーカー数であること
package main

import (
	"fmt"
	"sync"
	"time"
)

// キューを取り出し、処理をする
func worker(sem <-chan int, wg *sync.WaitGroup) {
	for num := range sem {
		time.Sleep(1 * time.Second)
		fmt.Println("process", num)
		wg.Done()
	}
}

func main() {
	sem := make(chan int)
	var wg sync.WaitGroup

	// Consumer, ワーカープール
	for i := 0; i < 20; i++ {
		go worker(sem, &wg)
	}

	// Producer, タスク生成
	for i := 0; i <= 200; i++ {
		wg.Add(1)
		// タスク生成したら、データをキューに追加
		go func(x int) {
			sem <- x
		}(i)
	}
	wg.Wait()
	close(sem)
}
