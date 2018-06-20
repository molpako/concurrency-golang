// 一度だけ関数を実行したい時に使う sync.Once
// 初期化処理を一度だけ行いたい時に使う
// init()を使う方がメジャー. 初期化処理を必要な時まで遅延させたい時に sync.Once
package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("初期化処理")
}

var once sync.Once

func main() {
	// 三回呼び出しても一度しか呼ばれない

	once.Do(initialize)
	once.Do(initialize)
	once.Do(initialize)
}
