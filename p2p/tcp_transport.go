package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn

	// if we are accepting the conn --> outbound == false
	// if we are dialing to a remote node --> outbound == true
	outbound bool
}



type TCPTransportOpts struct {
	ListenAddress string 
	HandShakeFunc HandShakeFunc
	Decoder Decoder
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener


	mu sync.RWMutex
	peers map[net.Addr]Peer
}


func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,

	}
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer{
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}



func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		fmt.Printf("TCP error : %s\n", err)
		return err
	}

	go t.startAcceptLoop()
	return err

}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		
		if err != nil {
			fmt.Printf("TCP Error : %s\n", err)
		}

		fmt.Printf("New incoming connection %+v\n", conn)
		go t.handleConnection(conn, true)
	}
}



func (t *TCPTransport) handleConnection(conn net.Conn, outbound bool) {
	// peer := NewTCPPeer(conn, true)

	// if err := t.HandShakeFunc(peer); err != nil {
	// 	conn.Close()
	// 	fmt.Printf("TCP HandShakeError: %s\n", err)
	// }	

	msg := &Message{}
	for {

		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}
		fmt.Printf("New Message: %+v  :", string(msg.Payload))

	}
	



}

