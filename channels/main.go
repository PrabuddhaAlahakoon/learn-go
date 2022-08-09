package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	//ch := make(chan int)

	// for j := 0; j < 5; j++ {
	// 	wg.Add(2)

	// 	go func() {
	// 		i := <-ch
	// 		fmt.Println(i)
	// 		wg.Done()
	// 	}()

	// 	go func() {
	// 		ch <- 42
	// 		wg.Done()
	// 	}()

	// }

	//part 2

	//ch := make(chan int)
	// wg.Add(2)

	// go func(ch <-chan int) {
	// 	i := <-ch
	// 	fmt.Println(i)
	// 	wg.Done()
	// }(ch)

	// go func(ch chan<- int) {
	// 	ch <- 42
	// 	wg.Done()
	// }(ch)

	//part 3

	ch := make(chan int, 50)
	wg.Add(2)
	go func(ch <-chan int) {
		// for i := range ch {
		// 	fmt.Println(i)
		// }

		//same thing as commented code above
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
