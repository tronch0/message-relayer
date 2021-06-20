package config

import (
	messagetype2 "message-relayer/service/model/messagetype"
)

type Config struct {
	MsgTypeStoredLength        map[messagetype2.MessageType]int
	MsgTypeImportanceOrderDesc []messagetype2.MessageType
}
