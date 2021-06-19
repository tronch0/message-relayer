package subscriber

import (
	"bytes"
	"log"
	"message-relayer/service/model"
	"message-relayer/service/model/messagetype"
	"testing"
)

type Subscriber struct {
	incomingChan chan model.Message
	logger       *log.Logger
	doneChan chan bool
	subscriberId int
	expectedMsgs []model.Message
	t *testing.T
}

func New(logger *log.Logger, r model.MessageRelayer, subscriberId int, msgs []model.Message, t *testing.T) (*Subscriber,chan bool) {
	c := make(chan model.Message)
	doneChan := make(chan bool)

	res := &Subscriber{
		incomingChan: c,
		logger:       logger,
		doneChan: doneChan,
		subscriberId: subscriberId,
		expectedMsgs: msgs,
		t: t,
	}

	r.SubscribeToMessages(messagetype.StartNewRound,res.incomingChan)

	return res, doneChan
}

func (s *Subscriber) Listen() {
	for msg := range s.incomingChan {
		s.logger.Printf("subscriber-%d: got new message, type: %d, body: %s", s.subscriberId,msg.Type,msg.Data)

		found := false
		for _, expectedMsg := range s.expectedMsgs {
			if expectedMsg.Type == msg.Type && bytes.Equal(expectedMsg.Data, msg.Data) {
				found = true
			}
		}

		if found == false {
			s.t.Fatalf("message recived is not in the expected messages")
		}
	}
	s.doneChan <- true
}
