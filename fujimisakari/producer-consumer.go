// タスク生成と処理を分ける: Producer-Consumer パターン
// 同時に複数のタスクを実行した場合に利用するパターン
// `buffer.go` との違いはタスク数の制限方法がバッファチャネルではなくワーカー数であること
// チャネルでProducerとConsumerを接続する
// チャネルは、複数のgoroutineで同時に読み込みを行なっても、
// 必ず一つのgoroutineだけが結果を受け取れる（消失したり、複製できてしまわない）
// したがって、Consumer側を増やすことで、安全に処理速度をスケールできる
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
