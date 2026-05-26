package p2p

// Peer interface is the represantation 
// of the remote node we are communicating with 
type Peer interface {
	Close() error
}

// Transport is the socket,
//  this means how are we doing communication
//  with the nodes using (TCP, UDP, Websockets)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}