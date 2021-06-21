package config

import (
	messagetype2 "message-relayer/service/model/messagetype"
)

type Config struct {
	MsgTypeStoredLength   map[messagetype2.MessageType]int
	MsgTypeBroadcastOrder []messagetype2.MessageType
	LogToFile             bool
}
