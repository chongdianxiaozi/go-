package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "parameter", "1")
	go HandelRequest3(ctx)

	time.Sleep(10 * time.Second)
}
