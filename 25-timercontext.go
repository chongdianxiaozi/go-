package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest2(ctx context.Context) {
	go WriteRedis2(ctx)
	go WriteDatabase2(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running.")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteRedis2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteRedis Done.")
			return
		default:
			fmt.Println("WriteRedis running.")
			time.Sleep(2 * time.Second)
		}
	}
}

func WriteDatabase2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("WriteDatabase Done.")
			return
		default:
			fmt.Println("WriteDatabase running.")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go HandelRequest2(ctx)

	time.Sleep(10 * time.Second)
}
