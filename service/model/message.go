package model

import (
	messagetype2 "message-relayer/service/model/messagetype"
)

type Message struct {
	Type messagetype2.MessageType
	Data []byte
}





