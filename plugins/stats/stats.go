package stats

import (
	"testbot/events"
	"testbot/players"

	"testbot/log"
	"time"
)


type plugininside struct{}

var Plug plugininside = plugininside{}

var connectionTimes map[int]time.Time

func (p plugininside) Init() chan<- events.Event{
	log.Log(log.LOG_INFO, "Starting plugin Stats")

	connectionTimes = make(map[int]time.Time)

	in := make(chan events.Event)
	go runner(in)
	return in
}


func runner(evts <-chan events.Event) {
	for {
		evt := <-evts
		switch t := evt.(type) {
		case events.EventClientInfo:
			if _, ok := connectionTimes[t.Client]; !ok {
				pl, ok := players.GetPlayer(t.Client)
				if ok && !pl.IsBot(){
					connectionTimes[t.Client] = time.Now()
				}
			}
		case events.EventClientDisconnect:
			if connectTime, ok := connectionTimes[t.Client]; ok {
				duration := time.Now().Sub(connectTime).Minutes()
				if (duration > 0) { //TODO: set this to a correct threshold (like 10 minutes with config)
					newTimeSpan(t.Client, duration)
				}
				delete(connectionTimes, t.Client)
			}
		default:
			log.Log(log.LOG_INFO, "Very weird that we are here, type is ", t)
		}
	}
}

var deps []events.EventType = []events.EventType{events.EVT_CLIENT_INFO, events.EVT_CLIENT_DISCONNECT}

func (p plugininside) IsDep(e events.EventType) bool{
	for _, v := range deps {
		if (v == e) {
			return true
		}
	}
	return false
}

func (p plugininside) GetName() string {
	return "stats"
}