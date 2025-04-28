package main

import (
	"context"
	"fmt"
	"time"
)

func sendvalue(c context.Context) {
	for {
		select {
		case <-c.Done():
			fmt.Println("crossed deadline ")
			return
		default:
			fmt.Println("working ", c.Value("userid"))
			time.Sleep(500 * time.Millisecond)
		}

	}
}

/* func main() {
	ctxt := context.Background()
	c := context.WithValue(ctxt, "userid", "12")
	go sendvalue(c)
	time.Sleep(1000 * time.Millisecond)

} */
