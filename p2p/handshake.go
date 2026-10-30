package p2p

// ErrInvalidHandshake is returned if the handshake between
// the local and remote node could not be established

// handshakeFunc is

type HandshakeFunc func(Peer) error

func NOPHandshakeFunc(Peer) error { return nil }
