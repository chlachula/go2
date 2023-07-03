package file_YMD_match

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	/* "github.com/udhos/equalfile" */)

var re1 = regexp.MustCompile(`(?m)(?P<src>.+)\/(?P<yyyy>\d{4})_(?P<mm>\d\d)_(?P<dd>\d\d)(?P<x>.*)\/(?P<filename>.+)`)

type RecordType = struct {
	md5   string
	path  string
	Path2 string
}

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

func findMatch(t RecordType, processed []RecordType) (bool, RecordType) {
	for _, p := range processed {
		if p.md5 == t.md5 {
			return true, p
		}
	}
	return false, t
}

// to find file matches
// find . -type f -exec md5sum {} \; >> toProcess.md5sums #or processes.md5sums
func FindMatches(toProcess, processed []RecordType) {
	f := 0
	for it, t := range toProcess {
		if found, r := findMatch(t, processed); found {
			toProcess[it].Path2 = r.path
			f += 1
		}
	}
	var a, b int
	for _, t := range toProcess {
		if t.Path2 != "" {
			a += 1
		} else {
			fmt.Printf("%v\n", t)
			b += 1
		}
	}
	fmt.Printf("found= %d, a=%d + b=%d = %d \n", f, a, b, len(toProcess))

}
func SumsToArr(filePath string) []RecordType {
	fileLines := make([]RecordType, 0)
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	defer readFile.Close()
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		s := strings.Split(line, " ")
		if len(s) == 2 {
			var rec RecordType
			rec.md5 = s[0]
			rec.path = s[1]
			fileLines = append(fileLines, rec)
		}
	}
	return fileLines
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
		if _, err := os.Stat(target); err == nil {
			/*
				equal, err := equalfile.CompareFile(sourcePathname, target)
				if err != nil {
					fmt.Printf("equal: error: %v\n", err)
					os.Exit(3)
				}

				if equal {
					fmt.Println("equal: files match")
					os.Exit(0)
				}

				fmt.Println("equal: files differ")
			*/
			//fmt.Println("DEL ", target)
			return true
		} else {
			fmt.Println("echo NO-MATCH ", sourcePathname)
			return false
		}
	} else {
		fmt.Println("echo NO-PATTERN ", sourcePathname)
	}
	return false
}
