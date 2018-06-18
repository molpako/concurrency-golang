// チャネルの送受信を通じて処理結果の取得を必要になるとこまで後回しにするパターン
package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func loadGameStage() chan string {
	promise := make(chan string)
	go func() {
		fmt.Println("load stage")
		time.Sleep(2 * time.Second)
		promise <- "done stage"
	}()
	return promise
}

func loadGameResource(futureStage <-chan string) chan []string {
	promise := make(chan []string)
	go func() {
		resource := []string{<-futureStage}
		var mu sync.Mutex
		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			fmt.Println("load character")
			time.Sleep(1 * time.Second)
			mu.Lock()
			resource = append(resource, "done character")
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()
			fmt.Println("load filed material")
			time.Sleep(1 * time.Second)
			mu.Lock()
			resource = append(resource, "done field material")
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()
			fmt.Println("load property")
			time.Sleep(1 * time.Second)
			mu.Lock()
			resource = append(resource, "done property")
			mu.Unlock()
		}()
		wg.Wait()
		promise <- resource
	}()
	return promise
}

func loadGameMenu(futureResource <-chan []string) chan []string {
	promise := make(chan []string)
	go func() {
		resource := <-futureResource
		fmt.Println("load menu frame")
		time.Sleep(time.Second)
		menu := append(resource, "done menu frame")
		promise <- menu
	}()
	return promise
}

func main() {
	futureStage := loadGameStage()
	fmt.Println("loading stage...")
	futureResource := loadGameResource(futureStage)
	futureMenu := loadGameMenu(futureResource)
	fmt.Println(strings.Join(<-futureMenu, "\n"))
}
