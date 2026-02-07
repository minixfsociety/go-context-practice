package main

import (
	"context"
	"fmt"
)

func FinalSearch(ctx context.Context) {
	ch := make(chan string)
	go Air(ctx, ch)
	go Hotel(ctx, ch)
	for i := 1; i < 3; i++ {
		select {
		case msq := <-ch:
			fmt.Println("Message:", msq)
		case <-ctx.Done():
			fmt.Println("Time's up :(")
			return
		}
	}
}
