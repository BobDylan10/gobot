package rcon

import (
	"net"
	"fmt"
	"strings"
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
        fmt.Printf("Some error %v", err)
        return
	}
	for {
		select {
		case cmd := <-commands:
			conn.Write([]byte(fmt.Sprintf(rconsendstring, password, cmd)))
			buffer := make([]byte, maxBufferSize)
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Printf("Some error %v", err)
				return
			}
			ans := string(buffer[0:n])
			ans = strings.TrimPrefix(ans, rconreplystring)
			answers<-ans
			//fmt.Println(ans)
		case <-done:
			return
		}
	}
}