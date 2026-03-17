package main

import "fmt"


func add(a,b int)int{
	return a+b
}
func main() {
	//now i want middle execute last
	//defer means when all the line executes and last one will print
	//when 2 defer executes it use like stack(lifo)
	fmt.Println("Starting of the Program")
	data:=add(5,6)
	defer fmt.Println(data)
	defer fmt.Println("Middle of the Program")
	defer fmt.Println("End of the Program")

}


