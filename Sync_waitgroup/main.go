package main

import (
	"fmt"
	"sync"
)

//We use WaitGroup to make sure all goroutines complete before program exits
// WaitGroup = synchronization tool to wait for goroutines

func worker(i int,wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("worker %d started\n ", i)
	//some task is happend
	fmt.Printf("worker %d end\n", i)
}

func main() {

	fmt.Println("Explore goroutine started")

	//now wg matlab konse goroutines checklist mee add huye
	var wg sync.WaitGroup
	//st 3 worker goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i,&wg)
	}
	//wait for all worker
	wg.Wait()
	fmt.Println("Workers task completed")
}
