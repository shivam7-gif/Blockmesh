package p2p

// handshakeFunc is 

type HandshakeFunc func(any) error


func NOPHandshakeFunc(any) error {return nil}
