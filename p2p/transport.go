package p2p

// Peer represents a node in the network
type Peer interface {
}

// transport is anything that handles the communication between nodes in the networks
// this can be of the form(tcp , udp , websockets)
type Transport interface {
	ListenAndAccept() error
}
