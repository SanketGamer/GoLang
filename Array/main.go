package main

import "fmt"

func main() {
	var name [5]string
	name[0] = "ssr"
	name[2] = "Aakash"
	fmt.Println("Names of Personality",name)

	var numbers = [5]int{1,2,3,4,51}
	fmt.Println(numbers)
	fmt.Println("Length of person",len(numbers))

	fmt.Println("Value of name at 2nd index is",name[2])

	//in Go integer default value always set as 0
	//in string empty space or string
	var price[5]string
	fmt.Println(price)
	fmt.Printf("number is %q\n",price)
}