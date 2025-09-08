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
	// 连接地址
	addr string
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
func (pool *TcpPool) Get() (net.Conn, error) {
	// 锁定
	pool.mu.Lock()
	defer pool.mu.Unlock()
	// 获取连接，若没有，则创建连接
	for {

		select {
		case idleConn, ok := <-pool.idleList:
			// 判断channel是否被关闭
			if !ok {
				return nil, errors.New("idle list was closed")
			}
			// 判断连接是否超时
			// pool.config.IdleTimeout 比较idleConn.putTime
			if pool.config.IdleTimeout > 0 { // >0 说明设置了超时时间
				// 判断puttime + timeout是否在now之前
				if idleConn.putTime.Add(pool.config.IdleTimeout).Before(time.Now()) {
					// 如果创建时间+超时时间在当前时间之前，说明当前已经超时了
					// 关闭连接并且查找下一个连接
					_ = pool.config.Factory.Close(idleConn.conn)
					continue
				}
			}

			// 判断连接是否可用
			if err := pool.config.Factory.Ping(idleConn.conn); err != nil {
				// ping失败，连接不可用
				// 关闭连接继续查找
				_ = pool.config.Factory.Close(idleConn.conn)
				continue
			}
			// 进行超时判断和连接可用性判断，都通过后代表找到了可用连接，直接返回即可
			// 使用的连接计数
			log.Println("get conn from idleconn")
			pool.openingConnNum++
			return idleConn.conn, nil
		default:
			// 如果channel中没有可用连接，那么就走default进行创建新连接
			// 创建连接之前需要判断，是否到达连接池最大连接数
			if pool.openingConnNum >= pool.config.MaxConnNum {
				return nil, errors.New("max opening connection")
				// 另一种方案，阻塞
				// continue
			}
			// 创建连接
			conn, err := pool.config.Factory.Factory(pool.addr)
			if err != nil {
				return nil, err
			}
			log.Println("get conn from factory")
			pool.openingConnNum++
			return conn, nil
		}
	}
}
func (pool *TcpPool) Put(conn net.Conn) error {
	// 锁
	pool.mu.Lock()
	defer pool.mu.Unlock()
	// 安全校验
	if conn == nil {
		return errors.New("connection is not exists")
	}
	// 判断空闲连接列表是否存在
	if pool.idleList == nil {
		// 关闭连接
		_ = pool.config.Factory.Close(conn)
		return errors.New("idlelist is not exists")
	}
	// 放回连接
	select {
	case pool.idleList <- &IdleConn{
		conn:    conn,
		putTime: time.Now(),
	}:
		pool.openingConnNum--
		return nil
	default:
		_ = pool.config.Factory.Close(conn)
		return nil
	}
}

// 释放连接池
func (pool *TcpPool) Release() error {
	// 并发安全锁
	pool.mu.Lock()
	defer pool.mu.Unlock()
	// 确定连接池是否被释放
	if pool.idleList == nil {
		return nil
	}
	// 关闭idlelist
	close(pool.idleList)
	// 释放全部空闲连接
	for i := range pool.idleList {
		_ = pool.config.Factory.Close(i.conn)
		// i.conn.Close()
	}

	return nil
}
func (pool *TcpPool) Len() int {
	return len(pool.idleList)
}

const defautlMaxConnNum = 10
const defautlInitConnNum = 1

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
		addr:           addr,
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
