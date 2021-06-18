package model

import "message-relayer/model/messagetype"

type Config struct {
	MessageTypeToQueueSize map[messagetype.MessageType]int
	MessageTypeImportanceOrderDesc []messagetype.MessageType
}
