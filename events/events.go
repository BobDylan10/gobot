package events

//Describes the type of the event
type EventType int

const (
	EVT_CLIENT_BEGIN EventType = iota
	EVT_CLIENT_KILL
	EVT_CLIENT_SAY
	EVT_CLIENT_INFO
	EVT_CLIENT_COMMAND
)

//Indices for Event elements
type EventIndex int

const (
	EVT_IND_CLIENT EventType = iota
	EVT_IND_TARGET
	EVT_IND_TEXT
 EVT_IND_DICT
)

type Event interface{
	EventType() EventType
}

type EventSay struct {
	Client int
	Text string
}
func (e EventSay) EventType() EventType {
	return EVT_CLIENT_SAY
}

type EventClientInfo struct {
	Client int
	Data map[string] string
}
func (e EventClientInfo) EventType() EventType {
	return EVT_CLIENT_INFO
}


type EventCommand struct {
	Client int
	Command string
	Args string
}
func (e EventCommand) EventType() EventType {
	return EVT_CLIENT_COMMAND
}