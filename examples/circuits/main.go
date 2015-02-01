package main

import (
	"github.com/sycamoreone/orc/control"
	"log"
)

func main() {
	log.SetFlags(log.Ldate)
	c, err := control.Dial(":9051")
	if err != nil {
		log.Print("circuits: ", err)
	}
	err = c.Auth("supersecretpasswd")
	if err != nil {
		log.Print("circuits", err)
		return
	}
	log.Print("circuits: connected to a Tor router.")

	reply, err := c.GetInfo("circuit-status")
	if err != nil {
		log.Print("circuits: ", err)
		return
	}
	log.Print(reply)
}
