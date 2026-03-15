package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("X is greater than 5")
	}	else{
		fmt.Println("NO")
	}
	z:=12
	if z>10{
		fmt.Println("z is grater than 10")
	}else if z>11{
		fmt.Println("Z is grater than 11 but ")
	}else{
		fmt.Println("Sorry")
	}

	y:=11
	if x>5 && y>5 && z>5{
		fmt.Println(x,y,z)
	}else{
		fmt.Println("no worries")
	}
}