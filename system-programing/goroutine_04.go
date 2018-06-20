package main

import (
	"fmt"
	"time"
)

func main() {
	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		go func() {
			// goroutine が起動するときにはループが回りきって
			// 全部の task が最後のタスクになってしまう
			fmt.Println(task)
		}()
	}
	time.Sleep(time.Second)
}
