package ex05

import (
	"fmt"
	"os"
)

func Exec(numOfSend int) {
	msg1 := make(chan int)
	msg2 := make(chan int)
	go func() {
		for {
			select {
			case x := <-msg1:
				fmt.Printf("channel1: %d\n", x)
				msg2 <- x + 1
			default:
				// do nothing
			}
		}
	}()

	msg1 <- 0
	for {
		select {
		case x := <-msg2:
			if x > numOfSend {
				os.Exit(0)
			}
			fmt.Printf("channel2: %d\n", x)
			msg1 <- x + 1
		default:
			// do nothing
		}
	}
}
