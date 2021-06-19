package relayer

import (
	"log"
	"message-relayer/service/config"
	"message-relayer/service/model"
	"message-relayer/service/model/messagetype"
	"message-relayer/service/utils"
)

type Relayer struct {
	socket model.NetworkSocket
	logger *log.Logger

	subscriberMap map[messagetype.MessageType][]chan<- model.Message
	messagesQueues map[messagetype.MessageType]*utils.Stack
	messageTypeImportanceDesc []messagetype.MessageType
}

func NewRelayer(socket model.NetworkSocket, logger *log.Logger, config *config.Config) *Relayer {
	res := &Relayer{
		logger: logger,
		socket: socket,
		subscriberMap: make(map[messagetype.MessageType][]chan<- model.Message),
		messagesQueues: make(map[messagetype.MessageType]*utils.Stack),
		messageTypeImportanceDesc: config.MessageTypeImportanceOrderDesc,
	}

	for msgType, queueSize := range config.MessageTypeToQueueSize {
		res.messagesQueues[msgType] = utils.NewStack(queueSize)
	}

	logger.Println("instantiated relayer")

	return res
}

func (r *Relayer) SubscribeToMessages(msgType messagetype.MessageType, messages chan<- model.Message) {

	if subscribers, isFound := r.subscriberMap[msgType]; isFound {
		r.subscriberMap[msgType] = append(subscribers, messages)
	} else {
		r.subscriberMap[msgType] = []chan<- model.Message{messages}
	}

	r.logger.Printf("relayer - added new subscriber for message-type %d", msgType)
}

func (r *Relayer) Start() { // we should setup a termination policy, specific error type or max error count
	r.logger.Println("relayer - start listening")

	r.consumeAndStoreMessages()
	r.processQueuedMessages()
}

func (r *Relayer) consumeAndStoreMessages() {
	maxErrCounter := 3
	for maxErrCounter > 0 {

		msg, err := r.socket.Read()
		if err != nil {
			r.logger.Printf("error reading a message, err: %v", err)
			maxErrCounter--
		} else {
			r.logger.Printf("relayer - queue message (type: %d)", msg.Type)
			r.messagesQueues[msg.Type].Push(msg)
		}
	}
}

func (r *Relayer) processQueuedMessages() {


	for _, msgType := range r.messageTypeImportanceDesc { // iterate over messages by type in importance order (DESC)


		messagesStack, isExist := r.messagesQueues[msgType]
		if isExist == false {
			continue
		}

		for msg := messagesStack.Pop(); msg != nil; msg = messagesStack.Pop() { // iterate over all messages from the same type

			r.logger.Printf("relayer - broadcast message (type: %d) to %d subscribers", msg.Type, len(r.subscriberMap[msg.Type]))

			for _, subscriber := range r.subscriberMap[msg.Type] { // iterate over all subscribers subscribed to this type
				subscriber <- *msg
			}
		}

		r.logger.Printf("adsahuifehufhnesuihvsiuhkjbckjashciohsaiobjkfabkjasbdas")

	}
	r.logger.Printf("adsahuifehufhnesuihvsiuhkjbckjashciohsaiobjkfabkjasbdas")
}