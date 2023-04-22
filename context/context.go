package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	userId := 10
	val, err := fetchUserData(context.Background(), userId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
	fmt.Println("took:", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	ch := make(chan Response)
	go func() {
		val, err := fetchSthThirdParty()
		ch <- Response{val, err}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("timeout fetching data")
		case resp := <-ch:
			return resp.value, resp.err
		}
	}
}
func fetchSthThirdParty() (int, error) {
	time.Sleep(201 * time.Millisecond)
	return 666, nil
}
