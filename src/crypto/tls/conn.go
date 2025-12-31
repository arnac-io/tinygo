package tls

import (
	"net"
)

// A TLSConn represents a secured connection.
// It implements the net.Conn interface.
type TLSConn struct {
	*net.TLSConn
}

func (c *TLSConn) ConnectionState() ConnectionState {
	panic("not implemented TLSConn.ConnectionState")
	return ConnectionState{}
}
