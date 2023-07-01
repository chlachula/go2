package file_YMD_match

import (
	"fmt"
	"os"
	"regexp"
)

var re1 = regexp.MustCompile(`(?m)(?P<src>.+)\/(?P<yyyy>\d{4})_(?P<mm>\d\d)_(?P<dd>\d\d)(?P<x>.*)\/(?P<filename>.+)`)

func getParams(compRegEx *regexp.Regexp, str string) (paramsMap map[string]string) {
	paramsMap = make(map[string]string)

	match := compRegEx.FindStringSubmatch(str)

	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}

// returns true if sourceDir/yyyy_mm_dd/filename exists
// at      targetDir/yyyy/mm/yyyy_mm_dd/filename
func ExistsYMD(sourcePathname, targetDir string) bool {
	if re1.Match([]byte(sourcePathname)) {
		p := getParams(re1, sourcePathname)
		yyyy := p["yyyy"]
		mm := p["mm"]
		dd := p["dd"]
		x := p["x"]
		filename := p["filename"]
		target := fmt.Sprintf("%s/%s/%s/%s_%s_%s%s/%s", targetDir, yyyy, mm, yyyy, mm, dd, x, filename)
		fmt.Println("DEBUG target = ", target)
		if _, err := os.Stat(target); err == nil {
			return true
		} else {
			return false
		}
	}
	return false
}
