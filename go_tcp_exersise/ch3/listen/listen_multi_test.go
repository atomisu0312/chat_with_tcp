package ch03

import (
	"net"
	"testing"
)

func TestListenerMulti(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:55000")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("bound to %q", listener.Addr())

	defer func() { _ = listener.Close() }()

	for {
		conn, err := listener.Accept()

		if err != nil {
			t.Log(err)
			return
		}
		go func(c net.Conn) {
			defer c.Close()

			// show connection details
			t.Logf("accepted connection from %s", c.RemoteAddr())
		}(conn)
	}

}
