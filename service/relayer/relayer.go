package relayer

import (
	"log"
	"message-relayer/service/model"
	config2 "message-relayer/service/model/config"
	"message-relayer/service/model/messagetype"
	"message-relayer/service/utils"
)

type Relayer struct {
	socket model.NetworkSocket
	logger *log.Logger

	typeToSubscribers   map[messagetype.MessageType][]chan<- model.Message
	subscribersChannels map[chan<- model.Message]bool
	typeToSavedMsgs     map[messagetype.MessageType]*utils.Stack
	typeBroadcastOrder  []messagetype.MessageType
}

func NewRelayer(socket model.NetworkSocket, logger *log.Logger, config *config2.Config) *Relayer {
	res := &Relayer{
		logger:              logger,
		socket:              socket,
		typeToSubscribers:   make(map[messagetype.MessageType][]chan<- model.Message),
		subscribersChannels: make(map[chan<- model.Message]bool),
		typeToSavedMsgs:     make(map[messagetype.MessageType]*utils.Stack),
		typeBroadcastOrder:  config.MsgTypeBroadcastOrder,
	}

	for msgType, queueSize := range config.MsgTypeStoredLength {
		res.typeToSavedMsgs[msgType] = utils.NewStack(queueSize)
	}

	res.logger.Println("instantiated relayer")

	return res
}

func (r *Relayer) SubscribeToMessages(msgType messagetype.MessageType, msgChan chan<- model.Message) {

	if subscribers, isFound := r.typeToSubscribers[msgType]; isFound {
		r.typeToSubscribers[msgType] = append(subscribers, msgChan)
	} else {
		r.typeToSubscribers[msgType] = []chan<- model.Message{msgChan}
	}

	r.subscribersChannels[msgChan] = true

	r.logger.Printf("relayer - added new subscriber for message-type %d", msgType)
}

func (r *Relayer) Listen() { // we should setup a termination policy, specific error type, signal channel or  max error count
	r.logger.Println("relayer - start listening")
	go r.processIncomingTraffic()
}
func (r *Relayer) processIncomingTraffic() {
	r.consumeAndStoreMessages()
	r.processQueuedMessages()
	r.closedAllSubscribersChannels()
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
			r.typeToSavedMsgs[msg.Type].Push(msg)
		}
	}
}

func (r *Relayer) processQueuedMessages() {

	for _, msgType := range r.typeBroadcastOrder { // iterate over messages by type in broadcast order (DESC)
		messagesStack, isExist := r.typeToSavedMsgs[msgType]
		if isExist == false {
			continue
		}

		for msg := messagesStack.Pop(); msg != nil; msg = messagesStack.Pop() { // iterate over all messages from the same type

			r.logger.Printf("relayer - broadcast message (type: %d) to %d subscribers", msg.Type, len(r.typeToSubscribers[msg.Type]))

			for _, subscriber := range r.typeToSubscribers[msg.Type] { // iterate over all subscribers subscribed to this type
				//  non blocking send
				select {
				case subscriber <- *msg:
				default:
				}

			}
		}

	}
}

func (r *Relayer) closedAllSubscribersChannels() {
	for c := range r.subscribersChannels {
		close(c)
	}
}
