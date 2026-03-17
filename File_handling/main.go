package main

import (
	"fmt"
	"io/ioutil"
	//"os"
)

// file handling -
func main() {
	//os basically a pacakge to create,delete,update,read like that
	// file,err:=os.Create("example.txt")  //this means create a file example.txt
	// if err!=nil{
	// 	fmt.Println("Error while creating file")
	// 	return
	// }
	// defer file.Close() //Close means the file resourses is now free at the end
	// content:="Hello World"
	// _,error:=io.WriteString(file,content)  //we can add the content using this io.WriteString(file,content) this return 2 things byte,errors so we can ignore byte
	// if error!=nil{
	// 	fmt.Println("Error while writing file: ",error)
	// }
	// fmt.Println("Successfully created file")

	//Reading file Buffer
	// file, err := os.Open("example.txt")
	// if err != nil {
	// 	fmt.Println("Error while creating file")
	// 	return
	// }
	// defer file.Close()
	// //create a buffer to read file content
	// buffer:=make([]byte,1024)

	// //Read file content into buffer
	// for{
	// 	n,err:=file.Read(buffer)
	// 	if err==io.EOF{  //end of file
	// 		break
	// 	}
	// 	//fmt.Println(n)
	// 	if err!=nil{
	// 		fmt.Println("Error while reading file",err)
	// 		return
	// 	}
	// 	//Process the read content
	// 	fmt.Println(string(buffer[:n]))  //buffer me jo data a raha wo string me ayega aur print 
	// }


	//Read the entire content file into byte slice
	content, err := ioutil.ReadFile("example.txt")
	if err!=nil{
		fmt.Println("Err while reading the file",err)
		return
	}
	fmt.Println(string(content))

}
