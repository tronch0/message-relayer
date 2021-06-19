package service

import (
	"fmt"
	"log"
	"message-relayer/service/config"
	"message-relayer/service/model"
	"message-relayer/service/model/messagetype"
	"message-relayer/service/relayer"
	"message-relayer/service/test/networksocket"
	sub2 "message-relayer/service/test/subscriber"
	"os"
)

func New() {

	socket := networksocket.New()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	config := getRelayerConfig()

	r := relayer.NewRelayer(socket, logger, config)
	doneChan := setupSubscribers(r,logger)
	r.Start()

	if <-doneChan != true {
		fmt.Println("assert error")
	}
}

func getRelayerConfig() *config.Config {
	importanceOrder := []messagetype.MessageType{messagetype.StartNewRound, messagetype.ReceivedAnswer}

	msgTypeToQueueSize := make(map[messagetype.MessageType]int)
	msgTypeToQueueSize[messagetype.StartNewRound] = 2
	msgTypeToQueueSize[messagetype.StartNewRound] = 1

	return  &config.Config{
		MessageTypeToQueueSize: msgTypeToQueueSize,
		MessageTypeImportanceOrderDesc: importanceOrder,
	}
}

func setupSubscribers(r model.MessageRelayer, logger *log.Logger) chan bool {
	s, doneChan := sub2.New(logger, r)
	go s.Listen()

	return doneChan
}



