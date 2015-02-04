package main

import (
	"log"
	"time"
	"strings"

	"github.com/sycamoreone/orc/tor"
	"github.com/sycamoreone/orc/control"
)

func main() {
	log.SetFlags(log.Ldate)

	// Start a new tor process.
	config := tor.NewConfig()
	config.Timeout = 20 * time.Second

	config.Set("SocksPort", 9055)
	config.Set("ControlPort", 9051)
	config.Set("HashedControlPassword", "16:97E10A1BE81C795E600BBD9C4A912BF853FC213F867A6C07869108EFBB")
	err := config.Err()
	if err != nil {
		log.Fatal("resolve: error setting configuration parameters", err)
	}

	cmd, err := tor.NewCmd(config)
	if err != nil {
		log.Fatal(err)
	}
	defer cmd.KillUnlessExited()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("resolve: Tor 100% bootstrapped")

	// Connect to the control port.
	c, err := control.Dial(":9051")
	if err != nil {
		log.Fatal("resolve: ", err)
	}

	err = c.Auth("supersecretpasswd")
	if err != nil {
		log.Fatal("resolve:", err)
	}

	log.Print("resolve: Connected to the ControlPort")

	c.SetEvents([]string{"ADDRMAP"}) // Request asyncronous ADDRMAP events,
	c.Resolve("torproject.org")      // and resolve torproject.org.

	returnChan := make(chan bool)
	c.Handle("ADDRMAP", func(r *control.Reply) {
		ip := strings.SplitN(r.Text, " ", 4)[2]
		log.Printf("resolve: torproject.org resolved to IP  address %s\n", ip)
		returnChan <- true
	})

	go c.ReceiveToChan()
	go c.Serve()
	<-returnChan

	log.Print("resolve: Shutting down the tor proxy now.")

	err = c.Signal(control.SignalShutdown)
	if err != nil {
		log.Fatal("resolve: ", err)
	}
	cmd.Wait()
	return
}
