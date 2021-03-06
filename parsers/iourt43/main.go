package iourt43

import (
    "regexp"
    "fmt"
    "strings"
    "testbot/events"
    "strconv"

    "testbot/log"
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


func clearLine(line string) string {
    line = strings.TrimSpace(line)

	lineclear := `^(?:[0-9:]+\s?)?`

	var compRegEx = regexp.MustCompile(lineclear)

	return compRegEx.ReplaceAllLiteralString(line, "")
}

func getLineParts(line string) (map[string] string, error) {
	line = clearLine(line)
	for _, v := range _lineFormats {
		form := regexp.MustCompile(v)
		if (form.MatchString(line)) {
			return getParams(form, line), nil
		}
	}
	return nil, &parseError{line}
}

func createEvent(args map[string] string) events.Event {
    log.Log(log.LOG_DEBUG, "Line to map gave", args)
    action := strings.ToLower(args["action"])
    data := args["data"]
    switch action {
    case "say":
        cid, err := strconv.Atoi(args["cid"])
        if err != nil {
            fmt.Println("Error converting cid")
        }
        return events.EventSay{Client: cid, Text: args["text"]} //Might need the name as well
    case "clientuserinfo":
        tmp := strings.SplitN(data, " ", 2)
        cid := tmp[0]
        data = tmp[1]
        pid, err := strconv.Atoi(cid)
        if err != nil {
            log.Log(log.LOG_ERROR, "Error converting cid", cid)
        }
        return events.EventClientInfo{Client: pid, Data: parseInfo(data)}
    case "clientdisconnect":
        cid := data
        pid, err := strconv.Atoi(cid)
        if err != nil {
            log.Log(log.LOG_ERROR, "Error converting cid", cid)
        }
        return events.EventClientDisconnect{Client: pid}
    case "kill":
        killer, errk := strconv.Atoi(args["acid"])
        victim, errv := strconv.Atoi(args["cid"])
        weapon, errw := strconv.Atoi(args["aweap"])
        if errk != nil || errv != nil || errw != nil {
            log.Log(log.LOG_ERROR, "Error converting kill line integers")
        }
        return events.EventClientKill{Killer: killer, Victim: victim, Weapon:weapon}
    case "initgame":
        info := parseInfo(data)
        cmap := info["mapname"]
        delete(info, "mapname")
        info["map"] = cmap
        return events.EventInitGame{Data: info}
    default:
        return nil
    }
}

func ParseLine(line string) events.Event {
    r, e := getLineParts(line)
    if (e != nil) {
        log.Log(log.LOG_ERROR, "Error parsing line", line)
    }
    //fmt.Println(r)
	return createEvent(r)
}