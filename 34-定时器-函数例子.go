package main

import (
	"log"
	"time"
)

func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(1 * time.Second)

	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C: // 超时
		println("WaitChannel timeout!")
		return false
	}
}

func DelayFunction() {
	timer := time.NewTimer(5 * time.Second)

	select {
	case <-timer.C:
		log.Println("Delayed 5s, start to do something.")
	}
}

func AfterDemo() {
	log.Println(time.Now())
	<-time.After(1 * time.Second)
	log.Println(time.Now())
}

func AfterFuncDemo() {
	log.Println("AfterFuncDemo start: ", time.Now())
	time.AfterFunc(1*time.Second, func() {
		log.Println("AfterFuncDemo end: ", time.Now())
	})

	time.Sleep(2 * time.Second) // 等待协程退出
}

func TickerDemo() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Ticker tick.")
	}
}

func TickerLaunch() {
	ticker := time.NewTicker(5 * time.Minute)
	maxPassenger := 30 // 每车最大装载人数
	passengers := make([]string, 0, maxPassenger)

	for {
		passenger := GetNewPassenger() // 获取一个新乘客
		if passenger != "" {
			passengers = append(passengers, passenger)
		} else {
			time.Sleep(1 * time.Second)
		}

		select {
		case <-ticker.C: // 时间到，发车
			Launch(passengers)
			passengers = []string{}
		default:
			if len(passengers) >= maxPassenger { // 时间未到，车已满座，发车
				Launch(passengers)
				passengers = []string{}
			}
		}
	}
}

func GetNewPassenger() string {
	return "one"
}

func Launch(passengers []string) []string {
	return passengers
}

func WrongTicker() {
	for {
		select {
		case <-time.Tick(1 * time.Second):
			log.Printf("Resource leak!")
		}
	}
}

func main() {

}
