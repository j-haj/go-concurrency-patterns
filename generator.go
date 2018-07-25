package main

import (
	"fmt"
	"time"
)

func Filter(filterVal int, source <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		encounteredFirstVal := false
		for {
			v := <-source
			if v == filterVal {
				if !encounteredFirstVal {
					encounteredFirstVal = true
					output <- v
				}
			} else {
				output <- v
			}
		}
	}()

	return output
}

// SimpleGenerator sends `str` to the returned string chan every `delay`
// milliseconds.
func SimpleGenerator(str string, delay time.Duration) <-chan string {
	output := make(chan string)
	go func() {
		for {
			<-time.After(delay * time.Millisecond)
			output <- str
		}
	}()
	return output
}

// SequenceGenerator returns an int channel that sends a sequence of ints
// starting at 0 every `delay` milliseconds.
func SequenceGenerator(delay time.Duration) <-chan int {
	i := 0
	output := make(chan int)
	go func() {
		for {
			<-time.After(delay * time.Millisecond)
			output <- i
			i++
		}
	}()
	return output
}

func ModWrapper(m int, input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		for {
			v := <-input
			output <- (v % m)
		}
	}()
	return output
}

func main() {
	m := Filter(0, ModWrapper(3, SequenceGenerator(time.Duration(1000))))
	for {
		select {
		case v := <-m:
			fmt.Printf("Received: %d\n", v)
		}
	}
}
