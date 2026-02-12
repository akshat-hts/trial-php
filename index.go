package main

import (
	"fmt"
	// "time"
	"sync"
)

func printOdd(oddCh chan bool, evenCh chan bool, wg * sync.WaitGroup) {
	defer wg.Done()			// defer exec after func is done, cnt is reduced by 1
	for i := 1; i <= 10; i += 2 {
		<-oddCh
		fmt.Println("Odd: ", i)
		evenCh <- true
	}
}
func printEven(oddCh chan bool, evenCh chan bool, wg * sync.WaitGroup) {
	defer wg.Done()			// defer exec after func is done, cnt -= 1
	for i := 2; i <= 10; i += 2 {
		<-evenCh		// after 5th itr, this fnc blocks forever
		fmt.Println("Even: ", i)
		if (i < 10) {
			oddCh <- true
		}
	}
}
// func deferExample() {
// 	fmt.Println("Start")
// 	defer fmt.Println("End")
// 	fmt.Println("Middle");
// }
func main() {
	oddCh := make(chan bool)
	evenCh := make(chan bool)
	
	var wg sync.WaitGroup	// counter = 0
	wg.Add(2)				// counter += 2; thus cnt = 2

	go printOdd(oddCh, evenCh, &wg)
	go printEven(oddCh, evenCh, &wg)
	
	oddCh <- true	// run above fnc first, then send signal

	wg.Wait()	// waits till cnt = 0
	// deferExample(); 
}