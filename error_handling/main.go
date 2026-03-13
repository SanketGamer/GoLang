package main

import "fmt"


// func divide(a,b float64) (float64,error){
// 	if b==0{
// 		return 0,fmt.Errorf("denominator not must be zero")
// 	}
// 	return a/b,nil
// }

func divide(a,b float64) (float64,string){ //this is string use as a err handling
	if b==0{
		return 0,"denominator not must be zero"
	}
	return a/b,"nil"
}
func main() {
	fmt.Println("Started Error handling in the functions")
	ans,_:= divide(10,2)  //underscore basically serves as a write-only variable that allows you to discard values returned by a function to ignore specific values(matlab hamm err variable k sath kuch khelna nhi cha rehe hey blacklist wale me dall diya means underscore me dall diya)
	fmt.Println("Division of two number",ans)
}