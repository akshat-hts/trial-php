package main

import (
	"fmt"
	"sync"
)
/*
Ping: 1
Pong: 2
Ping: 3
Pong: 4
*/

func ping(pi, po chan int, wg * sync.WaitGroup) {
	defer wg.Done()
	for {
		num := <-pi
		fmt.Println("Ping: ", num)
		// sender <- receiver
		if (num == 19) {
			po <- num + 1 
			break;
		} else if (num < 20) {
			po <- num + 1
		} 
	}
}

func pong(pi, po chan int, wg * sync.WaitGroup) {
	defer wg.Done()
	for {
		num := <-po
		fmt.Println("Pong: ", num)
		num++
		if (num < 20) {
			pi <- num
		} else {
			break; 
		}
	}
}

func main() {
	pi := make(chan int)
	po := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	
	go ping(pi, po, &wg)
	go pong(pi, po, &wg)

	pi <- 1

	wg.Wait()

		
}