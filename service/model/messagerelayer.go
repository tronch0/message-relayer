package model

import (
	messagetype2 "message-relayer/service/model/messagetype"
)

type MessageRelayer interface {
	SubscribeToMessages(msgType messagetype2.MessageType, messages chan<- Message)
}
