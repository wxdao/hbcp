package hbcp

import (
	"net"
	"strings"
)

type Context struct {
	conn net.Conn
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
