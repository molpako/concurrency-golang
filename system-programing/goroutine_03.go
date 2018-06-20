// gorouine で協調動作をするには goroutine を実行する親スレッドとこの間でデータのやりとりが必要
package main

import (
	"fmt"
	"time"
)

func sub1(c int) {
	fmt.Println("share by arguments:", c*c)
}

func main() {
	// 引数渡し
	go sub1(10)

	// クロージャのキャプチャ渡し
	c := 20
	go func() {
		fmt.Println("share by capture", c*c)
	}()
	time.Sleep(time.Second)
}
