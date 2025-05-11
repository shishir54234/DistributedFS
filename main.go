package main

import (
	"fmt"
	"github.com/shishir54234/DistributedFS/p2p"
	"log"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3004",
		HandShakeFunc: p2p.NOPHandShakeFunc,
		Decoder:       &p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	fmt.Println("HEYY")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatalf("Failed to listen and accept: %v", err)
	}

}
