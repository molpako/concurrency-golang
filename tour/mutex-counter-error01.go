// goroutine間で通信が必要ない場合は、コンフリクトを避けるため一度に一つの
// goroutineだけが変数にアクセスできるようにしたい場合（排他制御）は mutex を使用する
package main

import (
	"sync"
)

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc は指定されたキーのカウンタをインクリメントする
// Lock と Unlock で囲むことで排他制御で実行するコードを定義できる
func (c *SafeCounter) Inc(key string) {
	// ロックせずにインクリメントするとエラーが起きる
	c.v[key]++
}

// Value は指定されたキーのカウントの値を返す
// mutexがUnlockされることを保証するために defer を使うこともできる
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()

	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 100; i++ {
		go c.Inc("somekey")
	}

	//time.Sleep(time.Second)
	//fmt.Println(c.Value("somekey"))
}
