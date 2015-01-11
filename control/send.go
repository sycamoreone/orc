// Package control implements part of the 'Tor control protocol (Version 1)'.
// See https://gitweb.torproject.org/torspec.git/blob/HEAD:/control-spec.txt
package control

import (
	"errors"
	"net"
	"net/textproto"
)

// Conn represents a Tor Control Protocol connection to a Tor server.
type Conn struct {
	// text is the textproto.Conn used by the Conn.
	text *textproto.Conn

	AsyncReplies chan *Reply
	SyncReplies  chan *Reply
}

// Client returns a new Tor Control Protocol connection
// using conn as the underlying transport.
func Client(conn net.Conn) *Conn {
	c := new(Conn)
	text := textproto.NewConn(conn)
	c.text = text
	c.AsyncReplies = make(chan *Reply)
	c.SyncReplies = make(chan *Reply)
	return c
}

// Dial connects to the given network address using net.Dial and then
// starts and returns a new Tor Control Protocol connection.
func Dial(addr string) (*Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return Client(conn), nil
}

// Cmd represents a command send from the client to the server.
type Cmd struct {
	Keyword   string
	Arguments []string // optional list of arguments; may be nil
	Data      string   // will be dot escaped by Send
}

// Send sends a command to the Tor server.
func (c Conn) Send(cmd Cmd) (err error) {
	if len(cmd.Keyword) == 0 {
		return errors.New("empty Keyword in Cmd")
	}
	line := cmd.Keyword
	for _, arg := range cmd.Arguments {
		if len(arg) != 0 {
			line = line + " " + arg
		}
	}
	_, err = c.text.Cmd("%s", line)

	if len(cmd.Data) == 0 {
		return err
	}
	if cmd.Keyword[0] != '+' {
		return errors.New("protocol error: CmdData present, but Keyword no leading '+;")
	}
	w := c.text.DotWriter()
	defer w.Close()
	_, err = w.Write([]byte(cmd.Data))
	return err
}

// dquote double quotes the string s.
func dquote(s string) string {
	return "\"" + s + "\""
}

// Auth authenticates a connection using the hashed password mechanism.
// Pass an empty string to authenticate without password.
func (c Conn) Auth(passwd string) (err error) {
	cmd := Cmd{Keyword: "AUTHENTICATE", Arguments: []string{dquote(passwd)}}
	err = c.Send(cmd)
	if err != nil {
		return err
	}
	reply, err := c.Receive()
	if err != nil {
		return err
	}
	if reply.Status != 250 || reply.Text != "OK" {
		return errors.New("authentication error: " + reply.Text)
	}
	return nil
}

// GetInfo sends a GETINFO command to the server.
func (c Conn) GetInfo(key string) error {
	cmd := Cmd{Keyword: "GETINFO", Arguments: []string{key}}
	err := c.Send(cmd)
	return err
}

// Resolve launches a remote hostname lookup for addr.
func (c Conn) Resolve(addr string) error {
	cmd := Cmd{Keyword: "RESOLVE", Arguments: []string{addr}}
	err := c.Send(cmd)
	return err
}

// SetEvents sends a SETEVENTS command to the server.
func (c Conn) SetEvents(keys []string) error {
	cmd := Cmd{Keyword: "SETEVENTS", Arguments: keys}
	err := c.Send(cmd)
	return err
}

type Signal string

const (
	SignalReload        Signal = "RELOAD"        // Reload config items.
	SignalShutdown      Signal = "SHUTDOWN"      // Controlled shutdown.
	SignalDump          Signal = "DUMP"          // Dump log information about open connections and circuits.
	SignalDebug         Signal = "DEBUG"         // Switch all open logs to loglevel debug.
	SignalHalt          Signal = "HALT"          // Immediate shutdown: clean up and exit now.
	SignalClearDNSCache Signal = "CLEARDNSCACHE" // Forget the client-side cached IPs for all hostnames.
	SignalNewNym        Signal = "NEWNYM"        // Switch to clean circuits, so new application requests don't share any circuits with old ones.
	SignalHeartbeat     Signal = "HEARTBEAT"     // Dump an unscheduled Heartbeat message to log.
)

// Signal sends a SIGNAL command to the server.
func (c Conn) Signal(s Signal) error {
	cmd := Cmd{Keyword: "SIGNAL", Arguments: []string{string(s)}}
	err := c.Send(cmd)
	return err
}
