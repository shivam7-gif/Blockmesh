
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
		shakeHands: NOPHandshakeFunc,
		listenAddress: listenAddr,
		peers:         make(map[net.Addr]*TCPPeer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {

	var err error

	t.listener ,err = net.Listen("tcp",t.listenAddress)

	if err != nil{
		return err
	}
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


type Temp struct{}
 

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(peer, true)

	if err := t.shakeHands(peer); err != nil{
		conn.Close()

		return ErrInvalidHandshake
	}
	// Read loop
	msg := &Temp{}
	for{
		if err := t.shakeHands.Decode(conn , msg) ; err != nil{

				fmt.Printf("tcp error : %s\n",err)
				continue 
			
		}
	}

}
