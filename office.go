package hbcp

import (
	"bufio"
	"net"
	"encoding/base64"
)

type Office struct {
	handler Handler
	mode    int
	metaHandlers map[string]MetaHandler
}

func NewOffice(handler Handler, metaHandlers map[string]MetaHandler) (office Office) {
	office = Office{}
	office.handler = handler
	if metaHandlers == nil {
		metaHandlers = make(map[string]MetaHandler)
	}
	metaHandlers["b64"] = b64
	office.metaHandlers = metaHandlers
	return
}

func (office *Office) Serve(address string) (err error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Temporary() {
				continue
			}
			return err
		}
		go office.Attach(conn)
	}
}

func (office *Office) Join(address string) (err error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	office.Attach(conn)
	return
}

func (office *Office) Attach(conn net.Conn) (err error) {
	office.handler.EmitJoin(&Context{office: office, conn: conn})
	err = office.handleConn(conn)
	return
}

func (office *Office) handleConn(conn net.Conn) (err error) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := ConstructMsg(reader, office.metaHandlers)
		if err != nil {
			conn.Close()
			office.handler.EmitClose(&Context{conn: conn})
			return err
		}
		office.handler.EmitMsg(&Context{office: office, conn: conn}, msg)
	}
}

func b64(param string, value string) (rvalue []byte, err error) {
	rvalue, err = base64.StdEncoding.DecodeString(value)
	return
}
