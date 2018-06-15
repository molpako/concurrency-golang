// 基本的な並行パターン
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("-- 1 --")
	go func() {
		fmt.Println("-- 2 --")
		time.Sleep(2 * time.Second)
		fmt.Println("-- 3 --")
	}()
	// goroutine の開始後、3秒sleep
	time.Sleep(3 * time.Second)
}
