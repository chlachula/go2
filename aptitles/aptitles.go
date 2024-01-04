package aptitles

import (
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`(?m).*"ap(?P<yymmdd>\d{6})\.html">(?P<title>.*)<\/a>.*`)

var tmpfile = `2021 August 02:  <a href="ap210802.html">The Hubble Ultra Deep Field in Light and Sound</a><br>
2021 August 01:  <a href="ap210801.html">Pluto in Enhanced Color</a><br>
2021 July 31:  <a href="ap210731.html">Remembering NEOWISE</a><br>
2021 July 30:  <a href="ap210730.html">Mimas in Saturnlight</a><br>
2021 July 29:  <a href="ap210729.html">The Tulip and Cygnus X 1</a><br>`

func LoadAPODarchive() error {
	lines := strings.Split(tmpfile, "\n")
	fmt.Printf(`"titles": [%s`, "\n")
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 2 {
			fmt.Printf(` {"ymd":"%s", "title":"%s"},%s`, match[1], match[2], "\n")
		}
	}
	fmt.Printf(`]%s`, "\n")
	return nil
}

func SearchTitle(yymmdd string) (string, error) {
	if len(yymmdd) != 6 {
		return "", fmt.Errorf("yymmdd: %s does not have length 6", yymmdd)
	}
	return "OK", nil
}
