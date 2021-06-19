package subscriber

import (
	"log"
	"message-relayer/service/model"
	"message-relayer/service/model/messagetype"
)

type Subscriber struct {
	incomingChan chan model.Message
	logger       *log.Logger
	doneChan chan bool
}

func New(logger *log.Logger, r model.MessageRelayer) (*Subscriber,chan bool) {
	c := make(chan model.Message)
	doneChan := make(chan bool)

	res := &Subscriber{
		incomingChan: c,
		logger:       logger,
		doneChan: doneChan,
	}

	r.SubscribeToMessages(messagetype.StartNewRound,res.incomingChan)

	return res, doneChan
}

func (s *Subscriber) Listen() {
	for msg := range s.incomingChan {
		s.logger.Printf("subscriber: got new message, type: %d, body: %s", msg.Type,msg.Data)
	}
	s.doneChan <- true
}
