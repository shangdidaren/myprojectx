package network

type NetAddr string

type RPC struct {
	From    NetAddr
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC
	Connect(transport Transport) error
	SendMessage(netAddr NetAddr, payload []byte) error
	Addr() NetAddr
}
