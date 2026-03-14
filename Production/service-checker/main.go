package main

import (
	"fmt"
	"net"
	"time"
)

// This function checks if the service is running
func isRunning(name, address string) {
	timeout := 2 * time.Second
	conn, err := net.DialTimeout("tcp", address, timeout)
	
	if err != nil {
		fmt.Printf("%s: false\n", name)
		return
	}
	
	conn.Close()
	fmt.Printf("%s: true\n", name)
}

func main() {
	// IMPORTANT: Change "localhost" to your actual Server IPs if they are remote
	services := map[string]string{
		"RabbitMQ":  "localhost:5672",
		"Redis":     "localhost:6379",
		"Mosquitto": "localhost:1883",
	}

	fmt.Println("--- Service Status ---")
	for name, addr := range services {
		isRunning(name, addr)
	}
}