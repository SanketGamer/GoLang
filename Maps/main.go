package main

import "fmt"

//in Go map is a data structure that provides an unordered collection of key-value pairs,where each key must be unique.
//it is similar to dictionary or hashmap in other programming language
//maps are used to associate values with keys and allow for effcient retrieval pf values based on those keys.
func main(){
   
	//name <-> grade
	//make is just a function to use map,slice,channels
	strudentGrade:=make(map[string]int)

	strudentGrade["Prince"]=34
	strudentGrade["Sanket"]=16
	strudentGrade["Rahul"]=48
	strudentGrade["Sani"]=52

	fmt.Println("Marks of Sani : ",strudentGrade["Sani"])
	strudentGrade["Sanket"]=99
	fmt.Println("Marks of Sanket : ",strudentGrade["Sanket"])
	//delete take two argument map name,key
	//delete(strudentGrade,"Sani")
	//fmt.Println("Marks of Sani : ",strudentGrade["Sani"])

	//Checking if key exists
	grades,exists:=strudentGrade["Rahul"]
	fmt.Println("Grade is ",grades)
	fmt.Println("Rahul exist",exists)

	Grades,Exists:=strudentGrade["Rahul"]
	fmt.Println("Grade is ",Grades)
	fmt.Println("Rahul exist",Exists)

	for index,value:=range strudentGrade{
		fmt.Printf("Key is %s and marks is %d\n",index,value)
	}

	person:=map[string]int{
		"s":10,
		"r":20,
		"c":30,
	}

	for index,value:=range person{
		fmt.Println(index,value)
	}
}