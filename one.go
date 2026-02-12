package main

import (
	"fmt"
	"sync"
)
func lessThan (ch chan bool, wg * sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println("One: ", i)
	}
	ch <- true
	
}
func greaterThan(ch chan bool, wg * sync.WaitGroup) {
	defer wg.Done()
	<-ch
	for i := 6; i <= 10; i++ {
		fmt.Println("Two: ", i)
	}
}
func main() {
	ch := make(chan bool)

	var wg sync.WaitGroup 	// var wg of type sync.WaitGroup
							// eqv of wg := sync.WaitGroup
	wg.Add(2)

	go lessThan(ch, &wg)
	go greaterThan(ch, &wg)

	wg.Wait();

	fmt.Println("EOP")
}