package main

func readChan(chanName <-chan int) {
	println("readChan")
	val := <-chanName
	println(val)
}

func writeChan(chanName chan<- int) {
	println("writeChan")
	chanName <- 1
}

func main() {
	var myChan = make(chan int, 10)
	writeChan(myChan)
	readChan(myChan)
}
