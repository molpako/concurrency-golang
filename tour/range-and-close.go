// 送信側はもう送信する値がない時それを示すためチャネルを close できる
// 受信側は2つ目の返り値でそのチャネルが close しているか確認できる（errみたいな感じ
// 通常は close する必要はない. close するのは、これ以上値がこないことを受信側が知る必要が
// ある時だけ. 例えば range ループを終了するという場合.
package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		time.Sleep(100 * time.Millisecond)
		c <- x
		x, y = y, y+x
	}
	// 送信側のチャネルを close する
	close(c)
}

func main() {

	// バッファとして使うチャネルの生成, 第2引数にバッファの長さを指定
	n := 10
	ch1 := make(chan int, n)

	go fibonacci(cap(ch1), ch1)

	// チャネル c が閉じられるまでチャネルから値を繰り返し受信し続ける
	for i := range ch1 {
		fmt.Println(i)
	}

	var v int
	var ok bool

	ch2 := make(chan int, 10)

	go fibonacci(cap(ch2), ch2)

	// 受信する値がない、かつチャネルが閉じられているなら
	// ok は false になる
	for i := 0; i < n+1; i++ {
		v, ok = <-ch2
		fmt.Printf("値: %v, received: %v\n", v, ok)
	}
}
