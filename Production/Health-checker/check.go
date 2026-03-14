
package main

//the net package in Go is used to check network connectivity or status, but more precisely:
import (
	"fmt"
	"time"
	"net"
)

func Check(destination string,port string) string{
	address:=destination+":"+port
	timeout:=time.Duration(2*time.Second)
	conn,err:=net.DialTimeout("tcp",address,timeout)
	var status string

	if err !=nil{
		status=fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v",destination,err)
	} else{
		//RemoteAddr basically website port mention
		status=fmt.Sprintf("[UP] %v is reachable,\n From:%v\n To: %v",destination,conn.LocalAddr(),conn.RemoteAddr())
	}
	return status
}