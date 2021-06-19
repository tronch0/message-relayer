package messagetype

type MessageType int

const (
	StartNewRound MessageType = 1 << iota
	ReceivedAnswer
	Undefined	// added default value
)

