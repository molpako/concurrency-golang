// forループ内を全てgoroutineで実行すれば、並列化する
// ループ変数の実体は一つしかないため、goroutineの引数として渡し、
// goroutineごとにコピーが作られるようにする
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	tasks := []string{
		"cmake",
		"cmake .--build Release",
		"cpack",
	}
	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(task string) {
			// ジョブを実行
			// このサンプルでは出力だけする
			fmt.Println(task)
			time.Sleep(2 * time.Second)
			wg.Done()
		}(task)
	}
	wg.Wait()
}
