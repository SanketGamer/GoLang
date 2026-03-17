package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"net/url"
	"strings"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func GetReq() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")
	if err != nil {
		fmt.Println("Error getting: ", err)
		return
	}
	defer res.Body.Close() //program may leak resources
	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting responce: ", res.Status)
	}
	// this is a generic responce this is not a good way to write
	// //this ReadAll returns byte of data
	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading : ", err)
	// 	return
	// }
	// fmt.Println("Data : ",string(data))

	//best way
	var todo Todo
	//jo bhi data rahega ussko decode karega normal objects me aur store karega todo variable
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error Decoding", err)
		return
	}
	fmt.Println("Todo: ", todo)

	fmt.Println("Title ", todo.Title)
	fmt.Println("completed responce: ", todo.Completed)
}
func PostReq() {
	todo := Todo{
		UserId:    23,
		Title:     "rahul",
		Completed: true,
	}
	//Convert the Todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling : ", err)
		return
	}
	//convert json data to string
	jsonString := string(jsonData)
	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	//post take 3 argument - url,contentType,body(ek tare ka reader hota hey)

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	res, err := http.Post(myUrl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error sending request : ", err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Responce : ", string(data))
	fmt.Println("Responce status: ", res.Status)
}
func UpdateReq() {
	todo := Todo{
		UserId:    2222,
		Title:     "SanketGamer in Golang",
		Completed: false,
	}
	//convert todo struct to JSON
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	const myurl = "https://jsonplaceholder.typicode.com/todos/1"
	//convert json data to string
	jsonString := string(jsonData)
	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	//create put req
	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error creating Put request : ", err)
		return
	}
	req.Header.Set("Content-type", "application/json")
	//send req
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending req", err)
		return
	}
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Responce : ", string(data))

}
func DeleteReq(){
  const myurl="https://jsonplaceholder.typicode.com/todos/1"

  req,err:=http.NewRequest(http.MethodDelete,myurl,nil)
  if err!=nil{
	fmt.Println("Error creating  Delete req: ",err)
	return
  }
  //send the req
  client:=http.Client{}
  res,err:=client.Do(req)
  if err!=nil{
	fmt.Println("Error sending req: ",err)
	return
  }
  defer res.Body.Close()

  data,_:=ioutil.ReadAll(res.Body)
  fmt.Println("Responce: ",string(data))
  fmt.Println("Responce status :",res.Status)
}
func main() {
	//GetReq()
	//PostReq()
	//UpdateReq()
	DeleteReq()
}
