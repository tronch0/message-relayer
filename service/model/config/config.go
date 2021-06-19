package config

import (
	messagetype2 "message-relayer/service/model/messagetype"
)

type Config struct {
	MessageTypeToQueueSize map[messagetype2.MessageType]int
	MessageTypeImportanceOrderDesc []messagetype2.MessageType
}
