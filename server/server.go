package server

import (
	"testbot/server/rcon"
	"regexp"
	"strings"
)

var (
	done = make(chan bool)
	cmds = make(chan string)
	aws = make(chan string)
)

func Init() {
	go rcon.RconRunner(done, cmds, aws)
}

func Close() {
	<-done
}

func CallServer(command string) string{
	cmds<-command
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
		if (i >= 3) {
			if (reg.MatchString(line)) {
				getParams(reg, line)
			}
		}
	}
}

func f2(command string) string {
	return "b"
}