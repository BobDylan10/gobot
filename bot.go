package main

import (
    "bufio"
    "fmt"
	"os"
	"time"
	"regexp"

	"os/signal"
	"syscall"

	"testbot/log"

	"testbot/database"

	"testbot/parsers/iourt43"
	"testbot/server"
	"testbot/events"

	"testbot/plugins/mappings"
	"testbot/plugins"

	"testbot/players"
)
//This takes a string, maps its named results into a map
func getParams(regEx, url string) (paramsMap map[string]string) {

    var compRegEx = regexp.MustCompile(regEx)
	match := compRegEx.FindStringSubmatch(url)
	if match == nil {
		fmt.Print("No match")
	}

    paramsMap = make(map[string]string)
    for i, name := range compRegEx.SubexpNames() {
        if i > 0 && i <= len(match) {
            paramsMap[name] = match[i]
        }
    }
    return
}

func reader(path string) {
	file, err := os.Open(path) //We should put the cursor at the end of file to avoid re-reading from scratch
	if err != nil {
		log.Log(log.LOG_ERROR, err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	read_rate := time.NewTicker(500 * time.Millisecond)

	for {
		select {
			case <-read_rate.C:
				scanner = bufio.NewScanner(file)
				//s, _ := file.Stat()
				//fmt.Println("Size:", s.Size())
				for scanner.Scan() {
					txt := scanner.Text()
					//fmt.Println("Parsing " + txt)
					e := iourt43.ParseLine(txt)
					if (e != nil) {
						pluginSchedule(e)
					}
				}
				//fmt.Println("End round")
				if err := scanner.Err(); err != nil {
					log.Log(log.LOG_ERROR, err.Error())
				}
		}
	}
}


func pluginSchedule(evt events.Event) {
	fmt.Println(evt)
	//First, send to players. This is not send in a goroutine because it is safer to wait for the player module to finish its work before continuing the rest
	players.CollectEvents(evt)
	//Then, send to plugins
	for plugin, evtmap := range pluginInBuffers {
		if (mappings.IsDep(plugin, evt.EventType())) {
			fmt.Println("Trying to send to plugin", plugin)
			evtmap<-evt
			select {
			case transmit := <-outEvents: //This will be done in another goroutine in the future ? It might be of use to use blocking channels. It might give some really weird behaviours, it might be of importance to monitor the execution time of the plugins
				fmt.Println("Should transmit event", transmit)
			default:
				fmt.Println("No event to transmit")
			}
		}
	}
}

//Not sure if this is really needed at this point. Inter plugin communication can be handled at the plugin layer directly
func restransmitEvents() {
	for {
		a := <-outEvents
		fmt.Println("Should retransmit", a)
	}
}

var outEvents chan plugins.PassEvent
var pluginInBuffers map[plugins.Plugin](chan events.Event)
var testMap map[plugins.Plugin]int

func newChan() (chan events.Event) {
	return make(chan events.Event)
}

func initPlugins() {
	//Start runners
	pluginInBuffers = make(map[plugins.Plugin](chan events.Event))
	outEvents = make(chan plugins.PassEvent, 2)
	testMap = make(map[plugins.Plugin]int)
	testMap[plugins.PLUGIN_CMD] = 3


	players.Init()
	for plugin, runner := range mappings.Runners {
		fmt.Println("Initializing", plugin)
		pluginInBuffers[plugin] = newChan()
		go runner(pluginInBuffers[plugin], outEvents) //TODO: call an Init function that return an inc channel that will be used to communicate with the plugin
	}
	fmt.Println("pluginInBuffer, ", pluginInBuffers)
	fmt.Println("testMap, ", testMap)

	//go restransmitEvents()
}

func main() {
	log.Log(log.LOG_INFO, "Starting b0t")
	path := "/home/guillaume/Documents/Urt/q3ut4/games.log"
	

	database.SetDatabase("gobot:gobot@/gobot_db?parseTime=true") //Init database
	initPlugins()
	server.Init()
	//server.Say("^5 Bot starting up")
	//test := make(map[int](chan string))
	go reader(path)
	//server.CallServer("kick 0")

	//Cleanup and signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		fmt.Println("Exiting bot")
		go database.CloseDatabase() //Launched in goroutines incase they are frozen for some reason
		go server.Close()

		time.Sleep(time.Second)//Let 1 second to close what needs to be closed

		os.Exit(0)
	}()

	done := make(chan bool)

	<-done //Wait infinity
}