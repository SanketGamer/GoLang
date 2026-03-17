package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Say Hello")
	//time.Sleep(3000 * time.Millisecond)
	fmt.Println("Say again Hello")
}

func sayHi() {
	fmt.Println("Hi Rahul")
}
func main() {

	fmt.Println("Start learning goroutines")
	go sayHello()  //run this function in background (concurrently)
	go sayHi()
	time.Sleep(1000 * time.Millisecond)
	
}
