package relayer

import (
	"bytes"
	"io"
	"log"
	"message-relayer/service/model"
	configuration "message-relayer/service/model/config"
	"message-relayer/service/model/messagetype"
	"message-relayer/service/relayer/testing/mock/networksocket"
	"message-relayer/service/relayer/testing/mock/subscriber"
	"os"
	"testing"
)

func TestOneSubTwoMsgType(t *testing.T) {
	config := getServiceConfig()
	logger := log.New(getLogOutput(config.LogToFile), "", log.Ldate|log.Ltime|log.Lshortfile)

	msgs := []model.Message{
		{Type: messagetype.StartNewRound, Data: []byte("1")},
		{Type: messagetype.StartNewRound, Data: []byte("2")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("3")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("4")},
		{Type: messagetype.StartNewRound, Data: []byte("5")},
		{Type: messagetype.StartNewRound, Data: []byte("6")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("7")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("8")},
	}
	expectedOutputIds := [][]byte{[]byte("6"), []byte("5"), []byte("8")}

	socket := networksocket.New(msgs)
	relayer := NewRelayer(socket, logger, config)

	sub1, sub1MsgChan := subscriber.New(
		logger,
		relayer,
		1,
		[]messagetype.MessageType{messagetype.StartNewRound, messagetype.ReceivedAnswer},
		true,
	)

	sub1.Listen()
	relayer.Listen()

	// assert result messages
	i := 0
	for msg := range sub1MsgChan {
		if i >= len(expectedOutputIds) || bytes.Equal(msg.Data, expectedOutputIds[i]) == false {
			t.Fatalf("error on result assert, expected msg.data: %s, actual: %s", expectedOutputIds[i], msg.Data)
		}
		i++
	}

	// assert result length
	if i != len(expectedOutputIds) {
		t.Fatalf("error on result length assert, expected len: %d, actual: %d", len(expectedOutputIds), i)
	}

}

func TestOneSubOneMsgTypeReceivedAnswer(t *testing.T) {
	config := getServiceConfig()
	logger := log.New(getLogOutput(config.LogToFile), "", log.Ldate|log.Ltime|log.Lshortfile)

	msgs := []model.Message{
		{Type: messagetype.StartNewRound, Data: []byte("1")},
		{Type: messagetype.StartNewRound, Data: []byte("2")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("3")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("4")},
		{Type: messagetype.StartNewRound, Data: []byte("5")},
		{Type: messagetype.StartNewRound, Data: []byte("6")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("7")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("8")},
	}
	expectedOutputIds := [][]byte{[]byte("8")}

	socket := networksocket.New(msgs)
	relayer := NewRelayer(socket, logger, config)

	sub1, sub1MsgChan := subscriber.New(
		logger,
		relayer,
		1,
		[]messagetype.MessageType{messagetype.ReceivedAnswer},
		true,
	)

	sub1.Listen()
	relayer.Listen()

	// assert result messages
	i := 0
	for msg := range sub1MsgChan {

		if i >= len(expectedOutputIds) || bytes.Equal(msg.Data, expectedOutputIds[i]) == false {
			t.Fatalf("error on result assert, expected msg.data: %s, actual: %s", expectedOutputIds[i], msg.Data)
		}
		i++
	}

	// assert result length
	if i != len(expectedOutputIds) {
		t.Fatalf("error on result length assert, expected len: %d, actual: %d", len(expectedOutputIds), i)
	}

}

func TestOneSubOneMsgTypeStartNewRound(t *testing.T) {
	config := getServiceConfig()
	logger := log.New(getLogOutput(config.LogToFile), "", log.Ldate|log.Ltime|log.Lshortfile)

	msgs := []model.Message{
		{Type: messagetype.StartNewRound, Data: []byte("1")},
		{Type: messagetype.StartNewRound, Data: []byte("2")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("3")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("4")},
		{Type: messagetype.StartNewRound, Data: []byte("5")},
		{Type: messagetype.StartNewRound, Data: []byte("6")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("7")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("8")},
	}
	expectedOutputIds := [][]byte{[]byte("6"), []byte("5")}

	socket := networksocket.New(msgs)
	relayer := NewRelayer(socket, logger, config)

	sub1, sub1MsgChan := subscriber.New(
		logger,
		relayer,
		1,
		[]messagetype.MessageType{messagetype.StartNewRound},
		true,
	)

	sub1.Listen()
	relayer.Listen()

	// assert result messages
	i := 0
	for msg := range sub1MsgChan {

		if i >= len(expectedOutputIds) || bytes.Equal(msg.Data, expectedOutputIds[i]) == false {
			t.Fatalf("error on result assert, expected msg.data: %s, actual: %s", expectedOutputIds[i], msg.Data)
		}
		i++
	}

	// assert result length
	if i != len(expectedOutputIds) {
		t.Fatalf("error on result length assert, expected len: %d, actual: %d", len(expectedOutputIds), i)
	}

}

