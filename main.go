package main

import (
	"fmt"
	"time"
)

func printEven() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Even: ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printOdd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd: ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {

	fmt.Println("Hello world")
	go printEven()
	go printOdd()

	time.Sleep(3 * time.Second)

}