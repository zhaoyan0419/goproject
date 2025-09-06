package netProgram

import (
	"errors"
	"log"
	"net"
	"sync"
	"time"
)

// 连接池接口
type Pool interface {

	// 获取连接
	Get() (net.Conn, error)
	// 放回连接，不是关闭
	Put(conn net.Conn) error
	// 释放池，关闭全部连接
	Release() error
	// 有效连接的长度
	Len() int
}

// 连接工厂接口
type ConnFactory interface {
	// 生产连接
	Factory(addr string) (net.Conn, error)
	// 关闭连接
	Close(net.Conn) error
	// Ping
	Ping(net.Conn) error
}

// 连接池典型配置
type PoolConfig struct {
	// 初始化连接数,至少保持多少个连接
	InitConnNum int
	// 最大连接数，池中最多支持多少连接
	MaxConnNum int
	// 最大空闲连接数，池中最多有多少可用的连接
	MaxIdleNum int
	// 空闲连接超时时间，多久后空闲连接会被释放
	IdleTimeout time.Duration

	// 生产连接的工厂
	Factory ConnFactory
}

// 空闲连接类型（管理的连接）
type IdleConn struct {
	// 连接本身
	conn net.Conn
	// 时间，用于判断空闲是否超时
	putTime time.Time
}

// 连接池结构
type TcpPool struct {
	// 配置信息
	config PoolConfig
	// 运行时信息
	// 使用的连接数量
	openingConnNum int
	// 空闲的连接链表，chan
	idleList chan *IdleConn
	//并发安全锁
	mu sync.RWMutex
}

// Tcp连接工厂类型
type TcpConnFactory struct {
}

// 产生连接方法
func (*TcpConnFactory) Factory(addr string) (net.Conn, error) {
	// 校验参数合理性
	if addr == "" {
		return nil, errors.New("addr is empty")
	}
	// 建立连接
	conn, err := net.DialTimeout("tcp", addr, time.Second*5)
	if err != nil {
		return nil, err
	}
	// return
	return conn, err
}

// close关闭连接
func (*TcpConnFactory) Close(conn net.Conn) error {
	return conn.Close()
}

// Ping健康检测
func (*TcpConnFactory) Ping(conn net.Conn) error {
	return nil
}

// TcpPool实现Pool接口
func (*TcpPool) Get() (net.Conn, error) {
	return nil, nil
}
func (*TcpPool) Put(conn net.Conn) error {
	return nil
}
func (*TcpPool) Release() error {
	log.Println("Release All Connections")
	return nil
}
func (*TcpPool) Len() int {
	return 0
}

const defautlMaxConnNum = 100
const defautlInitConnNum = 10

func NewTcpPool(addr string, PoolConfig PoolConfig) (*TcpPool, error) {

	// 校验参数
	// j考验工厂的存在
	if PoolConfig.Factory == nil {
		return nil, errors.New("Factory is not exists")
	}

	if addr == "" {
		return nil, errors.New("addr is empty")
	}

	// 最大连接数校验
	if PoolConfig.MaxConnNum == 0 {
		// a.return错误
		return nil, errors.New("max conn is zero")
		// b. 人为修改一个合理的
		PoolConfig.MaxConnNum = defautlMaxConnNum
	}
	if PoolConfig.InitConnNum == 0 {
		PoolConfig.InitConnNum = defautlInitConnNum
	} else if PoolConfig.InitConnNum > PoolConfig.MaxConnNum {
		PoolConfig.InitConnNum = PoolConfig.MaxConnNum
	}
	if PoolConfig.MaxIdleNum == 0 {
		PoolConfig.MaxIdleNum = PoolConfig.InitConnNum
	} else if PoolConfig.MaxIdleNum > PoolConfig.MaxConnNum {
		PoolConfig.MaxIdleNum = PoolConfig.MaxConnNum
	}

	// 初始化TcpPool对象
	pool := TcpPool{
		config:         PoolConfig,
		openingConnNum: 0,
		idleList:       make(chan *IdleConn, PoolConfig.MaxConnNum),
		mu:             sync.RWMutex{},
	}
	// 初始化连接
	// 根据InitConnNum的配置来创建
	for i := 0; i < PoolConfig.InitConnNum; i++ {
		conn, err := pool.config.Factory.Factory(addr)
		if err != nil {
			// 通常以为这，连接池初始化失败
			// 释放已经存在的连接
			pool.Release()
			return nil, err
		}
		// 连接创建成功
		// 加入到空闲连接队列中
		pool.idleList <- &IdleConn{
			conn:    conn,
			putTime: time.Now(),
		}
	}
	// 返回
	return &pool, nil
}