func TestTwoSubTwoMsgType(t *testing.T) {
	config := getServiceConfig()
	logger := log.New(getLogOutput(config.LogToFile), "", log.Ldate|log.Ltime|log.Lshortfile)

	msgs := []model.Message{
		{Type: messagetype.StartNewRound, Data: []byte("1")},
		{Type: messagetype.StartNewRound, Data: []byte("2")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("3")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("4")},
		{Type: messagetype.StartNewRound, Data: []byte("5")},
		{Type: messagetype.StartNewRound, Data: []byte("6")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("7")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("8")},
	}
	expectedOutput := [][]byte{[]byte("6"), []byte("5"), []byte("8")}

	socket := networksocket.New(msgs)
	relayer := NewRelayer(socket, logger, config)

	sub1, sub1MsgChan := subscriber.New(
		logger,
		relayer,
		1,
		[]messagetype.MessageType{messagetype.StartNewRound, messagetype.ReceivedAnswer},
		true,
	)

	sub2, sub2MsgChan := subscriber.New(
		logger,
		relayer,
		2,
		[]messagetype.MessageType{messagetype.StartNewRound, messagetype.ReceivedAnswer},
		true,
	)

	sub1.Listen()
	sub2.Listen()

	relayer.Listen()

	// assert result messages - sub 1
	i := 0
	for msg := range sub1MsgChan {
		if i >= len(expectedOutput) || bytes.Equal(msg.Data, expectedOutput[i]) == false {
			t.Fatalf("error on result assert, expected msg.data: %s, actual: %s", expectedOutput[i], msg.Data)
		}
		i++
	}

	// assert result length - sub 1
	if i != len(expectedOutput) {
		t.Fatalf("error on result length assert, expected len: %d, actual: %d", len(expectedOutput), i)
	}

	// assert result messages - sub 2
	i = 0
	for msg := range sub2MsgChan {
		if i >= len(expectedOutput) || bytes.Equal(msg.Data, expectedOutput[i]) == false {
			t.Fatalf("error on result assert, expected msg.data: %s, actual: %s", expectedOutput[i], msg.Data)
		}
		i++
	}

	// assert result length  - sub 2
	if i != len(expectedOutput) {
		t.Fatalf("error on result length assert, expected len: %d, actual: %d", len(expectedOutput), i)
	}

}

func TestTwoSubDifferentMsgType(t *testing.T) {
	config := getServiceConfig()
	logger := log.New(getLogOutput(config.LogToFile), "", log.Ldate|log.Ltime|log.Lshortfile)

	msgs := []model.Message{
		{Type: messagetype.StartNewRound, Data: []byte("1")},
		{Type: messagetype.StartNewRound, Data: []byte("2")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("3")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("4")},
		{Type: messagetype.StartNewRound, Data: []byte("5")},
		{Type: messagetype.StartNewRound, Data: []byte("6")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("7")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("8")},
	}
	sub1ExpectedOutput := [][]byte{[]byte("6"), []byte("5")}
	sub2ExpectedOutput := [][]byte{[]byte("8")}
	socket := networksocket.New(msgs)
	relayer := NewRelayer(socket, logger, config)

	sub1, sub1MsgChan := subscriber.New(
		logger,
		relayer,
		1,
		[]messagetype.MessageType{messagetype.StartNewRound},
		true,
	)

	sub2, sub2MsgChan := subscriber.New(
		logger,
		relayer,
		2,
		[]messagetype.MessageType{messagetype.ReceivedAnswer},
		true,
	)

	sub1.Listen()
	sub2.Listen()

	relayer.Listen()

	// assert result messages - sub 1
	i := 0
	for msg := range sub1MsgChan {
		if i >= len(sub1ExpectedOutput) || bytes.Equal(msg.Data, sub1ExpectedOutput[i]) == false {
			t.Fatalf("error on result assert, expected msg.data: %s, actual: %s", sub1ExpectedOutput[i], msg.Data)
		}
		i++
	}

	// assert result length - sub 1
	if i != len(sub1ExpectedOutput) {
		t.Fatalf("error on result length assert, expected len: %d, actual: %d", len(sub1ExpectedOutput), i)
	}

	// assert result messages - sub 2
	i = 0
	for msg := range sub2MsgChan {
		if i >= len(sub2ExpectedOutput) || bytes.Equal(msg.Data, sub2ExpectedOutput[i]) == false {
			t.Fatalf("error on result assert, expected msg.data: %s, actual: %s", sub2ExpectedOutput[i], msg.Data)
		}
		i++
	}

	// assert result length  - sub 2
	if i != len(sub2ExpectedOutput) {
		t.Fatalf("error on result length assert, expected len: %d, actual: %d", len(sub2ExpectedOutput), i)
	}

}

