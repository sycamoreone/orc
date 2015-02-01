package main

import (
	"github.com/sycamoreone/orc/control"
	"log"
)

func main() {
	log.SetFlags(log.Ldate)
	c, err := control.Dial(":9051")
	if err != nil {
		log.Print("shutdown: ", err)
	}
	err = c.Auth("supersecretpasswd")
	if err != nil {
		log.Print("shutdown: ", err)
		return
	}
	log.Print("shutdown: connected; shutting down the Tor router")

	err = c.Signal(control.SignalShutdown)
	if err != nil {
		log.Print("shutdown: ", err)
		return
	}
}
