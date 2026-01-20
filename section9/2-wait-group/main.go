package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Printf("sayHello %s\n", message)
}

func main() {
	/*
	 1. Add outside of the goroutine
	 2. You must decrease the counter by calling wg.Done inside the goroutine
	 3. Do not forget to call wg.Wait()
	 4. Always pass a reference/pointer of the wait group variable instead of a copy
	*/

	var wg sync.WaitGroup

	fmt.Println("Mai() Go Routine")

	totalJobs := 5
	for i := 0; i < totalJobs; i++ {
		wg.Add(1)
		go sayHello(fmt.Sprintf("Job %d", i), time.Second, &wg)
	}

	// go sayHello("first", time.Second, &wg)
	// go sayHello("second", time.Second, &wg)
	// go sayHello("third", 2*time.Second, &wg)
	// go sayHello("fourth", 3*time.Second, &wg)
	// go sayHello("fifth", time.Second, &wg)

	fmt.Println("Last Message from Main()")

	wg.Wait()
}
