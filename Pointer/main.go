package main

import "fmt"

//in Go,pointer is a variable that stores the memory address of another variable.
//Pointer are used to indirectly refer to the value stored in a variable,rather than the value itself
//simple word pointer dusre variable ka data store nhi karta bo address store karta hey.



func modyfyValueReference(num *int){
  *num=*num*20
}
func main() {
	num := 2

	var ptr *int
	ptr = &num
	fmt.Println(ptr)
	fmt.Println(*ptr)

	//use case
	var pointer *int  //default pointer is null
	if pointer == nil {
		fmt.Println("Pointer is not assigned")
	}

	value:=10
	modyfyValueReference(&value)
	fmt.Println(value)

}
