package main

import (
	"fmt"
	"time"
)

func main() {

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
