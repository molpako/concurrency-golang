// 並行・並列化の第一歩は「重い処理をタスクに分けること」
// 重い処理をgoroutineのなかで実行して非同期化する
package main

import (
	"io/ioutil"
	"os"
)

func main() {
	inputs := make(chan []byte)

	// create temp file
	tmpfile1, _ := ioutil.TempFile("", "example")
	defer os.Remove(tmpfile1.Name())

	tmpfile2, _ := ioutil.TempFile("", "example")
	defer os.Remove(tmpfile2.Name())

	go func() {
		a, _ := ioutil.ReadFile(tmpfile1.Name())
		inputs <- a
	}()

	go func() {
		b, _ := ioutil.ReadFile(tmpfile2.Name())
		inputs <- b
	}()

}
