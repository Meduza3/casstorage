package p2p

import "testing"

func TestNewTCPTransportSetsListenAddress(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)

	if tr.listenAddress != listenAddr {
		t.Errorf("Wanted: %q got: %q", listenAddr, tr.listenAddress)
	}
}

func TestListenAndAcceptOKAddr(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)

	err := tr.ListenAndAccept()
	if err != nil {
		t.Errorf("Critical error")
	}
}

func TestListenAndAcceptFailsWithIncorrectAddr(t *testing.T) {
	listenAddr := "xxxxxx"
	tr := NewTCPTransport(listenAddr)

	err := tr.ListenAndAccept()
	if err == nil {
		t.Errorf("Shouldn't have accepted %q as a listenAddr", listenAddr)
	}
}
