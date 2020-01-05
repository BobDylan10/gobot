package events

//Describes the type of the event
type EventType int

const (
	EVT_CLIENT_BEGIN EventType = iota
	EVT_CLIENT_KILL
	EVT_CLIENT_SAY
)

// //Indices for Event elements
// type EventIndex int

// const (
// 	EVT_IND_CLIENT EventType = iota
// 	EVT_IND_TARGET
// 	EVT_IND_TEXT
//  EVT_IND_DICT
// )

type Event struct {
	Evt EventType
	Client int
	Target int
	Data map[string] string
}

func NewEvent(tpe EventType) *Event {
	n := Event{Evt: tpe}

	return &n
}