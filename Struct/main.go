package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}
type Contact struct {
	email string
	phone string
}
type Address struct {
	HouseId int
	area    string
	state   string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Add     Address
}

func main() {
	// var sanket Person
	// fmt.Println(sanket)
	// sanket.FirstName = "guddu"
	// sanket.LastName = "agarwal"
	// sanket.Age = 22
	// fmt.Println(sanket)

	// person1 := Person{
	// 	FirstName: "Jitu",
	// 	LastName:  "Singh",
	// 	Age:       23,
	// }
	// fmt.Println(person1)

	//new keyword
	//var person2 = new(Person)  //so person2 as a pointer pointing to Person
	//person2.FirstName = "Sani"
	//person2.LastName = "Das"
	//person2.Age = 34
	//fmt.Println(person2)

	var employee Employee
	employee.Person_Details = Person{
		FirstName: "ssr",
		LastName:  "das",
		Age:       26,
	}
	employee.Person_Contact.email = "contact@gmail.com"
	employee.Person_Contact.phone = "123"

	employee.Person_Add = Address{
		HouseId: 4356,
		area:    "Near",
		state:   "himachal",
	}
	fmt.Println(employee)
}
