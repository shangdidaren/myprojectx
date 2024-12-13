package main

import (
	"my_projectx/network"
	"time"
)

func main() {
	trLocal := network.NewLocalTransport("Local")
	trRemote := network.NewLocalTransport("Remote")
	err := trLocal.Connect(trRemote)
	if err != nil {
		panic(err)
	}
	err = trRemote.Connect(trLocal)
	if err != nil {
		panic(err)
	}

	go func() {
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				_ = trLocal.SendMessage(trRemote.Addr(), []byte("hello world"))
			} else {
				_ = trRemote.SendMessage(trLocal.Addr(), []byte("hello world"))
			}
			time.Sleep(1 * time.Second)
		}

	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal, trRemote},
	}
	s := network.NewServer(
		opts,
	)
	s.Start()
}
