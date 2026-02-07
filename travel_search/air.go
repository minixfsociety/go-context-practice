package main

import (
	"context"
	"time"
)

func Air(ctx context.Context, ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Tickets found for $500"
}
