package rpc

import (
	"encoding/binary"
	"io"
	"net"
)

// 测试网络中读写数据的情况

// 会话连接的结构体
type Session struct {
	conn net.Conn
}

// 构造方法
func NewSession(conn net.Conn) *Session {
	return &Session{conn: conn}
}

// 向连接中去写数据
func (s *Session) Write(data []byte) error {
	// 定义写数据的格式
	// 4字节头部 + 可变体的长度
	buf := make([]byte, 4+len(data))
	// 写入头部，记录数据长度
	binary.BigEndian.PutUint32(buf[:4], uint32(len(data))) // 写入4字节大端序长度头
	// 将整个数据，放到4后边
	copy(buf[4:], data)
	_, err := s.conn.Write(buf)
	if err != nil {
		return err
	}
	return nil
}

// 从连接读数据
func (s *Session) Read() ([]byte, error) {
	// 读取头部记录的长度
	header := make([]byte, 4)
	// 按长度读取消息
	_, err := io.ReadFull(s.conn, header)
	if err != nil {
		return nil, err
	}
	// 读取数据
	dataLen := binary.BigEndian.Uint32(header)
	data := make([]byte, dataLen)
	_, err = io.ReadFull(s.conn, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
