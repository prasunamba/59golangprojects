package main

import (
	"context"
	"fmt"
	"time"
)

func senddata(c context.Context) {
	for {
		select {
		case <-c.Done():
			fmt.Println("crossed deadline ")
			return
		default:
			fmt.Println("working ")
			time.Sleep(500 * time.Millisecond)
		}

	}
}

/* func main() {
	ctxt := context.Background()
	deadline := time.Now().Add(2 * time.Second)
	c, _ := context.WithDeadline(ctxt, deadline)
	go senddata(c)
	time.Sleep(3 * time.Second)

} */
