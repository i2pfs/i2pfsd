package i2p

import (
	"github.com/i2pfs/i2pfsd/consts"
	"github.com/xaionaro-go/log"
	"github.com/xaionaro-go/errors"
	"net"
)

type Connection interface {
	net.Conn

	Loop()
	SendMessage(msg interface{}) error
	IsClosed() bool
	SetMessageHandler(newHandler func(conn Connection, buf []byte) error)
}

type connection struct {
	net.Conn
	messageHandler func(conn Connection, buf []byte) error
	isClosed       bool
}

func NewConnection(rawConn net.Conn) Connection {
	return &connection{
		Conn: rawConn,
	}
}

func (conn *connection) SetMessageHandler(newHandler func(conn Connection, buf []byte) error) {
	conn.messageHandler = newHandler
}

func (conn *connection) Loop() {
	buf := make([]byte, consts.MESSAGE_MAX_SIZE)
	for {
		log.Debugln("Connection: Read")
		n, err := conn.Read(buf)
		if err != nil {
			conn.Close()
			return
		}
		log.Debugln("Connection: received: " + string(buf[:n]))
		err = conn.messageHandler(conn, buf[:n])
		if err != nil {
			conn.Close()
			return
		}
		if conn.IsClosed() {
			return
		}
	}
}

func (conn connection) IsClosed() bool {
	return conn.isClosed
}

func (conn *connection) Close() error {
	err := conn.Conn.Close()
	if err == nil {
		conn.isClosed = true
	}
	return err
}

func (conn *connection) SendMessage(msg interface{}) error {
	return errors.NotImplemented
}
