package naming

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	patternLast2Numbers = regexp.MustCompile(`\.([\d]{2})$`)
)

func getNumberSplit(s string) int {
	w := strings.Split(s, `.`)
	if len(w) < 2 {
		return 0
	}
	outString := w[len(w)-1]
	r, err := strconv.ParseInt(outString, 10, 32)
	if err != nil {
		return 0
	}
	return int(r)
}

func getNumberRegexp(s string) int {
	d := patternLast2Numbers.FindStringSubmatch(s)
	if d == nil {
		return 0
	}
	r, err := strconv.ParseInt(d[1], 10, 32)
	if err != nil {
		return 0
	}
	return int(r)
}
