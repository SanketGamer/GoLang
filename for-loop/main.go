package main

import (
	"fmt"
)

//in Go the for loop is the Primary way to create loops
func main(){

	//1
	for i:=0;i<5;i++{
		fmt.Println("Number is :",i)
	}

	//2
	//to make infinite loop
	counter:=0
	for {
		fmt.Println("Infinte loop")
		counter++
		if(counter==3){
			break
		}
	}

	//3
	//range keyword - simplifies looping over elements of a collection like slices,arrays,maps,and string.
	//simple word range keyword looping ko asan bana deta hey
	numbers:=[]int{1,2,3,4,5}
	for index,value:=range numbers{
		fmt.Printf("index of numbers is %d,and value is %d\n",index,value)
	}

	data:="hello world!"
	for index,char:=range data{
		fmt.Printf("index of data is %d,index of value is %c\n",index,char)
	}
}