package main

import (
	"log"
	"message-relayer/service"
)

func main() {
	log.Printf("** MessageRelayer service starting... **")
	service.New()
	//if err := run(); err != nil {
	//	fmt.Fprintf(os.Stdout, "%s\n", err)
	//	os.Exit(1)
	//}
}

//func run() error {
//	return service.New()
//}