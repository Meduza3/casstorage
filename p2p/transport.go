package p2p

// Peer represents the remote node
type Peer interface {
}

// Transport is anything that handles the communication between the nodes
// (TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
}
