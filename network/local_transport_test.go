package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	transA := NewLocalTransport("A")
	transB := NewLocalTransport("B")

	transA.Connect(transB)
	transB.Connect(transA)
	assert.Equal(t, transA.peers[transB.addr], transB)
	assert.Equal(t, transB.peers[transA.addr], transA)

}

func TestSendMessage(t *testing.T) {
	transA := NewLocalTransport("A")
	transB := NewLocalTransport("B")
	transA.Connect(transB)
	transB.Connect(transA)

	msg := []byte("Hello blockchain")
	// 发送消息
	n := transA.SendMessage(transB.addr, msg)

	rpc := <-transB.Consume()
	// 是否为空
	assert.Nil(t, n)
	// 消息是否一致
	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, transA.addr)
}
