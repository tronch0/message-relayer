package model

import (
	"message-relayer/service/model/messagetype"
)

type Message struct {
	Type messagetype.MessageType
	Data []byte
}
