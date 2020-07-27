package main

import (
	"bufio"
	"flag"
	"os"
	"time"

	"os/signal"
	"syscall"

	"testbot/log"

	"testbot/database"

	"testbot/events"
	"testbot/parsers/iourt43"
	"testbot/server"

	"testbot/plugins"

	"testbot/players"
)

func reader(path string) {
	file, err := os.Open(path) //We should put the cursor at the end of file to avoid re-reading from scratch
	if err != nil {
		log.Log(log.LOG_ERROR, err.Error())
	}
	defer file.Close()

	file.Seek(0, 2) //Read from the end of file

	scanner := bufio.NewScanner(file)

	read_rate := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-read_rate.C:
			scanner = bufio.NewScanner(file)
			for scanner.Scan() {
				txt := scanner.Text()
				e := iourt43.ParseLine(txt)
				if e != nil {
					pluginSchedule(e)
				}
			}
			if err := scanner.Err(); err != nil {
				log.Log(log.LOG_ERROR, err.Error())
			}
		}
	}
}

func pluginSchedule(evt events.Event) {
	log.Log(log.LOG_DEBUG, "Scheduling plugins for event", evt)
	//First, send to players. This is not sent in a goroutine because it is safer to wait for the player module to finish its work before continuing the rest
	players.CollectEvents(evt)
	//Then, send to plugins
	for plugin, evtmap := range pluginInBuffers {
		//TODO: build the list of plugins to be scheduled per event to avoid computation
		if plugins.Plugins[plugin].IsDep(evt.EventType()) {
			//TODO: Here we should add a timeout incase a plugin is stuck
			evtmap <- evt
		}
	}
	//TODO: wait for all plugins to be done ? No, call directly plugin functions
}

var pluginInBuffers map[plugins.PluginID](chan<- events.Event)

func initPlugins() {
	//Start runners
	pluginInBuffers = make(map[plugins.PluginID](chan<- events.Event))

	players.Init()
	for _, pluginid := range plugins.GetPluginOrder() {
		log.Log(log.LOG_INFO, "Initializing plugin", plugins.Plugins[pluginid].GetName())
		pluginInBuffers[pluginid] = plugins.Plugins[pluginid].Init()
		time.Sleep(100 * time.Millisecond)
	}
}

func argsParse() {
	cfgPath := flag.String("cfg", "config/config.json", "The path to the config file")
	flag.Parse()
	log.Log(log.LOG_INFO, "Config file at", *cfgPath)
}

func main() {
	log.Log(log.LOG_INFO, "Starting b0t")
	argsParse()
	path := "/home/guillaume/Documents/Urt/q3ut4/games.log"

	database.SetDatabase("gobot:gobot@/gobot_db?parseTime=true") //Init database
	initPlugins()
	server.Init()
	server.Say("^5 GoBot starting up")
	go reader(path)

	//Cleanup and signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signals
		log.Log(log.LOG_INFO, "Exiting bot")
		go database.CloseDatabase() //Launched in goroutines incase they are frozen for some reason
		go server.Close()

		time.Sleep(time.Second) //Let 1 second to close what needs to be closed

		os.Exit(0)
	}()

	server.Reload() //Reload server so that everything works correctly

	done := make(chan bool)

	<-done //Wait infinity
}