func TestOneBufferedAndOneNonBuffered(t *testing.T) {
	config := getServiceConfig()
	logger := log.New(getLogOutput(config.LogToFile), "", log.Ldate|log.Ltime|log.Lshortfile)

	msgs := []model.Message{
		{Type: messagetype.StartNewRound, Data: []byte("1")},
		{Type: messagetype.StartNewRound, Data: []byte("2")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("3")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("4")},
		{Type: messagetype.StartNewRound, Data: []byte("5")},
		{Type: messagetype.StartNewRound, Data: []byte("6")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("7")},
		{Type: messagetype.ReceivedAnswer, Data: []byte("8")},
	}
	sub1ExpectedOutput := [][]byte{[]byte("6")}
	sub2ExpectedOutput := [][]byte{[]byte("6"), []byte("5")}
	socket := networksocket.New(msgs)
	relayer := NewRelayer(socket, logger, config)

	sub1, sub1MsgChan := subscriber.New(
		logger,
		relayer,
		1,
		[]messagetype.MessageType{messagetype.StartNewRound},
		false,
	)

	sub2, sub2MsgChan := subscriber.New(
		logger,
		relayer,
		2,
		[]messagetype.MessageType{messagetype.StartNewRound},
		true,
	)

	sub1.Listen()
	sub2.Listen()

	relayer.Listen()

	// assert result messages - sub 1
	i := 0
	for msg := range sub1MsgChan {
		if i >= len(sub1ExpectedOutput) || bytes.Equal(msg.Data, sub1ExpectedOutput[i]) == false {
			t.Fatalf("error on result assert, expected msg.data: %s, actual: %s", sub1ExpectedOutput[i], msg.Data)
		}
		i++
	}

	// assert result length - sub 1
	if i != len(sub1ExpectedOutput) {
		t.Fatalf("error on result length assert, expected len: %d, actual: %d", len(sub1ExpectedOutput), i)
	}

	// assert result messages - sub 2
	i = 0
	for msg := range sub2MsgChan {
		if i >= len(sub2ExpectedOutput) || bytes.Equal(msg.Data, sub2ExpectedOutput[i]) == false {
			t.Fatalf("error on result assert, expected msg.data: %s, actual: %s", sub2ExpectedOutput[i], msg.Data)
		}
		i++
	}

	// assert result length  - sub 2
	if i != len(sub2ExpectedOutput) {
		t.Fatalf("error on result length assert, expected len: %d, actual: %d", len(sub2ExpectedOutput), i)
	}
}

func getServiceConfig() *configuration.Config {
	broadcastOrder := []messagetype.MessageType{messagetype.StartNewRound, messagetype.ReceivedAnswer}

	msgTypeToQueueSize := make(map[messagetype.MessageType]int)
	msgTypeToQueueSize[messagetype.StartNewRound] = 2
	msgTypeToQueueSize[messagetype.ReceivedAnswer] = 1

	return &configuration.Config{
		MsgTypeStoredLength:   msgTypeToQueueSize,
		MsgTypeBroadcastOrder: broadcastOrder,
		LogToFile:             false,
	}
}

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
