package p2p

import (
	"testing"
    "github.com/stretchr/testify/assert"

)


func TestTCPTransport(t *testing.T) {
	listenAddress := ":3000"
	opts := TCPTransportOpts{
		ListenAddress: listenAddress,	
		HandShakeFunc: NOPHandShakeFunc,
		Decoder: GOBDecoder{},
	}
	tr := NewTCPTransport(opts)

	assert.Equal(t, listenAddress, tr.ListenAddress)

	// we need the server to listen 
	assert.Nil(t, tr.ListenAndAccept())


}