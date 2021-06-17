package model

type NetworkSocket interface {
	Read() (Message, error)
}

type Message struct {
	Type MessageType
	Data []byte
}

type MessageRelayer interface {
	SubscribeToMessages(msgType MessageType, messages chan<- Message)
}



