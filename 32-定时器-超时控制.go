package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTimer(1 * time.Second)
	defer t.Stop()

	go func() {
		select {
		case <-doWork():
			fmt.Println("正常退出")
			return
		case <-t.C:
			fmt.Println("超时")
			return
		}
	}()

	time.Sleep(5 * time.Second)
}

func doWork() chan struct{} {
	var ch = make(chan struct{})

	go func() {
		// 模拟超时
		time.Sleep(2 * time.Second)
		ch <- struct{}{}
	}()
	return ch
}
