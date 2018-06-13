// チャネルはバッファとしても使える
// バッファを持つチャネルを初期化するには make の
// 引数にバッファの長さを与える
package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
