package iourt43

import (
	"testing"
	//"fmt"
)

func TestClearLine(t *testing.T) {
	s := "  1579:03 ConnectInfo: 0: E24F9B2702B9E4A1223E905BF597FA92: ^w[^2AS^w]^2Lead: 3: 3: 24.153.180.106:2794"
	ret := clearLine(s)
	if (ret != "ConnectInfo: 0: E24F9B2702B9E4A1223E905BF597FA92: ^w[^2AS^w]^2Lead: 3: 3: 24.153.180.106:2794") {
		t.Errorf("ParseLine took %s and return %s", s, ret)
	}
}

// func TestParser(t *testing.T) {
// 	s := "151:123 ClientUserinfo: 0 \\ip\\192.168.43.2:27961\\challenge\\-1284990544\\qport\\522\\protocol\\68\\snaps\\20\\name\\Esty\\rate\\25000\\sex\\male\\handicap\\100\\color2\\5\\color1\\4\\authc\\0\\cl_guid\\57221752F1988CC10DB254A791F54271"

// 	fmt.Print(ParseLine(s))
// }