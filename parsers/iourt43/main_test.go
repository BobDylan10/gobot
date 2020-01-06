package iourt43

import (
	"testing"
	"fmt"
)

func TestClearLine(t *testing.T) {
	s := "  1579:03 ConnectInfo: 0: E24F9B2702B9E4A1223E905BF597FA92: ^w[^2AS^w]^2Lead: 3: 3: 24.153.180.106:2794"
	ret := clearLine(s)
	if (ret != "ConnectInfo: 0: E24F9B2702B9E4A1223E905BF597FA92: ^w[^2AS^w]^2Lead: 3: 3: 24.153.180.106:2794") {
		t.Errorf("ParseLine took %s and return %s", s, ret)
	}
}

func TestParseUserInfo(t *testing.T) {
	//2 \ip\145.99.135.227:27960\challenge\-232198920\qport\2781\protocol\68\battleye\1\name\[SNT]^1XLR^78or...
	// 7 n\[SNT]^1XLR^78or\t\3\r\2\tl\0\f0\\f1\\f2\\a0\0\a1\0\a2\0

	//m1 := map[string]string{"battleye": "1", "challenge": "-232198920", "ip":"145.99.135.227:27960", "protocol":"68", "qport":"2781"}
	//m2 := map[string]string{"n":"[SNT]^1XLR^78or", "a0":"0", "a1":"0", "a2":"0", "r":"2", "t":"3", "tl":"0"}
	s1 := "\\ip\\145.99.135.227:27960\\challenge\\-232198920\\qport\\2781\\protocol\\68\\battleye\\1"
	s2 := "n\\[SNT]^1XLR^78or\\t\\3\\r\\2\\tl\\0\\f0\\\\f1\\\\f2\\\\a0\\0\\a1\\0\\a2\\0"
	fmt.Println(parseUserInfo(s1))
	fmt.Println(parseUserInfo(s2))
}

// func TestParser(t *testing.T) {
// 	s := "151:123 ClientUserinfo: 0 \\ip\\192.168.43.2:27961\\challenge\\-1284990544\\qport\\522\\protocol\\68\\snaps\\20\\name\\Esty\\rate\\25000\\sex\\male\\handicap\\100\\color2\\5\\color1\\4\\authc\\0\\cl_guid\\57221752F1988CC10DB254A791F54271"

// 	fmt.Print(ParseLine(s))
// }