package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("done \n")
				return
			default:
				fmt.Println("wait 1")
				time.Sleep(2 * time.Second)

			}

		}
	}(ctx)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("done 2  \n")
				return
			default:
				fmt.Println("wait 2")
				time.Sleep(2 * time.Second)

			}

		}

	}(ctx)

	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(6 * time.Second)
}
