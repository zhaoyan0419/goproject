package netProgram

import "testing"

func TestTcpServer(t *testing.T) {
	TcpServer()
}
func TestTcpClient(t *testing.T) {
	TcpClient()
}
func TestTcpTimeOutClient(t *testing.T) {
	TcpTimeOutClient()
}

func TestTcpBackLogServer(t *testing.T) {
	TcpBackLogServer()
}

func TestTcpBackLogClient(t *testing.T) {
	TcpBackLogClient()
}

func TestTcpServerRW(t *testing.T) {
	TcpServerRW()
}

func TestTcpClientRW(t *testing.T) {
	TcpClientRW()
}
