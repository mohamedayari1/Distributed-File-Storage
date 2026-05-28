package main

import (
	"fmt"
	"log"

	"github.com/mohamedayari1/Distributed-File-Storage/p2p"
)

func OnPeer(p2p.Peer) error {
	fmt.Println("Doing some logic with the Peer outside of the TCPTransport")
	return nil 
}


func main() {
	opts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",	
		HandShakeFunc: p2p.NOPHandShakeFunc,
		Decoder: p2p.DefaultDecoder{},
		OnPeer: OnPeer,
	}

	tr := p2p.NewTCPTransport(opts)	

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("we gucci")	
	go func() {
		msg := <- tr.Consume()
		fmt.Printf("%+v Sent : %+v\n", msg.RemoteAddr, string(msg.Payload))
	}()
	
	select {}
}   

