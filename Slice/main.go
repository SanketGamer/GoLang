package main

import "fmt"

func main() {
	 numbers := []int{1, 2, 3}
	// numbers=append(numbers, 6,7,8)
	// fmt.Println(numbers)
	// fmt.Printf("number has data type:  %T\n",numbers)
	// name:=[]string{}
	// fmt.Println(name)

	 fmt.Println("slice: ",numbers)
	 fmt.Println("Length: ",len(numbers))
	 fmt.Println("Capacity: ",cap(numbers))

	 //make function
	 number:=make([]int,3,5)
	 number=append(number,4)
	 number=append(number,6)
	 number=append(number,8)
	 fmt.Println("slice: ",number)
	 fmt.Println("Length: ",len(number))
	 fmt.Println("Capacity: ",cap(number))

	 //we can not inisialize slice like this we need to use make function
	 //var stock=[]string
	 stock:=make([]int,0)
	 fmt.Println(stock)


}