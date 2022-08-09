package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

// use go run -race main.go to see if routines are racing with each other
func main() {

	// wg.Add(1)

	// go sayHello()

	// var msg = "Hello2"
	// go func() {
	// 	fmt.Println(msg)
	// }()
	// time.Sleep(100 * time.Microsecond)

	// wg.Add(1)

	// var msg = "Hello2"
	// go func(msg string) {
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(msg)

	//runtime.GOMAXPROCS(100)

	for i := 0; i < 10; i++ {
		wg.Add(2)

		m.RLock()
		go sayHello()

		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello! %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
