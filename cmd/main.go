package main

import (
	"log"
	"message-relayer/service"
	configuration "message-relayer/service/model/config"
	"message-relayer/service/model/messagetype"
)

func main() {
	log.Printf("** MessageRelayer service starting... **")

	c := getServiceConfig()
	service.New(c)
}

// in a normal case, we'll get all config definitions from the user (from a file or env-var)
func getServiceConfig() *configuration.Config {
	broadcastOrder := []messagetype.MessageType{messagetype.StartNewRound, messagetype.ReceivedAnswer}

	msgTypeToQueueSize := make(map[messagetype.MessageType]int)
	msgTypeToQueueSize[messagetype.StartNewRound] = 2
	msgTypeToQueueSize[messagetype.StartNewRound] = 1

	return &configuration.Config{
		MsgTypeStoredLength:   msgTypeToQueueSize,
		MsgTypeBroadcastOrder: broadcastOrder,
		LogToFile:             false,
	}
}
