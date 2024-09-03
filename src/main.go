package main

import (
	"log"

	"github.com/Quiqui-dev/distributed_storage_go/p2p"
)


func main() {

	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr: ":3000",
		Decoder: p2p.DefaultDecoder{},
		HandshakeFunc: p2p.NOPHandshakeFunc,
	}
	
	log.Println("Starting app")

	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}