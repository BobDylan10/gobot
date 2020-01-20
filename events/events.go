package events

//TODO: Replace all the clients with real players objects

//Describes the type of the event
type EventType int

const (
	EVT_CLIENT_BEGIN EventType = iota
	EVT_CLIENT_SAY
	EVT_CLIENT_DISCONNECT
	EVT_CLIENT_INFO
	EVT_CLIENT_KILL
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

type EventClientDisconnect struct {
		Client int
	}
	func (e EventClientDisconnect) EventType() EventType {
		return EVT_CLIENT_DISCONNECT
	}