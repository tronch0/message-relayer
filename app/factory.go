package app

import (
	"fmt"
	"log"
	"message-relayer/app/relayer"
	"message-relayer/model"
	"message-relayer/model/messagetype"
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

	select {}
}

func getRelayerConfig() *model.Config {
	importanceOrder := []messagetype.MessageType{messagetype.StartNewRound, messagetype.ReceivedAnswer}

	msgTypeToQueueSize := make(map[messagetype.MessageType]int)
	msgTypeToQueueSize[messagetype.StartNewRound] = 2
	msgTypeToQueueSize[messagetype.StartNewRound] = 1

	return  &model.Config{
		MessageTypeToQueueSize: msgTypeToQueueSize,
		MessageTypeImportanceOrderDesc: importanceOrder,
	}
}

func getNetworkSocket() model.NetworkSocket {
	return &NS{
		c: 1,
	}
}

type NS struct {
	c int
}

func (n *NS) Read() (model.Message, error){
	currSec := time.Now().Second()

	res := model.Message{Type: messagetype.Undefined, Data: nil}

	if currSec == 35 {
		return res, fmt.Errorf("couldn't retrive message")
	}
	if n.c != currSec {
		res.Type = messagetype.StartNewRound
		res.Data = []byte(strconv.Itoa(time.Now().Second()))
	}

	n.c = currSec

	return res, nil
}


func setupSubscribers(r model.MessageRelayer) {

}