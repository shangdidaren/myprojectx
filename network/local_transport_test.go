package network

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	err := tra.Connect(trb)
	assert.Nil(t, err)
	err = trb.Connect(tra)
	assert.Nil(t, err)

	assert.Equal(t, tra.peers[trb.Addr()], trb)
	assert.Equal(t, trb.peers[tra.Addr()], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	err := trb.Connect(tra)
	assert.Nil(t, err)
	err = tra.Connect(trb)
	assert.Nil(t, err)

	message := []byte("Hello World")
	err = tra.SendMessage(trb.Addr(), message)
	assert.Nil(t, err)

	rpc := <-trb.consumeCh
	assert.Equal(t, rpc.Payload, message)
}
