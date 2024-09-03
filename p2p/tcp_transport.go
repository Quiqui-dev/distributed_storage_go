package p2p

import (
	"fmt"
	"net"
	"sync"
)

// represents the remote node over TCP connection
type TCPPeer struct {
	conn net.Conn

	// dial to another peer would be outbound to the peer
	// if we accept from the peer it will be inbound
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}


type TCPTransportOpts struct {
	ListenAddr string 
	HandshakeFunc HandshakeFunc
	Decoder Decoder
}

type TCPTransport struct {
	TCPTransportOpts TCPTransportOpts
	listener net.Listener


	mu sync.RWMutex // mutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
	}
} 

func (t *TCPTransport) ListenAndAccept() error {
	
	var err error
	
	t.listener, err = net.Listen("tcp", t.TCPTransportOpts.ListenAddr)

	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}


func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept() 

		if err != nil {
			fmt.Printf("tcp accept error: %s\n", err)
		}

		
		go t.handleConn(conn)
	}
}

type Temp struct {}

func (t *TCPTransport) handleConn(conn net.Conn) {

	peer := NewTCPPeer(conn, true)

	if err := t.TCPTransportOpts.HandshakeFunc(peer); err != nil {
		fmt.Printf("tcp handshake error: %s\n", err)
		conn.Close()
		return 
	}

	// Read loop
	msg := &Message{}
	for {
		if err := t.TCPTransportOpts.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("tcp error: %s\n", err)
			continue
		}

		msg.From = conn.RemoteAddr()

		fmt.Printf("message: %+v\n", msg)
	}
}