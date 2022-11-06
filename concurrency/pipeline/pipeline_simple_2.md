# Pipeline Pattern

## Description
In progress ...

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var workers int = 3

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
	var wg sync.WaitGroup
	processTwo := func() {
		for num := range inCh {
			// simulate hard work
			time.Sleep(1 * time.Second)
			outCh <- num + 2
		}
	}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			processTwo()
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	return outCh
}

// step 3
func sumThree(inCh <-chan int) <-chan int {
	outCh := make(chan int)
	go func() {
		for num := range inCh {
			outCh <- num + 3
		}
		close(outCh)
	}()
	return outCh
}

func main() {
	for num := range sumThree(sumTwo(generateNumber(1, 2, 3, 4, 5, 6, 7, 8))) {
		fmt.Println("Printing...", num)
	}
}

```