package main

import (
	"fmt"
	"strconv"
)

func main() {
	// num := 42
	// fmt.Println(num)
	// fmt.Printf("%T",num)

	// var data float64=float64(num)
	// fmt.Println("Data is :",data)

	//int to string
	num:=123
	str:=strconv.Itoa(num)  //in built func to change num->string
	fmt.Println(str)
	fmt.Printf("%T\n",str)

	//string to int
	// number_string:="1234"
	// number_int,_:=strconv.Atoi(number_string) //string->int
	// fmt.Printf("%T\n",number_int)
	// fmt.Println(number_int)

	//string to float
	number_string:="3.14"
	number_float,_:=strconv.ParseFloat(number_string,64) //string->int
	fmt.Printf("%T\n",number_float)
	fmt.Println(number_float)
}