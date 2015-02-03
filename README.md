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

TODO
-----

I started writing orc because I wanted to play with the control protocol.
and I am planing to add to the library according to my own interests.
Still, if anybody is interested in a particular feature, please tell me!
I would be more than happy if this thing is actually useful to somebody.

Things that I would like to add include:

- [x] Add Handler functions for asynchronous events.
- [ ] Instead of sending a SETEVENTS command and installing a custom 
  Handler in two steps, there should be a catalog of ready-made functions
  that make sending the SETEVENTS command and installing a Handler a
  single call.
- [ ] Add a package to launch a slave Tor process.
- [ ] Add a nice way to provide and save configurations.
- [ ] Add helpers for launching and configuration of hidden services.
- [ ] Status codes like "551 Internal error" shouldn't be just integers
  but implement the error interface.
- [ ] Try to implement a few of the example use cases from the
  [Stem](https://stem.torproject.org/index.html) and
  [txtorcon](https://txtorcon.readthedocs.org/en/latest/) documentation,
  to see what functionality is missing.
- [ ] And, of course, we need to have yet another small and fun CLI client.
