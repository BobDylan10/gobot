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
    data := args["data"]
    switch action {
    case "say":
        cid, err := strconv.Atoi(args["cid"])
        if err != nil {
            fmt.Println("Error converting cid")
        }
        return &events.Event{Evt: events.EVT_CLIENT_SAY, Client: cid, Data: map[string]string{"text": args["text"]}} //Might need the name as well
    case "clientuserinfo":
        tmp := strings.SplitAfterN(data, " ", 2)
        cid := tmp[0]
        data = tmp[1]
        pid, err := strconv.Atoi(cid)
        if err != nil {
            fmt.Println("Error converting cid")
        }
        return &events.Event{Evt: events.EVT_CLIENT_INFO, Client: pid, Data: map[string]string{"data": data}}
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