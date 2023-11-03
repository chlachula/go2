package fileDateMatch

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func fileToLines(fname string) ([]string, error) {
	lines := make([]string, 0)
	bytes, err := os.ReadFile(fname)
	if err != nil {
		return lines, err
	}

	text := string(bytes)
	return strings.Split(text, "\n"), nil
}
func dateInsideFile(lines []string) (time.Time, error) {
	d := time.Now()
	// loop
	return d, nil
}

func getParams(compRegEx *regexp.Regexp, str string) (map[string]string, error) {
	paramsMap := make(map[string]string)
	if !compRegEx.MatchString(str) {
		return paramsMap, fmt.Errorf("no match for string '%s'", str)
	}
	match := compRegEx.FindStringSubmatch(str)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap, nil
}
func dateFromFilename(filename string) (time.Time, error) {
	var re = regexp.MustCompile(`(?m)(?P<yy>\d\d)(?P<mm>\d\d)(?P<dd>\d\d)`)
	d := time.Now()
	m, err := getParams(re, filename)
	if err != nil {
		return d, err
	}
	yymmdd := m["yy"] + m["mm"] + m["dd"]
	d, _ = time.Parse("060102", yymmdd)
	return d, nil
}

func FilenameDateMatch(filename string) error {
	if _, err := os.Stat(filename); err != nil {
		return fmt.Errorf("file %s does not exist", filename)
	}
	if lines, err := fileToLines(filename); err != nil {
		return err
	} else {
		if dateInside, err2 := dateInsideFile(lines); err != nil {
			return err2
		} else {
			if dateFilename, err3 := dateFromFilename(filename); err3 != nil {
				return err3
			} else {
				if !dateInside.Equal(dateFilename) {
					return fmt.Errorf("filename %s does not match date inside", filename)
				} else {
					return nil
				}

			}
		}
	}
}
