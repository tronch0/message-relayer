package app

import (
	"log"
	"message-relayer/app/relayer"
	"message-relayer/model"
)

func New() error {

	socket := getNetworkSocket()
	logger := log.Logger{}

	r := relayer.NewRelayer(socket, logger)
	setupSubscribers(r)

	return nil
}

func getNetworkSocket() model.NetworkSocket {

	return nil
}


func setupSubscribers(r model.MessageRelayer) {

}