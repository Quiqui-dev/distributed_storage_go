package p2p

import "net"

// holds any data being sent over each transport
// between two nodes
type Message struct {
	From net.Addr
	Payload []byte
}