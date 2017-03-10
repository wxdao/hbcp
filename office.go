package hbcp

import (
	"bufio"
	"net"
)

const (
	MODE_NIL   = 0
	MODE_SERVE = 1
	MODE_JOIN  = 2
)

type Office struct {
	handler Handler
	mode    int
}

func NewOffice(handler Handler) (office Office) {
	office = Office{}
	office.mode = MODE_NIL
	office.handler = handler
	return
}

func (office *Office) Serve(address string) (err error) {
	listener, err := net.Listen("tcp", address)
	office.mode = MODE_SERVE
	defer func() {
		office.mode = MODE_NIL
	}()
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
		office.handler.CallJoinFunc(&Context{conn: conn})
		go office.handleConn(conn)
	}
}

func (office *Office) Join(address string) (err error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	office.mode = MODE_JOIN
	defer func() {
		office.mode = MODE_NIL
	}()
	office.handler.CallJoinFunc(&Context{conn: conn})
	err = office.handleConn(conn)
	return
}

func (office *Office) Send(address string) (err error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return
	}
	office.mode = MODE_JOIN
	defer func() {
		office.mode = MODE_NIL
	}()
	err = office.handleConn(conn)
	return
}

func (office *Office) handleConn(conn net.Conn) (err error) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := ConstructMsg(reader)
		if err != nil {
			conn.Close()
			office.handler.CallCloseFunc(&Context{conn: conn})
			return err
		}
		office.handler.CallMsgFunc(&Context{conn: conn}, msg)
	}
}
