package main

import (
	"blockchain_go/network"
	"time"
)

// 服务端 链接
// 区块 交易
// 加密

func main() {
	transLocal := network.NewLocalTransport("LOCAL")
	transRemote := network.NewLocalTransport("REMOTE")

	transLocal.Connect(transRemote)
	transRemote.Connect(transLocal)

	go func() {
		for {
			// transRemote.SendMessage(transLocal.Addr(), []byte("Hello world"))
			transLocal.SendMessage(transRemote.Addr(), []byte("Hello world"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		// Transports: []network.Transport{transLocal},
		Transports: []network.Transport{transRemote},
	}
	s := network.NewServer(opts)
	s.Start()
}
