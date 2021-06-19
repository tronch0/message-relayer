package service

import (
	"log"
	"message-relayer/service/model"
	configuration "message-relayer/service/model/config"
	"message-relayer/service/relayer"
	"os"
)

func New(config *configuration.Config) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	socket := getNetworkSocket()
	r := relayer.NewRelayer(socket, logger, config)

	setupSubscribers(r,logger)
	r.Listen()
}

func setupSubscribers(r model.MessageRelayer, logger *log.Logger)  {
	// here we can subscribe to the relayer
}

func getNetworkSocket() model.NetworkSocket {
	return nil
}
