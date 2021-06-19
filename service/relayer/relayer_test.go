package relayer

import (
	"fmt"
	"io"
	"log"
	"message-relayer/service/model"
	configuration "message-relayer/service/model/config"
	"message-relayer/service/model/messagetype"
	"message-relayer/service/relayer/test/mock/networksocket"
	"message-relayer/service/relayer/test/mock/subscriber"
	"os"
	"testing"
)

// Test Scenarios:
// 1. Empty - No incoming traffic, No Subs.
// 2. No incoming with subscribers.
// 3. Incoming with no subscribers.
// 4. Exceed error count -> relayer should terminate.
// 5. Happy flow - 5 subs - 1 message type.
// 6. Happy flow - 1 subs - 3 message type.
// 7. Happy flow - 3 subs - 3 message type


func TestRelayerHappyPath(t *testing.T) {
	logger := log.New(getLogOutput(true), "", log.Ldate|log.Ltime|log.Lshortfile)
	config := getServiceConfig()

	msgs := []model.Message{
		model.Message{Type: messagetype.StartNewRound, Data: []byte("1")},
		model.Message{Type: messagetype.StartNewRound, Data: []byte("2")},
		model.Message{Type: messagetype.ReceivedAnswer, Data: []byte("3")},
		model.Message{Type: messagetype.ReceivedAnswer, Data: []byte("4")},
	}
	socket := networksocket.New(msgs)

	r := NewRelayer(socket, logger, config)

	s1, s1Chan := subscriber.New(logger, r, 1, []model.Message{}, t)
	go s1.Listen()

	r.Listen()

	if <-s1Chan != true {
		t.Fatalf("expected to get finish ack")
		fmt.Println("assert error")
	}
}


func getServiceConfig() *configuration.Config {
	importanceOrder := []messagetype.MessageType{messagetype.StartNewRound, messagetype.ReceivedAnswer}

	msgTypeToQueueSize := make(map[messagetype.MessageType]int)
	msgTypeToQueueSize[messagetype.StartNewRound] = 2
	msgTypeToQueueSize[messagetype.ReceivedAnswer] = 1

	return  &configuration.Config{
		MessageTypeToQueueSize: msgTypeToQueueSize,
		MessageTypeImportanceOrderDesc: importanceOrder,
	}
}

//func getMessages(r model.MessageRelayer, logger *log.Logger) chan bool {
//	res.Type = messagetype.StartNewRound
//	res.Data = []byte(fmt.Sprintf("“An ounce of prevention is worth a pound of cure.” - B.F (%d)", n.c + 1))
//}

func getLogOutput(isFile bool) io.Writer {
	if isFile {
		logFile, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		return logFile
	}

	return os.Stdout
}