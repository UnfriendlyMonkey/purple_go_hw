package main

import (
	"fmt"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func numsGenerator(numCh chan int) {
	for i := 0; i < 10; i++ {
		numCh <- r.Intn(101)
	}
	close(numCh)
}

func squareNums(numCh chan int, sqCh chan int) {
	for n := range numCh {
		res := n * n
		sqCh <- res
	}
	close(sqCh)
}

func main() {
	numCh := make(chan int)
	sqCh := make(chan int)

	go numsGenerator(numCh)
	go squareNums(numCh, sqCh)

	// Add timeout to prevent infinite waiting
	timeout := time.After(2 * time.Second)

	for {
		select {
		case res, ok := <-sqCh:
			if !ok {
				return
			}
			fmt.Println(res)
		case <-timeout:
			fmt.Println("Operation timed out")
			return
		}
	}
}
