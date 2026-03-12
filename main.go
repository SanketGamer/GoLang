package main

import (
	"fmt"
	// "mylearning/myutill"
)

func main(){
	fmt.Println("hello world")
	//myutill.PrintMsg("hello world");

	//declare variables
	var name string="sanket"
	var version = 76
	fmt.Println(name)
	fmt.Println(version)

	var money int=1000
	var currency=500
	fmt.Println(money)
	fmt.Println(currency)

	var dimesion float64=87.12
	fmt.Println(dimesion)

	var flag bool=true
	fmt.Println(flag)

	const data=67.12
	fmt.Println(data)

	//optimize way : - define as a datatype
	person :=123
	fmt.Println(person)

	//if we need any variable in other file or package we have uppercase 1st word but if we do same file that variable we asign small all the word same as function alsp
	var Public="data is important"
	var private="data is private"

	fmt.Println(Public)
	fmt.Println(private)

	// func PublicFunc(){
	// 	fmt.Println("public so 1st word uppercase")
	// }
	// func privateFunc(){
	// 	fmt.Println("private so small name")
	// }
}

