package main

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandPeriodicSequence(delay time.Duration, tag string) chan string {
	output := make(chan string)
	go func() {
		for {
			dur := (delay + time.Duration(rand.Int63n(1000))) * time.Millisecond
			<-time.After(dur)
			output <- tag
		}
	}()
	return output
}

func TwoChanFanIn(ch1, ch2 chan string) chan string {
	output := make(chan string)
	go func() {
		for {
			select {
			case v := <-ch1:
				output <- v
			case v := <-ch2:
				output <- v
			}
		}
	}()
	return output
}

func main() {
	ch1 := GenerateRandPeriodicSequence(100, "Jim")
	ch2 := GenerateRandPeriodicSequence(50, "Jane")
	fan := TwoChanFanIn(ch1, ch2)
	for {
		v := <-fan
		fmt.Printf("Received: %s\n", v)
	}
}
