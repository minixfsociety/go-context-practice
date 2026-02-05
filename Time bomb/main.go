package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	bomb := make(chan bool)
	go defuse(bomb)
	fmt.Println("ждем")
	select {
	case <-ctx.Done():
		fmt.Printf("Мы ждали слишкой долго: %v", ctx.Err())
	case result := <-bomb:
		fmt.Println(result)
	}
}

func defuse(ch chan bool) {
	time.Sleep(5 * time.Second)
	fmt.Println("Бомба унечтожина")
	ch <- true
}
