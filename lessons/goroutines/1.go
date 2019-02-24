package goroutines

import (
	"fmt"
	"time"
)

// Go's _select_ lets you wait on multiple channel
// operations. Combining goroutines and channels with
// select is a powerful feature of Go.

func goRoutinesExample() {

	// For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		c1 <- "one"
		close(c1)
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "two"
	}()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.

	timeout := time.After(10 * time.Second)

Loop:
	for {
		fmt.Println("uruchamiam pętlę")

		channelFirstClosed := false
		channelSecondClosed := false

		select {
		case msg2, ok := <-c2:
			if ok {
				fmt.Printf("received %v with status: %v \n", msg2, ok)
			} else {
				fmt.Printf("channel: %v closed with status: %v \n", "2", ok)
				channelFirstClosed = true
			}

		default:
			fmt.Println("brak wiadomosci na kanale 2")
		}

		select {
		case msg1, ok := <-c1:
			if ok {
				fmt.Printf("received %v with status: %v \n", msg1, ok)
			} else {
				fmt.Printf("channel: %v closed with status: %v \n", "1", ok)
				channelSecondClosed = true
			}
		default:
			fmt.Println("brak wiadomosci na kanale 1")
		}

		time.Sleep(time.Second * 2)

		if channelSecondClosed && channelFirstClosed {
			fmt.Println("Wszystkie kanały zamknięte, kończę")
			break Loop
		}

		select {
		case <-timeout:
			fmt.Println("To wszystko za długo trwa, przerywam....")
			break Loop
		default:
			fmt.Println("spróbuję jeszcze raz")
		}
	}
}
