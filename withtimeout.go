package main

import (
	"context"
	"fmt"
	"time"
)

func Senddata(c context.Context) {
	for {
		select {
		case <-c.Done():
			fmt.Println("context timeout ")
			return
		default:
			fmt.Println("working ")
			time.Sleep(500 * time.Millisecond)
		}

	}
}

/* func main() {
	ctxt := context.Background()
	c, _ := context.WithTimeout(ctxt, 3*time.Second)
	go senddata(c)
	time.Sleep(4 * time.Second)
} */
