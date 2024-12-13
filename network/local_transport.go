package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.RWMutex //保证peers线程安全
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}
func (t *LocalTransport) Consume() <-chan RPC {
	// todo 不懂
	return t.consumeCh
}
func (t *LocalTransport) Connect(tr Transport) error {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.peers[tr.Addr()] = tr.(*LocalTransport) // todo: why not lock tr.peers, and put t.Addr() to tr.peers
	// it's doesn't matter, it is what it is

	return nil
}

func (t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	t.lock.RLock()
	defer t.lock.RUnlock()

	peer, ok := t.peers[to]
	if !ok {
		return fmt.Errorf("%s can not send message to %s", t.addr, to)
	}
	peer.consumeCh <- RPC{
		From:    t.addr,
		Payload: payload,
	}

	return nil
}
