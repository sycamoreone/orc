package main

import (
	"log"
	"os"
	"time"

	"github.com/sycamoreone/orc/tor"
	//"github.com/sycamoreone/orc/control"
)

func main() {
	config := tor.NewConfig()
	config.Set("CookieAuthentication", 1)
	config.Set("SocksPort", 9055)
	config.Set("ControlPort", 9051)
	err := config.Err()
	if err != nil {
		log.Fatal("error setting configuration parameters")
	}

	cmd, _ := tor.NewCmd("", config)
	cmd.Cmd.Stdout = os.Stdout // DEBUG
	err = cmd.Start()
	if err != nil {
		return
	}

	time.Sleep(10 * time.Second)
	cmd.Cmd.Process.Signal(os.Interrupt)
	cmd.Wait()
	return
}
