package model

import "message-relayer/model/messagetype"

type Message struct {
	Type messagetype.MessageType
	Data []byte
}





