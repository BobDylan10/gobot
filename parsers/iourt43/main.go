package iourt43

import (
    "regexp"
    "fmt"
    "strings"
    "testbot/events"
    "strconv"
)


// // remove the time off of the line
//     _lineTime = re.compile(`^(?P<minutes>[0-9]+):(?P<seconds>[0-9]+).*')
//     _lineClear = re.compile(`^(?:[0-9:]+\s?)?')
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

var _lineFormats = [] string {
    // Radio: 0 - 7 - 2 - "New Alley" - "I'm going for the flag"
    `(?i)^(?P<action>Radio): ` +
    `(?P<data>(?P<cid>[0-9]+) - ` +
    `(?P<msg_group>[0-9]+) - ` +
    `(?P<msg_id>[0-9]+) - ` +
    `"(?P<location>.*)" - ` +
    `"(?P<text>.*)")$`,

    // Callvote: 1 - "map dressingroom"
    `(?i)^(?P<action>Callvote): (?P<data>(?P<cid>[0-9]+) - "(?P<vote_string>.*)")$`,

    // Vote: 0 - 2
    `(?i)^(?P<action>Vote): (?P<data>(?P<cid>[0-9]+) - (?P<value>.*))$`,

    // VotePassed: 1 - 0 - "reload"
    `(?i)^(?P<action>VotePassed): (?P<data>(?P<yes>[0-9]+) - (?P<no>[0-9]+) - "(?P<what>.*)")$`,

    // VoteFailed: 1 - 1 - "restart"
    `(?i)^(?P<action>VoteFailed): (?P<data>(?P<yes>[0-9]+) - (?P<no>[0-9]+) - "(?P<what>.*)")$`,

    // FlagCaptureTime: 0: 1234567890
    // FlagCaptureTime: 1: 1125480101
    `(?i)^(?P<action>FlagCaptureTime):\s(?P<cid>[0-9]+):\s(?P<captime>[0-9]+)$`,

    // 13:34 ClientJumpRunStarted: 0 - way: 1
    // 13:34 ClientJumpRunStarted: 0 - way: 1 - attempt: 1 of 5
    `(?i)^(?P<action>ClientJumpRunStarted):\s` +
        `(?P<cid>\d+)\s-\s` +
        `(?P<data>way:\s` +
        `(?P<way_id>\d+)` +
        `(?:\s-\sattempt:\s` +
        `(?P<attempt_num>\d+)\sof\s` +
        `(?P<attempt_max>\d+))?)$`,

    // 13:34 ClientJumpRunStopped: 0 - way: 1 - time: 12345
    // 13:34 ClientJumpRunStopped: 0 - way: 1 - time: 12345 - attempt: 1 of 5
    `(?i)^(?P<action>ClientJumpRunStopped):\s` +
        `(?P<cid>\d+)\s-\s` +
        `(?P<data>way:\s` +
        `(?P<way_id>\d+)` +
        `\s-\stime:\s` +
        `(?P<way_time>\d+)` +
        `(?:\s-\sattempt:\s` +
        `(?P<attempt_num>\d+)\sof\s` +
        `(?P<attempt_max>\d+` +
        `))?)$`,

    // 13:34 ClientJumpRunCanceled: 0 - way: 1
    // 13:34 ClientJumpRunCanceled: 0 - way: 1 - attempt: 1 of 5
    `(?i)^(?P<action>ClientJumpRunCanceled):\s` +
        `(?P<cid>\d+)\s-\s` +
        `(?P<data>way:\s` +
        `(?P<way_id>\d+)` +
        `(?:\s-\sattempt:\s` +
        `(?P<attempt_num>\d+)\sof\s` +
        `(?P<attempt_max>\d+))?)$`,

    // 13:34 ClientSavePosition: 0 - 335.384887 - 67.469154 - -23.875000
    // 13:34 ClientLoadPosition: 0 - 335.384887 - 67.469154 - -23.875000
    `(?i)^(?P<action>Client(Save|Load)Position):\s` +
        `(?P<cid>\d+)\s-\s` +
        `(?P<data>` +
        `(?P<x>-?\d+(?:\.\d+)?)\s-\s` +
        `(?P<y>-?\d+(?:\.\d+)?)\s-\s` +
        `(?P<z>-?\d+(?:\.\d+)?))$`,

    // 13:34 ClientGoto: 0 - 1 - 335.384887 - 67.469154 - -23.875000
    `(?i)^(?P<action>ClientGoto):\s` +
        `(?P<cid>\d+)\s-\s` +
        `(?P<tcid>\d+)\s-\s` +
        `(?P<data>` +
        `(?P<x>-?\d+(?:\.\d+)?)\s-\s` +
        `(?P<y>-?\d+(?:\.\d+)?)\s-\s` +
        `(?P<z>-?\d+(?:\.\d+)?))$`,

    // ClientSpawn: 0
    // ClientMelted: 1
    `(?i)^(?P<action>Client(Melted|Spawn)):\s(?P<cid>[0-9]+)$`,

    //Assist: 0 14 15: -[TPF]-PtitBigorneau assisted Bot1 to kill Bot2
    `(?i)^(?P<action>Assist):\s(?P<acid>[0-9]+)\s(?P<kcid>[0-9]+)\s(?P<dcid>[0-9]+):\s+(?P<text>.*)$`,

    // Generated with ioUrbanTerror v4.1:
    // Hit: 12 7 1 19: BSTHanzo[FR] hit ercan in the Helmet
    // Hit: 13 10 0 8: Grover hit jacobdk92 in the Head
    `(?i)^(?P<action>Hit):\s` +
        `(?P<data>` +
        `(?P<cid>[0-9]+)\s` +
        `(?P<acid>[0-9]+)\s` +
        `(?P<hitloc>[0-9]+)\s` +
        `(?P<aweap>[0-9]+):\s+` +
        `(?P<text>.*))$`,

    // 6:37 Kill: 0 1 16: XLR8or killed =lvl1=Cheetah by UT_MOD_SPAS
    // 2:56 Kill: 14 4 21: Qst killed Leftovercrack by UT_MOD_PSG1
    // 6:37 Freeze: 0 1 16: Fenix froze Biddle by UT_MOD_SPAS
    `(?i)^(?P<action>[a-z]+):\s` +
        `(?P<data>` +
        `(?P<acid>[0-9]+)\s` +
        `(?P<cid>[0-9]+)\s` +
        `(?P<aweap>[0-9]+):\s+` +
        `(?P<text>.*))$`,

    // ThawOutStarted: 0 1: Fenix started thawing out Biddle
    // ThawOutFinished: 0 1: Fenix thawed out Biddle
    `(?i)^(?P<action>ThawOut(Started|Finished)):\s` +
        `(?P<data>` +
        `(?P<cid>[0-9]+)\s` +
        `(?P<tcid>[0-9]+):\s+` +
        `(?P<text>.*))$`,

    // Processing chats and tell events...
    // 5:39 saytell: 15 16 repelSteeltje: nno
    // 5:39 saytell: 15 15 repelSteeltje: nno
    `(?i)^(?P<action>[a-z]+):\s` +
        `(?P<data>` +
        `(?P<cid>[0-9]+)\s` +
        `(?P<acid>[0-9]+)\s` +
        `(?P<name>.+?):\s+` +
        `(?P<text>.*))$`,

    // SGT: fix issue with onSay when something like this come and the match could` +nt find the name group
    // say: 7 -crespino-:
    // say: 6 ^5Marcel ^2[^6CZARMY^2]: !help
    `(?i)^(?P<action>[a-z]+):\s` +
        `(?P<data>` +
        `(?P<cid>[0-9]+)\s` +
        `(?P<name>[^ ]+):\s*` +
        `(?P<text>.*))$`,

    // 15:42 Flag Return: RED
    // 15:42 Flag Return: BLUE
    `(?i)^(?P<action>Flag Return):\s(?P<data>(?P<color>.+))$`,

    // Bombmode actions:
    // 3:06 Bombholder is 2
    `(?i)^(?P<action>Bombholder)(?P<data>\sis\s(?P<cid>[0-9]))$`,

    // was planted, was defused, was tossed, has been collected (doh, how gramatically correct!)
    // 2:13 Bomb was tossed by 2
    // 2:32 Bomb was planted by 2
    // 3:01 Bomb was defused by 3!
    // 2:17 Bomb has been collected by 2
    `(?i)^(?P<action>Bomb)\s` +
        `(?P<data>(was|has been)\s` +
        `(?P<subaction>[a-z]+)\sby\s` +
        `(?P<cid>[0-9]+).*)$`,

    // 17:24 Pop!
    `(?i)^(?P<action>Pop)!$`,

    // Falling thru? Item stuff and so forth
    `(?i)^(?P<action>[a-z]+):\s(?P<data>.*)$`,

    // Shutdowngame and Warmup... the one word lines
    `(?i)^(?P<action>[a-z]+):$`}

func clearLine(line string) string {
    line = strings.TrimSpace(line)

	lineclear := `^(?:[0-9:]+\s?)?`

	var compRegEx = regexp.MustCompile(lineclear)

	return compRegEx.ReplaceAllLiteralString(line, "")
}

func getLineParts(line string) (map[string] string, error) {
	line = clearLine(line)
    fmt.Println(line)
	for _, v := range _lineFormats {
		form := regexp.MustCompile(v)
		if (form.MatchString(line)) {
			return getParams(form, line), nil
		}
	}
	return nil, &parseError{line}
}

func createEvent(args map[string] string) *events.Event {
    action := strings.ToLower(args["action"])
    switch action {
    case "say":
        cid, err := strconv.Atoi(args["cid"])
        if err != nil {
            fmt.Println("Error converting cid")
        }
        return &events.Event{Evt: events.EVT_CLIENT_SAY, Client: cid, Data: map[string]string{"text": args["text"]}} //Might need the name as well
    default:
        return events.NewEvent(events.EVT_CLIENT_SAY)
    }
}

func ParseLine(line string) *events.Event {
    r, e := getLineParts(line)
    if (e != nil) {
        fmt.Printf("Error parsing line %s", line)
    }
    fmt.Println(r)
	return createEvent(r)
}