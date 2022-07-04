package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	ch := make(chan int, 10)

	go func() {
		var i = 1
		for {
			i++
			ch <- i
		}
	}()

	for {
		select {
		case x := <-ch:
			println(x)
		case <-time.After(3 * time.Minute):
			println(time.Now().Unix())
		}
	}
}
