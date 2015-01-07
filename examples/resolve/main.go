package main

import (
	"fmt"
	"github.com/sycamoreone/orc/control"
	"log"
	"strings"
	"time"
)

func main() {
	c, err := control.Dial(":9051")
	if err != nil {
		log.Fatalln(err)
	}
	err = c.Auth("supersecretpasswd")
	if err != nil {
		log.Fatalln(err)
	}

	errChan := make(chan error)

	// Call ReceiveToChan in a loop.
	go func() {
		for {
			err := c.ReceiveToChan()
			if err != nil {
				errChan <- err
			}
		}
	}()

	c.SetEvents([]string{"ADDRMAP"}) // Request asyncronous ADDRMAP events,
	c.Resolve("torproject.org")      // and resolve torproject.org.

	timer := time.Tick(10 * time.Second)
	for {
		select {
		case <-timer:
			return
		case r := <-c.AsyncReplies:
			ip := strings.SplitN(r.Text, " ", 4)[2]
			fmt.Printf("torproject.org resolved to IP  address %s\n", ip)
			return
		case r := <-c.SyncReplies:
			if r.Status != 250 || r.Text != "OK" {
				log.Print("unexpected sync reply: ", r)
			}
		case err := <-errChan:
			log.Print(err)
		}
	}
}
