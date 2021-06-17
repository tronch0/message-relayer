package main

import (
	"fmt"
	"log"
	"message-relayer/app"
	"os"
)

func main() {
	log.Printf("** MessageRelayer starting... **")

	if err := run(); err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	return app.New()
}