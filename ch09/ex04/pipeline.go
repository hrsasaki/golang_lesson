package ex04

import "fmt"

func Exec(numOfGoroutines int) {
	var channels []chan int
	for i := 0; i < numOfGoroutines; i++ {
		channels = append(channels, make(chan int))
	}

	// first goroutine of pipeline
	go func() {
		for i := 0; i < 5; i++ {
			channels[0] <- i
		}
		close(channels[0])
	}()

	for i := 1; i < numOfGoroutines; i++ {
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
		fmt.Print(x)
	}
}
