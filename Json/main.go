package main

import (
	"encoding/json"
	"fmt"
)

//this is a object
type Person struct {
	Name    string// `json:"name"`  this called we used lowercase json keys
	Age     int    //`json:"age"`
	isAdult bool   //`json:"isAdult"`
}

func main() {

	person := Person{Name: "John", Age: 34, isAdult: true}
	fmt.Println("person Data is :",person)
	fmt.Printf("%T\n",person)

	//covert Person into Json Encoding(Marshalling)
	//Marshal expect - object
	//Marshal(encoding) → Go struct(object) → JSON string
	jsonData,err:=json.Marshal(person)
	if err!=nil{
		fmt.Println("Error Marshaling",err)
		return
	}
	fmt.Println("Person data is",string(jsonData))

	//UnMarshal(decoding) : JSON → Go object/struct
	var personData Person
	err=json.Unmarshal(jsonData, &personData)
	if err!=nil{
		fmt.Println("Error unmarshaling",err)
		return
	}
	fmt.Println("Person data is: ",personData)
}