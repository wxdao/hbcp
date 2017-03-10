package hbcp

import (
	"bufio"
	"strings"
)

type Msg map[string]string

func ConstructMsg(reader *bufio.Reader) (msg Msg, err error) {
	msg = Msg{}
	for {
		row, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		row = strings.TrimSuffix(row, "\n")
		if row != "" {
			subs := strings.Split(row, ":")
			value := strings.Join(subs[1:], ":")
			msg[subs[0]] = value
			continue
		}
		return msg, err
	}
}
