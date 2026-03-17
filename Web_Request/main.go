package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//Go lang provide a pacakage net/http it helps web request connectivity

func main() {

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting get response from", err)
		return
	}
	//we stop the response body
	defer res.Body.Close()
	fmt.Printf("Type of response is %T\n", res)
	//fmt.Println("response: ",res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while reading response")
		return
	}
	fmt.Println("response is:", string(data))
}
