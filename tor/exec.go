// Package process supplies helper functions to start a tor binary as a slave process.
package tor

import (
	"os/exec"
	//	"bufio"
	//	"regexp"
)

// Cmd represents an tor executable to be run as a slave process.
type Cmd struct {
	Config *Config
	Cmd    *exec.Cmd // TODO: We probably shouldn't expose the exec.Cmd
}

// NewCmd returns a Cmd to run a tor process using the configuration values in config.
// The argument path is the path to the tor program to be run. If path is the empty string,
// $PATH is used to search for a tor executable.
func NewCmd(path string, config *Config) (*Cmd, error) {
	if path == "" {
		file, err := exec.LookPath("tor")
		if err != nil {
			return nil, err
		}
		path = file
	}
	return &Cmd{Config: config, Cmd: exec.Command(path, config.ToCmdLineFormat()...)}, nil
}

func (c *Cmd) Start() error {
	err := c.Cmd.Start()
	if err != nil {
		return err
	}
	// TODO: read output until one gets a "Bootstrapped 100%: Done" notice.
	return nil
}

func (c *Cmd) Wait() error {
	return c.Cmd.Wait()
}
