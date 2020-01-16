package rcon

//TODO: Better handling of network failures

import (
	"net"
	"strings"
	"time"
	"fmt"

	"testbot/log"
)
const rconsendstring = "\xff\xff\xff\xffrcon \"%s\" %s\n"
const rconreplystring = "\xff\xff\xff\xffprint\n"

var password = "abc"

const maxBufferSize = 4096

func RconRunner(done <-chan bool, commands <-chan string, answers chan<- string) {
	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:27960")
	if err != nil {
		return
	}
	conn, err := net.DialUDP("udp", nil, raddr)
	defer conn.Close()

	if err != nil {
		log.Log(log.LOG_ERROR, "Error", err)
	}
	for {
		select {
		case cmd := <-commands:
			log.Log(log.LOG_DEBUG, "Command: ", cmd)
			conn.Write([]byte(fmt.Sprintf(rconsendstring, password, cmd)))
			buffer := make([]byte, maxBufferSize)
			n, err := conn.Read(buffer)
			if err != nil {
				log.Log(log.LOG_ERROR, "Error", err)
				answers<-""
				continue
			}
			ans := string(buffer[0:n])
			ans = strings.TrimPrefix(ans, rconreplystring)
			answers<-ans
			//fmt.Println(ans)
			time.Sleep(250 * time.Millisecond)
		case <-done:
			return
		}
	}
}