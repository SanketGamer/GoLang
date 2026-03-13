package main

import "fmt"

func simpleFunction(){

}

func add(a,b int)(int){   //2nd bracket work as output what datatype and 1st one is as argument datatype
return a+b
}

func sub(a,b int)(res int){  //same but we difined res 1st because we have to return that value
	res=a+b
	return res
}

func main() {
	fmt.Println("We are learning functions in GO lang")
	simpleFunction()

	ans:=add(4,6)
	fmt.Println("add of two numbers is:",ans)
}