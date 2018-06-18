// 複数のチャネルを同時に扱いたい場合に利用するパターン
package main

import (
	"fmt"
	"sync"
)

type Server struct {
	req chan string
	res chan string
}

func (s Server) Start() {
	worker := make(chan string, 30)
	result := make(chan string)

	go listen(s, worker, result)
	go response(worker, result)
}

func listen(server Server, worker chan<- string, result <-chan string) {
	for {
		select {
		case req := <-server.req:
			worker <- req
		case res := <-result:
			server.res <- res
		}
	}
}

func response(worker <-chan string, result chan<- string) {
	for request := range worker {
		go func(req string) {
			result <- fmt.Sprintf("response from %s", req)
		}(request)
	}
}

func request(server Server, reqNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	server.req <- fmt.Sprintf("request %d", reqNum)
	fmt.Println(<-server.res)
}

func main() {
	server := Server{req: make(chan string), res: make(chan string)}
	server.Start()

	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go request(server, i, &wg)
	}
	wg.Wait()
}
