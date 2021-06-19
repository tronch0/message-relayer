package model

import (
	"message-relayer/service/model/messagetype"
)

type MessageRelayer interface {
	SubscribeToMessages(msgType messagetype.MessageType, messages chan<- Message)
}
