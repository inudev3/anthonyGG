package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()
	ch := make(chan any, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go fetchUserLikes(userName, ch, wg)
	go fetchUserMatch(userName, ch, wg)
	wg.Wait()
	close(ch)
	for res := range ch {
		fmt.Println("res: ", res)
	}

	fmt.Println("took: ", time.Since(start))
}
func fetchUser() string {
	time.Sleep(time.Millisecond * 100)
	return "BOB"
}
func fetchUserLikes(username string, ch chan<- any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)
	ch <- 11
	wg.Done()
}
func fetchUserMatch(username string, ch chan<- any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)
	ch <- "ANNA"
	wg.Done()
}
