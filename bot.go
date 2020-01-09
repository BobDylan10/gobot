package main

import (
    "bufio"
    "fmt"
    "log"
	"os"
	"time"
	"regexp"
	"testbot/parsers/iourt43"
	"testbot/server"
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
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	read_rate := time.NewTicker(500 * time.Millisecond)

	for {
		select {
			case <-read_rate.C:
				scanner = bufio.NewScanner(file)
				s, _ := file.Stat()
				fmt.Println("Size:", s.Size())
				for scanner.Scan() {
					txt := scanner.Text()
					fmt.Println("Parsing " + txt)
					fmt.Println(iourt43.ParseLine(txt))
				}
				fmt.Println("End round")
				if err := scanner.Err(); err != nil {
					log.Fatal(err)
				}
		}
	}
}

func main() {
	path := "/home/guillaume/Documents/Urt/q3ut4/games.log"
	server.Init()
	server.CallServer("say \"^2Reader starting up\"")

	go reader(path)
	server.CallServer("kick 0")

	done := make(chan bool)

	<-done //Wait infinity
}