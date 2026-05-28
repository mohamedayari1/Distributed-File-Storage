package p2p

import (
	"fmt"
	"net"
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
	OnPeer func(Peer) error
}

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpcchan chan RPC
	
}


func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpcchan: make(chan RPC),
	}
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer{
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}

func (peer *TCPPeer) Close() error {
	return peer.conn.Close() 
}


// Consume implements the transport interface, which will return a read-only channel
// for  reading incoming messages/RPCs sent from another peer
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcchan
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
	peer := NewTCPPeer(conn, outbound)
	if err := t.HandShakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP HandShakeError: %s\n", err)
	}	

	if t.OnPeer != nil {
		if err := t.OnPeer(peer); err != nil {
			conn.Close()
			fmt.Printf("TCP Peer Error (OnPeer method failed): %s", err)
		}
	}
	rpc := RPC{}
	rpc.RemoteAddr = conn.RemoteAddr()	
	for {

		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("TCP error: %s\n", err)
			continue
		}
		fmt.Printf("New Message: %+v  :", string(rpc.Payload))

		// The channel is supposed to be read only right ? why we are able to write on it ???
		t.rpcchan <- rpc
	}
	



}

