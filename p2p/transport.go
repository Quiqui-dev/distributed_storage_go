package p2p

// Represents the remote node (remote connection to us)
type Peer interface {

}


// anything that handles the communication
// between nodes in the network
type Transport interface {

	ListenAndAccept() error
}