package p2p

import (
	"fmt"
	"net"
	"sync"
)




type TCPTransport struct {
	listenAddress string
	listener net.Listener

	mu sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
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

		go t.handleConnection(conn)
	}
}



func (t *TCPTransport) handleConnection(conn net.Conn) {
	fmt.Printf("New incoming connection %+v\n", conn)
}

