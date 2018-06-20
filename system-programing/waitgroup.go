// 多数のgoroutineで実行しているジョブの終了待ちに使います
// ジョブ数が大量にあったり、可変個の場合、チャネルよりはやい
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// ジョブ数をあらかじめ登録
	// 必ずgoroutineの作成前に呼び出す
	wg.Add(2)

	go func() {
		// 非同期で仕事をする(1)
		fmt.Println("仕事１開始")
		time.Sleep(2 * time.Second)
		fmt.Println("仕事１終了")

		// Doneで完了を通知, ジョブ数をデクリメント
		wg.Done()
	}()

	go func() {
		// 非同期で仕事をする(2)
		fmt.Println("仕事２開始")
		fmt.Println("仕事２終了")

		// Doneで完了を通知, ジョブ数をデクリメント
		wg.Done()
	}()

	// 全ての処理が終わるのを待つ
	wg.Wait()
	fmt.Println("全部終了")
}
