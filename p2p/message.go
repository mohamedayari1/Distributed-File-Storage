package p2p

import "net"

// RPC holds any arbitrary data that is being sent over the transport/network
type RPC struct {
	RemoteAddr net.Addr
	Payload []byte
}