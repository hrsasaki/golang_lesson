package ex04

import (
	"fmt"
	"sync"
)

var channels []chan int
var num = 1

var initOnce sync.Once

func initChannels() {
	for i := 0; i < num; i++ {
		channels = append(channels, make(chan int))
	}
}

func Exec(numOfGoroutines int) {
	num = numOfGoroutines
	initChannels()

	// first goroutine of pipeline
	go func() {
		for i := 0; i < 5; i++ {
			channels[0] <- i
		}
		close(channels[0])
	}()

	for i := 1; i < num; i++ {
		// copy index variable
		index := i
		go func() {
			for x := range channels[index-1] {
				channels[index] <- x
			}
			close(channels[index])
		}()
	}

	// last goroutine of pipeline
	for x := range channels[len(channels)-1] {
		fmt.Println(x)
	}
}
