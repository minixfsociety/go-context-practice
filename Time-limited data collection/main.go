package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()
	ch := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 500
	}()
	select {
	case <-ctx.Done():
		fmt.Println("Поиск занял слишком много времени")
	case result := <-ch:
		fmt.Println("Билет найден за ценой:", result)
	}
}
