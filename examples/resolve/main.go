package main

import (
	"fmt"
	"github.com/sycamoreone/orc/control"
	"log"
	"strings"
	"os"
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

	c.SetEvents([]string{"ADDRMAP"}) // Request asyncronous ADDRMAP events,
	c.Resolve("torproject.org")      // and resolve torproject.org.

	c.Handle("ADDRMAP", func (r *control.Reply) {
		ip := strings.SplitN(r.Text, " ", 4)[2]
		fmt.Printf("torproject.org resolved to IP  address %s\n", ip)
		os.Exit(0)
	})

	// TODO: This is ugly. All methods that send commands should make sure
	// to read the synchromous replies. Then this would be necessary.
	go func () {
		for {
			_ = <- c.SyncReplies
		}
	}()

	c.ReceiveAndServe()
}
