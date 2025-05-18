package main

import (
	"fmt"
	"math/rand"
	"time"
)

func numsGenerator(numCh chan int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
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
	for res := range sqCh {
		fmt.Println(res)
	}
}
