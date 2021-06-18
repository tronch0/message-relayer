package model

import "message-relayer/model/messagetype"

type NetworkSocket interface {
	Read() (Message, error)
}

type Message struct {
	Type messagetype.MessageType
	Data []byte
}

type MessageRelayer interface {
	SubscribeToMessages(msgType messagetype.MessageType, messages chan<- Message)
}



