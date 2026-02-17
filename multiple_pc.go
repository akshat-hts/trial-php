package main

import (
	"fmt"
	"sync"	
)
func producerA(chA, chB chan bool, wg * sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		<-chA
		fmt.Println("Producer A: ", i)
		chB <- true
	}
	close(chA)
}
func producerB(chB, chC chan bool, wg * sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		<-chB
		fmt.Println("Producer B: ", i)
		chC <- true
	}
	close(chB)

}
func producerC(chC, chA chan bool, wg * sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		<-chC
		fmt.Println("Producer C: ", i)
		fmt.Println()
		if (i < 5) {
			chA <- true
		}
	}
	close(chC)

}
func main() {
	fmt.Println()
	var wg sync.WaitGroup
	chA, chB, chC := make(chan bool), make(chan bool), make(chan bool)
	wg.Add(3)

	go producerA(chA, chB, &wg)
	go producerB(chB, chC, &wg)
	go producerC(chC, chA, &wg)

	chA <- true

	wg.Wait()
}