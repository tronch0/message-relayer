package subscriber

import (
	"log"
	"message-relayer/service/model"
	"message-relayer/service/model/messagetype"
)

type Subscriber struct {
	incomingChan chan model.Message
	logger       *log.Logger
	msgChan      chan model.Message
	subscriberId int
}

func New(
	logger *log.Logger,
	r model.MessageRelayer,
	subscriberId int,
	subscribeToTypes []messagetype.MessageType,
	buffered bool,
	) (*Subscriber,chan model.Message) {

	outputChan := make(chan model.Message)
	var incomingChan chan model.Message

	if buffered {
		incomingChan = make(chan model.Message, 10)
	} else {
		incomingChan = make(chan model.Message)
	}

	res := &Subscriber{
		incomingChan: incomingChan,
		logger:       logger,
		msgChan:      outputChan,
		subscriberId: subscriberId,
	}

	for _, msgType := range subscribeToTypes {
		r.SubscribeToMessages(msgType,res.incomingChan)
		res.logger.Printf("subscriber %d subscribed to %d message type", msgType)
	}

	return res, outputChan
}

func (s *Subscriber) Listen() {
	go s.listen()
}

func (s *Subscriber) listen() {
	for msg := range s.incomingChan {
		s.logger.Printf("subscriber-%d: got new message, type: %d, body: %s", s.subscriberId,msg.Type,msg.Data)
		s.msgChan <- msg
	}

	close(s.msgChan)
}
