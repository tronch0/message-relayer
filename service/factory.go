package service

import (
	"fmt"
	"log"
	model2 "message-relayer/service/model"
	messagetype2 "message-relayer/service/model/messagetype"
	"message-relayer/service/relayer"
	"os"
	"strconv"
	"time"
)

func New() {

	socket := getNetworkSocket()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	config := getRelayerConfig()

	r := relayer.NewRelayer(socket, logger, config)
	setupSubscribers(r)
	r.Start()
}

func getRelayerConfig() *model2.Config {
	importanceOrder := []messagetype2.MessageType{messagetype2.StartNewRound, messagetype2.ReceivedAnswer}

	msgTypeToQueueSize := make(map[messagetype2.MessageType]int)
	msgTypeToQueueSize[messagetype2.StartNewRound] = 2
	msgTypeToQueueSize[messagetype2.StartNewRound] = 1

	return  &model2.Config{
		MessageTypeToQueueSize: msgTypeToQueueSize,
		MessageTypeImportanceOrderDesc: importanceOrder,
	}
}

func getNetworkSocket() model2.NetworkSocket {
	return &NS{
		c: 1,
	}
}

type NS struct {
	c int
}

func (n *NS) Read() (model2.Message, error){
	currSec := time.Now().Second()

	res := model2.Message{Type: messagetype2.Undefined, Data: nil}

	if currSec == 35 {
		return res, fmt.Errorf("couldn't retrive message")
	}
	if n.c != currSec {
		res.Type = messagetype2.StartNewRound
		res.Data = []byte(strconv.Itoa(time.Now().Second()))
	}

	n.c = currSec

	return res, nil
}


func setupSubscribers(r model2.MessageRelayer) {

}