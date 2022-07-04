package main

import "fmt"

func Greeting(prefix string, who ...string) {
	if who == nil {
		fmt.Printf("Nobody to say hi.")
	}

	for _, people := range who {
		fmt.Printf("%s %s\n", prefix, people)
	}
}

func main() {
	Greeting("nobody")
	Greeting("hello:", "Kobe", "Duncan", "Yao")
	guest := []string{"Yao", "Kobe", "Duncan"}
	Greeting("hi:", guest...)
}
