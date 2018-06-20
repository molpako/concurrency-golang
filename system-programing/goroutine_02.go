package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start sub()")
	// インラインで無名関数を作ってその場でgorouineで実行
	// go の後ろには関数名ではなく「関数呼び出し分」を書くので末尾に () が必要
	go func() {
		fmt.Println("sub() is running")
		time.Sleep(time.Second)
		fmt.Println("sub() is finished")
	}()
	time.Sleep(2 * time.Second)
}
