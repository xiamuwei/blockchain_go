package network

type NetAddr string

type RPC struct {
	From    NetAddr
	Payload []byte
}

type Transport interface {
	Consume() <-chan RPC               // 消费
	Connect(Transport) error           // 链接
	SendMessage(NetAddr, []byte) error // 发送消息
	Addr() NetAddr                     // 地址
}
