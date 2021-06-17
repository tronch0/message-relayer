package relayer

import (
	"log"
	"message-relayer/model"
)

func NewRelayer(socket model.NetworkSocket, logger log.Logger) *Relayer {
	res := &Relayer{
		logger: logger,
		socket: socket,
		subscriberMap: make(map[model.MessageType][]chan<-model.Message),
	}

	go res.listen()
	logger.Println("instantiate relayer")

	return res
}

type Relayer struct {
	socket model.NetworkSocket
	logger log.Logger
	subscriberMap map[model.MessageType][]chan<-model.Message
}

func (r *Relayer) SubscribeToMessages(msgType model.MessageType, messages chan<- model.Message) {

	if subscribers, isFound := r.subscriberMap[msgType]; isFound {
		r.subscriberMap[msgType] = append(subscribers, messages)
	} else {
		r.subscriberMap[msgType] = []chan<-model.Message{messages}
	}

	r.logger.Printf("relayer - new subscriber for message-type %d", msgType)
}

func (r *Relayer) listen() {
	maxErrCounter := 5 // we should setup a termination policy, specific error type or max error count
	r.logger.Println("relayer - start listening")
	for msg, err := r.socket.Read(); maxErrCounter > 0 ; {
		if err != nil {
			r.logger.Printf("error reading a message, err: %v", err)
			maxErrCounter--
		}
		r.processMessage(msg)
	}
}

func (r *Relayer) processMessage(msg model.Message) {
	r.logger.Printf("relayer - broadcast message (type: %d) to %d subscribers", msg.Type, len(r.subscriberMap[msg.Type]))

	for _, subscriber := range r.subscriberMap[msg.Type] {
		subscriber <- msg
	}
}
