package main

import (
	"log"
	"message-relayer/app"
)

func main() {
	log.Printf("** MessageRelayer service starting... **")
	app.New()
	//if err := run(); err != nil {
	//	fmt.Fprintf(os.Stdout, "%s\n", err)
	//	os.Exit(1)
	//}
}

//func run() error {
//	return app.New()
//}