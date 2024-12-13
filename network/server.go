package network

import (
	"fmt"
)

type ServerOpts struct {
	Transports []Transport
}

type Server struct {
	ServerOpts

	rpcCh chan RPC
	quit  chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quit:       make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	//ticker := time.NewTicker(1000 * time.Second)
free:
	for {
		select {
		case rpc := <-s.rpcCh:
			// TODO: handle protocol
			fmt.Println(rpc)
		case <-s.quit:
			break free
			//case <-ticker.C:
			//	fmt.Println("tick n second")
		}
	}
	fmt.Println("Server stopped")
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)

	}
}
