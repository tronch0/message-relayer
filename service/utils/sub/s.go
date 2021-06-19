package sub

import (
	"log"
	"message-relayer/service/model"
	"message-relayer/service/model/messagetype"
)

type Subscriber struct {
	incomingChan chan model.Message
	logger       *log.Logger
}

func New(logger *log.Logger, r model.MessageRelayer) *Subscriber {
	c := make(chan model.Message)

	res := &Subscriber{
		incomingChan: c,
		logger:       logger,
	}

	r.SubscribeToMessages(messagetype.StartNewRound,res.incomingChan)


	return res
}

func (s *Subscriber) Listen() {
	for msg := range s.incomingChan {
		s.logger.Printf("subscriber: got new message, type: %d, body: %s", msg.Type,msg.Data)
	}
}
