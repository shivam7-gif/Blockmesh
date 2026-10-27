
package p2p

import (
	"fmt"
	"net"
	"sync"
)
// TCPPeer represents the remote node over a tcp established connections.
type TCPPeer struct {
	// conn is the underlying connection of the peer
	conn     net.Conn
	// if we dial and retreive a conn => outbound == true
	// if we accept and retreieve a conn => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder  Decoder
	mu            sync.RWMutex
	peers         map[net.Addr]*TCPPeer
}




func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		HandshakeFunc: NOPHandshakeFunc,
		listenAddress: listenAddr,
		peers:         make(map[net.Addr]*TCPPeer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	t.listener = ln
	fmt.Printf("TCP server listening on %s\n", ln.Addr())

	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
			return
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, false)


	if err := t.shakeHands(conn); err != nil{
		conn.
	}
	buf := new(bytes.Buffer)

	// Read loop
	for{
		n , _  := conn.Read(buf)
		// msg := buff[:n]

	}
	t.mu.Lock()
	t.peers[conn.RemoteAddr()] = peer
	t.mu.Unlock()

	fmt.Printf("New incoming connection from: %s\n", conn.RemoteAddr())
}
