package p2p

type Message struct {
	remoteAddr string
	Payload []byte
}