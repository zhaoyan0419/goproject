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

// 短连接测试验证
func TestServerListenShort(t *testing.T) {
	ServerListenShort()
}

func TestClientDialShort(t *testing.T) {
	ClientDialShort()
}

// 心跳检测测试验证
func TestServerListenHB(t *testing.T) {
	ServerListenHB()
}

func TestClientDialHB(t *testing.T) {
	ClientDialHB()
}

func TestServerListenPool(t *testing.T) {
	ServerListenPool()
}

func TestClientUseTcpPool(t *testing.T) {
	ClientUseTcpPool()
}

// 粘包测试
func TestTcpServerSticky(t *testing.T) {
	TcpServerSticky()
}

func TestTcpClientSticky(t *testing.T) {
	TcpClientSticky()
}

// 自定义编解码器解决粘包问题测试
func TestTcpServerCoder(t *testing.T) {
	TcpServerCoder()
}

func TestTcpClientCoder(t *testing.T) {
	TcpClientCoder()
}

// 验证tcpconn

func TestTcpServerSpecial(t *testing.T) {
	TcpServerSpecial()
}

func TestTcpClientSpecial(t *testing.T) {
	TcpClientSpecial()
}
