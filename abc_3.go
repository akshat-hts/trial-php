package main

import (
	"fmt"
	"sync"
)
func printA(chA, chC chan bool, wg * sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-chC
		fmt.Println("A")
		chA <- true
	}
}
func printB(chB, chA chan bool, wg * sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-chA
		fmt.Println("B")
		chB <- true
	}
}
func printC(chC, chB chan bool, wg * sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		<-chB
		fmt.Println("C")
		if (i < 4) {
			chC <- true
		}
	}
}

func main() {
	chA := make(chan bool)
	chB := make(chan bool)
	chC := make(chan bool)

	var wg sync.WaitGroup
	wg.Add(3)

	go printA(chA, chC, &wg)
	go printB(chB, chA, &wg)
	go printC(chC, chB, &wg)

	chC <- true

	wg.Wait()

	close(chA)
	close(chB)
	close(chC)

}