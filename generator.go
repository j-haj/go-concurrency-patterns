package main

import (
	"fmt"
	"time"
)

func GenerateFilter(filterVal int, source <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		encounteredFirstVal := false
		for {
			select {
			case v := <-source:
				if v == filterVal {
					if !encounteredFirstVal {
						encounteredFirstVal = true
						output <- v
					}
				} else {
					output <- v
				}
			}
		}
	}()

	return output
}

func SimpleGenerator(str string, delay time.Duration) <-chan string {
	output := make(chan string)
	go func() {
		for {
			select {
			case <-time.After(delay * time.Millisecond):
				output <- str
			}
		}
	}()
	return output
}

func main() {
	ch := SimpleGenerator("Hello", time.Duration(1000))
	for {
		select {
		case v := <-ch:
			fmt.Printf("Received: %s\n", v)
		}
	}
}
