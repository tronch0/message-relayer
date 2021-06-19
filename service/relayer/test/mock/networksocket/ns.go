package networksocket

import (
	"fmt"
	"message-relayer/service/model"
	"message-relayer/service/model/messagetype"
)

func New(messagesToSend []model.Message) model.NetworkSocket {
	return &NS{
		messagesToSend: messagesToSend,
	}
}

type NS struct {
	messagesToSend []model.Message
}

func (n *NS) Read() (model.Message, error){
	if len(n.messagesToSend) == 0 {
		return model.Message{Type: messagetype.Undefined, Data: nil}, fmt.Errorf("no more messages")
	}
	res := n.messagesToSend[0]

	n.messagesToSend = n.messagesToSend[1:]

	return res, nil
}

