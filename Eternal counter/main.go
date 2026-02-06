package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go counter(ctx, 1)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func counter(ctx context.Context, number int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Счетчик остановлен")
			return
		default:
			fmt.Println("Текущее число:", number)
			number++
			time.Sleep((200 * time.Millisecond))
		}
	}
}
