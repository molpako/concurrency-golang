// チャネルは、チャネルオペレータ(<-)を用いて値の送受信ができる通り道
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// sum をチャネル c に送信する
	c <- sum
}
func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// マップとスライスのようにチャネルは使う前に生成する
	c := make(chan int)

	// 2つのgoroutine間で作業を分配する。両方の計算が終わると
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// receive from c
	// x には1つ目、 y には2つ目のgoroutineの結果が入る
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

}
