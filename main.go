package main

import (
	"fmt"
	"log"

	"github.com/mohamedayari1/Distributed-File-Storage/p2p"
)

func main() {
	opts := p2p.TCPTransportOpts{
		ListenAddress: ":3000",	
		HandShakeFunc: p2p.NOPHandShakeFunc,
		Decoder: p2p.DefaultDecoder{},
	}
	
	tr := p2p.NewTCPTransport(opts)	

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("we gucci")
	select {}
}   