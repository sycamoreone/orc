orc - Onion router control protocol library.
============================================

    go get github.com/sycamoreone/orc/control

Only some low-level functionality at the moment.

To do anything with this library you will have to read the
[control protocol specification](https://gitweb.torproject.org/torspec.git/tree/control-spec.txt)
beforehand.

Examples 
---------

The examples assume that a Tor router with ControlPort 9051 open and protected
by a password is running on localhost. You can start such a router temporarily
to run an example:

    > /usr/sbin/tor -f examples/torrc
    > go run examples/resolve/main.go

