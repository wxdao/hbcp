package hbcp

import (
	"net"
	"strings"
)

type Context struct {
	office *Office
	conn   net.Conn
}

func (context *Context) Respond(msg Msg) (err error) {
	buffer := ""
	for k, v := range msg {
		buffer += strings.Join([]string{k, v}, ":") + "\n"
	}
	buffer += "\n"
	_, err = context.conn.Write([]byte(buffer))
	return
}

func (context *Context) RemoteAddr() (addr net.Addr) {
	addr = context.conn.RemoteAddr()
	return addr
}

func (context *Context) LocalAddr() (addr net.Addr) {
	addr = context.conn.LocalAddr()
	return addr
}

func (context *Context) Close() (err error) {
	err = context.conn.Close()
	context.office.handler.CallCloseFunc(context)
	return
}
