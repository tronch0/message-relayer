package model

import "message-relayer/model/messagetype"

type MessageRelayer interface {
	SubscribeToMessages(msgType messagetype.MessageType, messages chan<- Message)
}
