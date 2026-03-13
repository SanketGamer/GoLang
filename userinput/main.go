package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, What's your name?")
	 //var name string
	 //1st way
	// fmt.Scan(&name)  //in go how we take input
	// fmt.Println("Hello, mr. ",name)
	//2nd way
	reader:=bufio.NewReader(os.Stdin)  //creates a new buffer reader that reads from standard input(os.Stdin)
	name, _ :=reader.ReadString('\n')
	fmt.Println("hello,Mr.",name)
	
}