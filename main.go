package main

import (
	"fmt"
)

// step 1
func generateNumber(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, num := range nums {
			ch <- num
		}
		close(ch)
	}()
	return ch
}

// step 2
func sumTwo(inCh <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		for num := range inCh {
			outCh <- num + 2
		}
		close(outCh)
	}()
	return outCh
}

func main() {
	for num := range sumTwo(generateNumber(1, 2, 3)) {
		fmt.Println("Printing...", num)
	}
}
