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
		buffer += strings.Join([]string{k, v.String()}, ":") + "\n"
	}
	buffer += "\n"
	_, err = context.conn.Write([]byte(buffer))
	return
}

func (context *Context) RemoteAddr() net.Addr {
	return context.conn.RemoteAddr()
}

func (context *Context) LocalAddr() net.Addr {
	return context.conn.LocalAddr()
}

func (context *Context) Close() (err error) {
	err = context.conn.Close()
	context.office.handler.EmitClose(context)
	return
}
