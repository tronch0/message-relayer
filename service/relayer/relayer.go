package relayer

import (
	"log"
	"message-relayer/service/config"
	model2 "message-relayer/service/model"
	messagetype2 "message-relayer/service/model/messagetype"
	"message-relayer/service/utils"
)

//const (
//	queueSizeMessageType1 = 2
//	queueSizeMessageType2 = 1
//)
//
//var (
//	messageImportanceOrderDESC  = []messagetype.MessageType{messagetype.StartNewRound,messagetype.ReceivedAnswer}
//)

func NewRelayer(socket model2.NetworkSocket, logger *log.Logger, config *config.Config) *Relayer {
	res := &Relayer{
		logger: logger,
		socket: socket,
		subscriberMap: make(map[messagetype2.MessageType][]chan<- model2.Message),
		messagesQueues: make(map[messagetype2.MessageType]*utils.Stack),
		messageTypeImportanceDesc: config.MessageTypeImportanceOrderDesc,
	}

	for msgType, queueSize := range config.MessageTypeToQueueSize {
		res.messagesQueues[msgType] = utils.NewStack(queueSize)
	}

	logger.Println("instantiated relayer")

	return res
}

type Relayer struct {
	socket model2.NetworkSocket
	logger *log.Logger

	subscriberMap map[messagetype2.MessageType][]chan<- model2.Message
	messagesQueues map[messagetype2.MessageType]*utils.Stack
	messageTypeImportanceDesc []messagetype2.MessageType
}

func (r *Relayer) SubscribeToMessages(msgType messagetype2.MessageType, messages chan<- model2.Message) {

	if subscribers, isFound := r.subscriberMap[msgType]; isFound {
		r.subscriberMap[msgType] = append(subscribers, messages)
	} else {
		r.subscriberMap[msgType] = []chan<- model2.Message{messages}
	}

	r.logger.Printf("relayer - new subscriber for message-type %d", msgType)
}

func (r *Relayer) Start() {
	maxErrCounter := 3 // we should setup a termination policy, specific error type or max error count
	r.logger.Println("relayer - start listening")
	for msg, err := r.socket.Read(); maxErrCounter > 0 ; {
		if err != nil {
			r.logger.Printf("error reading a message, err: %v", err)
			maxErrCounter--
		}
		r.queueMessage(msg)
	}


	r.processQueuedMessages()
}

func (r *Relayer) queueMessage(msg model2.Message) {
	r.logger.Printf("relayer - queue message (type: %d)", msg.Type)
	r.messagesQueues[msg.Type].Push(msg)
}

func (r *Relayer) processQueuedMessages() {

	// iterate over messages by type (from the most important to the lest important)
	for _, msgType := range r.messageTypeImportanceDesc {

		// iterate over all messages from the same type
		for {
			msg := r.messagesQueues[msgType].Pop()
			if msg == nil {
				break
			}

			r.logger.Printf("relayer - broadcast message (type: %d) to %d subscribers", msg.Type, len(r.subscriberMap[msg.Type]))

			// iterate over all subscribers and broadcast
			for _, subscriber := range r.subscriberMap[msg.Type] {
				subscriber <- *msg
			}
		}
	}

}



//r.logger.Printf("CRITICAL ERROR reached maximum errors count")
//r.logger.Printf("service shutting down...")