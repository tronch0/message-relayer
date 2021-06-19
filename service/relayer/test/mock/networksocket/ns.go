package networksocket

import (
	"fmt"
	"message-relayer/service/model"
	"message-relayer/service/model/messagetype"
)

func New() model.NetworkSocket {
	return &NS{
		c: 5,
	}
}

type NS struct {
	c int
}

func (n *NS) Read() (model.Message, error){
	res := model.Message{Type: messagetype.Undefined, Data: nil}

	if n.c < 0 {
		return res, fmt.Errorf("no more messages")
	}

	n.c = n.c - 1
	res.Type = messagetype.StartNewRound
	res.Data = []byte("“An ounce of prevention is worth a pound of cure.” - B.F")

	return res, nil
}
