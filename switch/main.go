package main

import "fmt"

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Unknown day")
	}

	month:="January"
	switch month{
	case "January","Feb":
		fmt.Println("Winter")
	case "mar","april":
		fmt.Println("Spring")
	default:
		fmt.Println("Summer")
	}
}