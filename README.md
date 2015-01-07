orc - Onion router control protocol library.
============================================

    go get github.com/sycamoreone/orc/control

Only some low-level functionality at the moment.

Examples 
---------

The examples assume that a Tor router with ControlPort 9051 open and not
protected by a password is running on localhost.  You can start such a
router temporarily to run an example:

    > /usr/sbin/tor -f examples/torrc
    > go run examples/resolve/main.go

