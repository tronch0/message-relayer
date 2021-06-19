package service

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"message-relayer/service/config"
	"message-relayer/service/model"
	"message-relayer/service/model/messagetype"
	"message-relayer/service/relayer"
	"message-relayer/service/utils/sub"
)

func New() {

	socket := getNetworkSocket()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	config := getRelayerConfig()

	r := relayer.NewRelayer(socket, logger, config)
	setupSubscribers(r,logger)
	r.Start()

	time.Sleep(10 * time.Second)
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

func getNetworkSocket() model.NetworkSocket {
	return &NS{
		c: 5,
	}
}

type NS struct {
	c int
}

func (n *NS) Read() (model.Message, error){
	res := model.Message{Type: messagetype.Undefined, Data: nil}

	if n.c < 0 {
		return res, fmt.Errorf("no more messages")
	}

	n.c = n.c - 1
	res.Type = messagetype.StartNewRound
	res.Data = []byte(strconv.Itoa(time.Now().Second()))

	return res, nil
}


func setupSubscribers(r model.MessageRelayer, logger *log.Logger) {
	s := sub.New(logger, r)
	go s.Listen()
}



