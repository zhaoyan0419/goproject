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

func TestServerWriteBlock(t *testing.T) {
	ServerWriteBlock()
}
func TestClientNoneRead(t *testing.T) {
	ClientNoneRead()
}

func TestServerListen(t *testing.T) {
	ServerListen()
}

func TestClientDial(t *testing.T) {
	ClientDial()
}

// 格式化消息传输测试
func TestServerListenFormat(t *testing.T) {
	ServerListenFormat()
}
func TestClientDialFormat(t *testing.T) {
	ClientDialFormat()
}
