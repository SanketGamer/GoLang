package main

import (
	"fmt"
	"net/url" //this pacakge we used for extract the url specific things
)

func main() {

	//schema is my https
	//host - "example.com:8000"
	//Path - "/path/to/resource"
	//RawQuery - "key1=value1&key2=value2"
	Myurl := "https://example.com:8000/path/to/resource?key1=value1&key2=value2"
	//exurl:="sanketdas1348"
	fmt.Printf("Type of %T\n",Myurl)

	parseUrl,err:=url.Parse(Myurl)
	if err!=nil{
		fmt.Println("Can not Parse url",err)
		return
	}
	fmt.Printf("Type of Url is: %T\n",parseUrl)
	fmt.Println("Scheme: ",parseUrl.Scheme)
	fmt.Println("Host: ",parseUrl.Host)
	fmt.Println("Path: ",parseUrl.Path)
	fmt.Println("RawQuery: ",parseUrl.RawQuery)

	//modifying url components
	parseUrl.Path="/newPath"
	parseUrl.RawQuery="username=sanket"

	//Constructing a Url string from url object
	newUrl:=parseUrl.String()
	fmt.Println("new url: ",newUrl)
}