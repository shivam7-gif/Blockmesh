package p2p

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOpts struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	handshakeFunc HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]*TCPPeer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	if opts.HandshakeFunc == nil {
		opts.HandshakeFunc = NOPHandshakeFunc
	}

	if opts.Decoder == nil {
		opts.Decoder = &GOBDecoder{}
	}

	return &TCPTransport{
		listenAddress: opts.ListenAddr,
		handshakeFunc: opts.HandshakeFunc,
		decoder:       opts.Decoder,
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

	if err := t.handshakeFunc(peer); err != nil {
		fmt.Printf("TCP handshake error: %s\n", err)
		conn.Close()
		return
	}

	t.mu.Lock()
	t.peers[conn.RemoteAddr()] = peer
	t.mu.Unlock()

	fmt.Printf("New peer connected: %s\n", conn.RemoteAddr())

	defer func() {
		conn.Close()

		t.mu.Lock()
		delete(t.peers, conn.RemoteAddr())
		t.mu.Unlock()
	}()
	msg := &Message{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			if err != io.EOF {
				fmt.Printf("TCP decode error: %s\n", err)
				continue
			}
			fmt.Printf("message : %+v\n",msg)
		}
	}
}

