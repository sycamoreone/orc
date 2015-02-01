package main

import (
	"github.com/sycamoreone/orc/control"
	"log"
	"strings"
)

func main() {
	log.SetFlags(log.Ldate)
	c, err := control.Dial(":9051")
	if err != nil {
		log.Fatalln("resolve", err)
	}
	err = c.Auth("supersecretpasswd")
	if err != nil {
		log.Fatalln("resolve: ", err)
	}

	c.SetEvents([]string{"ADDRMAP"}) // Request asyncronous ADDRMAP events,
	c.Resolve("torproject.org")      // and resolve torproject.org.

	exitChan := make(chan bool)
	c.Handle("ADDRMAP", func(r *control.Reply) {
		ip := strings.SplitN(r.Text, " ", 4)[2]
		log.Printf("torproject.org resolved to IP  address %s\n", ip)
		exitChan <- true
	})

	go c.ReceiveToChan()
	go c.Serve()
	<-exitChan
}
