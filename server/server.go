package server

import (
	"regexp"
	"strconv"
	"strings"
	"testbot/log"
	"testbot/server/rcon"
)

var (
	done = make(chan bool)
	cmds = make(chan string)
	aws  = make(chan string)
)

const lineLength = 80

func Init() {
	go rcon.RconRunner(done, cmds, aws)
}

func Close() {
	<-done
}

func CallServer(command string) string {
	log.Log(log.LOG_DEBUG, "Executing command", command)
	cmds <- command
	return <-aws
}

// num score ping guid   name            lastmsg address               qport rate
// --- ----- ---- ------ --------------- ------- --------------------- ----- -----
// 2     0   29 465030   ThorN                50 68.63.6.62:-32085      6597  5000
const _regPlayer = `(?i)^(?P<slot>[0-9]+)\s+` +
	`(?P<score>[0-9-]+)\s+` +
	`(?P<ping>[0-9]+)\s+` +
	`(?P<guid>[0-9a-zA-Z]+)\s+` +
	`(?P<name>.*?)\s+` +
	`(?P<last>[0-9]+)\s+` +
	`(?P<ip>[0-9.]+):(?P<port>[0-9-]+)\s+` +
	`(?P<qport>[0-9]+)` +
	`\s+(?P<rate>[0-9]+)$`

func getParams(regEx *regexp.Regexp, url string) (paramsMap map[string]string) {
	match := regEx.FindStringSubmatch(url)

	paramsMap = make(map[string]string)
	for i, name := range regEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return
}

func getPlayers() { //Quickfix for Urt, should be done inside rcon
	res := CallServer("status")
	lines := strings.Split(res, "\n")
	reg := regexp.MustCompile(_regPlayer)
	for i, line := range lines {
		if i >= 3 {
			if reg.MatchString(line) {
				getParams(reg, line)
			}
		}
	}
}

func wrap(longstring string) []string {
	res := make([]string, 0)
	parts := len(longstring) / lineLength
	i := 0
	for i = 0; i < parts; i++ {
		res = append(res, longstring[lineLength*i:lineLength*(i+1)])
	}
	rest := len(longstring) % lineLength
	if rest != 0 {
		res = append(res, longstring[lineLength*i:lineLength*i+rest])
	}
	return res
}

func Say(tosay string) {
	log.Log(log.LOG_VERBOSE, "Saying", tosay)
	lines := wrap(tosay)
	for _, line := range lines {
		CallServer("say \"" + line + "\"")
	}
	//Is it necessary to return something ?
}

func Kick(pid int) bool {
	p := strconv.Itoa(pid)
	r := CallServer("kick " + p)
	b := strings.Contains(r, "kicked")
	if b {
		log.Log(log.LOG_INFO, "Player "+p+" was kicked")
		return true
	} else {
		log.Log(log.LOG_INFO, "Player "+p+" was not kicked")
		return false
	}
	//TODO: return code to say if it worked
}

func BigText(tosay string) {
	log.Log(log.LOG_VERBOSE, "Bigtexting", tosay)
	CallServer("bigtext \"Bonjour " + tosay + "\"")
	//Is it necessary to return something ?
}

func GetMaps() []string {
	log.Log(log.LOG_VERBOSE, "Getting maps")
	maps := CallServer("fdir *.bsp")
	ret := strings.Split(maps, "\n")
	return ret
}

func Reload() {
	log.Log(log.LOG_INFO, "Reloading server.")
	CallServer("reload")
}
