package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := 2 * time.Second

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return fmt.Sprintf("[DOWN] %v unreachable: %v", address, err)
	}

	defer conn.Close()

	return fmt.Sprintf("[UP] %v reachable", address)
}