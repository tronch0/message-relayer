package main

import (
	"log"
	"message-relayer/service"
)

func main() {
	log.Printf("** MessageRelayer service starting... **")
	service.New()
}