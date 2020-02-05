package iourt43

import (
	"regexp"
)

func parseInfo (data string) map[string]string {
	if (data[0] != '\\') {
		data = "\\" + data
	}

	reg := regexp.MustCompile(`\\([^\\]+)\\([^\\]+)`)
	dic := reg.FindAllStringSubmatch(data, -1)

	uInfo := make(map[string]string)

	for _, v := range dic {
		uInfo[v[1]] = v[2]
	}

	return uInfo
}