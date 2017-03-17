package hbcp

import (
	"bufio"
	"strings"
)

type MsgValue struct {
	buf []byte
}

func NewMsgValue(buf []byte) *MsgValue {
	return &MsgValue{buf: buf}
}

func (v *MsgValue) String() string {
	return string(v.buf)
}

func (v *MsgValue) Bytes() []byte {
	return v.buf
}

type Msg map[string]MsgValue

type MetaHandler func (param string, value string) ([]byte, error)

func ConstructMsg(reader *bufio.Reader, metaHandlers map[string]MetaHandler) (msg Msg, err error) {
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
			keyWithMeta := subs[0]
			keySubs := strings.Split(keyWithMeta, ",")
			key := keySubs[0]
			meta := strings.Join(keySubs[1:], ",")
			mSubs := strings.Split(meta, ";")
			tag := mSubs[0]
			param := strings.Join(mSubs[1:], ";")
			if tag != "" {
				handler, exists := metaHandlers[tag]
				if !exists {
					continue
				}
				rvalue, err := handler(param, value)
				if err != nil {
					continue
				}
				msg[key] = *NewMsgValue(rvalue)
			} else {
				msg[key] = *NewMsgValue([]byte(value))
			}
			continue
		}
		return msg, err
	}
}
