package tor

import (
	"errors"
	"fmt"
	"strconv"
)

// A Config struct is used to configure a to be executed Tor process.
type Config struct {
	// Options is a map of configuration options to values to be used
	// as command line arguments or in a torrc configuration file.
	Options map[string]string
	err     error
}

func NewConfig() *Config {
	c := &Config{Options: make(map[string]string), err: nil}
	return c
}

func (c *Config) setErr(format string, a ...interface{}) {
	err := errors.New(fmt.Sprintf(format, a...))
	if c.err == nil {
		c.err = err
	}
}

func (c *Config) Set(option string, value interface{}) {
	switch v := value.(type) {
	case int:
		c.Options[option] = strconv.Itoa(v)
	case string:
		c.Options[option] = quote(v)
	default:
		c.setErr("value %v for option %s is not a string or int", value, option)
	}
}

func quote(s string) string {
	if s[0] == '"' && s[len(s)-1] == '"' {
		return s
	}
	return "\"" + s + "\""
}

// Err reports the first error that was encountered during the preceding calls to Set()
// and clears the saved error value to nil.
func (c *Config) Err() error {
	err := c.err
	c.err = nil
	return err
}

func (c Config) ToCmdLineFormat() []string {
	args := make([]string, 0)
	for k, v := range c.Options {
		args = append(args, "--"+k)
		args = append(args, v)
	}
	return args
}
