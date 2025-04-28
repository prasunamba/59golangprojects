package main

import (
	"context"
	"fmt"
	"time"
)

/*
	 func main() {

		ch1 := make(chan string)
		ch2 := make(chan string)

		go func() {
			time.Sleep(2 * time.Second)
			ch1 <- "Message from Channel 1"
			close(ch1)
		}()

		go func() {
			time.Sleep(1 * time.Second)
			ch2 <- "Message from Channel 2"
			close(ch2)
		}()
		// We use a flag to track if both channels are closed
		ch1Closed, ch2Closed := false, false

		for {
			select {
			case msg1, ok := <-ch1:
				if !ok {
					// ch1 is closed
					fmt.Println("ch1 is closed")
					ch1Closed = true
				} else {
					fmt.Println("msg1:", msg1)
				}
			case msg2, ok := <-ch2:
				if !ok {
					// ch2 is closed
					time.Sleep(1 * time.Second)

					fmt.Println("ch2 is closed")
					ch2Closed = true
				} else {
					fmt.Println("msg2:", msg2)
				}
			}

			// Exit condition: both channels are closed
			if ch1Closed && ch2Closed {
				break
			}
		}

}
*/

func send(c context.Context) {
	for {
		select {
		case <-c.Done():
			fmt.Println("context cancelled ")
			return
		default:
			fmt.Println("working ")
			time.Sleep(1 * time.Second)
		}

	}

}
func main() {

	ctxt := context.Background()
	c, cancel := context.WithCancel(ctxt)
	go send(c)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)

}
