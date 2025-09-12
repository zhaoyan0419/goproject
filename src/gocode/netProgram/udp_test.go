package netProgram

import "testing"

func TestUdpServerBasic(t *testing.T) {
	UdpServerBasic()
}

func TestUdpClientBasic(t *testing.T) {
	UdpClientBasic()
}

func TestUdpServerPeer(t *testing.T) {
	UdpServerPeer()
}

func TestUdpClientPeer(t *testing.T) {
	UdpClientPeer()
}

func TestUDPReceiverMulticast(t *testing.T) {
	UDPReceiverMulticast()
}

func TestUDPSenderMulticast(t *testing.T) {
	UDPSenderMulticast()
}
