package main
import (
	"fmt"
	"sync"
)
func producer(chn chan int, wg * sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		chn <- i
	}
	close(chn)
}

func consumer(chn chan int, wg * sync.WaitGroup) {
	defer wg.Done()
	// for i := 1; i <= 10; i++ {
	// 	temp := <-chn
	// 	num := temp * temp
	// 	fmt.Println(num)
	// }	// assumes 10 better soln below
	for num := range chn {
		fmt.Println(num * num)
	}
}
// for num := range chn => keeps receiving values until CHANNEL IS CLOSED
func main() {
	chn := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)

	go producer(chn, &wg)
	go consumer(chn, &wg)

	wg.Wait()
	fmt.Println("gg")
}