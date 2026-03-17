package main

import (
	"fmt"
	"time"
)

func main() {

	//"dd-mm-yyyy"
	//"yyyy-dd-mm"
	//in Go - if u want day->02 , month->01 , year->2006 ,current day->Monday, 24hr formate(current time)->15:04:05 , Am/Pm->2006/01/02 3:04 PM

	currentTime := time.Now()
	fmt.Println(currentTime)
	fmt.Printf("Type of currentTiming: %T\n",currentTime)

	formatted:=currentTime.Format("2006/01/02 3:04 PM")  //formate will give a string as a returned value
	fmt.Println(formatted)

	//parsing method
	layout_string:="2006-01-02"
	dateStr:="2023-11-25"
	formatted_time,_:=time.Parse(layout_string,dateStr)
	fmt.Println(formatted_time)

	//add 1 more day in currentTime
	new_date:=currentTime.Add(24*time.Hour)
	fmt.Println(new_date)
	formatted_new_date:=new_date.Format("2006/01/02 Monday")
	fmt.Println(formatted_new_date)
}