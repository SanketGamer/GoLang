package main

import (
	"fmt"
	"strings"
)

func main(){
	//i want separate of every string
	data:="apple,orange,banana"
	parts:=strings.Split(data,",")
	fmt.Println(parts)

	//i want how many times two occured
	str:="one two three four two two five"
	count:=strings.Count(str,"two")
	fmt.Println(count)

	//i want get rid all the space
	str="   Hello, Go!  "
	trimmed:=strings.TrimSpace(str) //but it only get rid the space before hello and after GO! 
	fmt.Println(trimmed)

	//i want to join two strings
	str1:="Prince"
	str2:="Agarwal"
	res:=strings.Join([]string{str1,str2}, " ")
	fmt.Println(res)
}