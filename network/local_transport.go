package network

import (
	"fmt"
	"sync"
)

// 搭建本地传输层， 实现transport接口
type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.RWMutex
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) Transport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC, 1024),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (local *LocalTransport) Consume() <-chan RPC {
	return local.consumeCh
}

func (local *LocalTransport) Connect(lt Transport) error {
	local.lock.Lock()
	defer local.lock.Unlock()
	local.peers[lt.Addr()] = lt.(*LocalTransport)
	return nil
}

func (local *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	local.lock.Lock()
	defer local.lock.Unlock()
	peer, ok := local.peers[to]
	if !ok {
		fmt.Errorf("could not send message from %s to %s", local.addr, to)
	}

	peer.consumeCh <- RPC{
		From:    local.addr,
		Payload: payload,
	}
	return nil
}

func (local *LocalTransport) Addr() NetAddr {
	return local.addr
}
