package main

import "fmt"

func main(){
	age:=25
	name:="Rahul"
	height:=8.624

	fmt.Println(age,name,height)
	fmt.Println("hello world")
	fmt.Printf("age is %d\n down",age) //printf means print same line 
	fmt.Printf("height is %.2f\n",height)   //.2 means it shows me 2 digit 
	fmt.Printf("Type of name is %T\n",name)  // %T means type of which data type u using
	fmt.Printf("name is %s\n",name) //%s means string
}