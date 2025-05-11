package p2p

import (
	//"bytes"
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn     net.Conn
	outbound bool // if this peer is the one making the connection this bool is true

}
type TCPTransportOpts struct {
	ListenAddr    string
	HandShakeFunc HandShakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TcpOpts  TCPTransportOpts
	listener net.Listener
	mu       sync.RWMutex
	peers    map[net.Addr]Peer
}

func NewTCPTransport(tcpOpts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{TcpOpts: tcpOpts,
		peers: make(map[net.Addr]Peer),
	}
}

func (transport *TCPTransport) ListenAndAccept() error {
	var err error
	transport.listener, err = net.Listen("tcp", transport.TcpOpts.ListenAddr)
	transport.startAcceptLoop()
	return err
}
func (transport *TCPTransport) startAcceptLoop() {
	for {
		conn, err := transport.listener.Accept()
		transport.mu.Lock()
		transport.handleConn(conn)
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}
	}
}

type Temp struct{}

func (transport *TCPTransport) handleConn(conn net.Conn) {

	fmt.Println("new incoming connection:", conn.RemoteAddr())
	peer := NewTCPTransport(transport.TcpOpts)
	if err := transport.TcpOpts.HandShakeFunc(peer); err != nil {
		fmt.Println("handshake error:", err)

		conn.Close()
		return

	}
	msg := &Message{}
	//buf := make([]byte, 1024)
	for {
		if err := transport.TcpOpts.Decoder.Decode(conn, msg); err != nil {
			fmt.Println("decode error:", err)
			continue
		}
		fmt.Println("decode message:", msg)
	}
	fmt.Printf("handshake finished: %s\n", conn.RemoteAddr())
}
