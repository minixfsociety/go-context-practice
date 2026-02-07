package main

import (
	"context"
	"time"
)

func Hotel(ctx context.Context, ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Hotel found for $300"
}
