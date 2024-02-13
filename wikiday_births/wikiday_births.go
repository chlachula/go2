package wikiday_births

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type DayRecord struct {
	MM      int
	DD      int
	Records []PersonRecord
}
type PersonRecord struct {
	Age  int
	YoB  int
	YoD  int
	Wiki string
}

var verbose bool = true
var re1 = regexp.MustCompile(`(?m).*<li>.*title="(\d{4})".*?<a href="/wiki/(.*?)".*\(d\. (.*)?\)`)

func getMMDD(mmdd string) (int, int) {
	if t, err := time.Parse("1/2", mmdd); err != nil {
		return 0, 0
	} else {
		return int(t.Month()), t.Day()
	}
}
func makeUrlString(mm, dd int) string {
	prefix := "https://en.wikipedia.org/wiki/" // February_13
	t := time.Date(2024, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s%s_%d", prefix, t.Format("January"), dd)
}
func WikiDay(mmdd string) {
	m, d := getMMDD(mmdd)
	urlString := makeUrlString(m, d)
	page, err := DownloadPage(urlString)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	dayRec, err1 := Find(m, d, page)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(1)
	}

	sort.SliceStable(dayRec.Records, func(i, j int) bool {
		//i,j are represented for two value of the slice .
		return dayRec.Records[i].Age < dayRec.Records[j].Age
	})
	bytes, _ := json.Marshal(dayRec)
	s := strings.ReplaceAll(string(bytes), "},", "},\n")
	s = strings.ReplaceAll(s, "[", "\n[")
	fmt.Printf("\ndayRec: %s\n\n", s)
}

func DownloadPage(urlString string) (string, error) {
	response, err := http.Get(urlString)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bytes, err1 := io.ReadAll(response.Body)
	if err1 != nil {
		return "", err1
	}
	if verbose {
		fmt.Printf("From %s downloaded %d bytes\n", urlString, len(bytes))
	}

	return string(bytes), nil
}

func extractInfo(line string) *PersonRecord {
	var r PersonRecord
	match := re1.FindStringSubmatch(line)
	if len(match) > 3 {
		y1, err1 := strconv.Atoi(match[1])
		y2, err2 := strconv.Atoi(match[3])
		if err1 == nil && err2 == nil {
			age := y2 - y1
			r.Age = age
			r.YoB = y1
			r.YoD = y2
			r.Wiki = match[2]
			return &r
		}
	}
	return nil
}

func Find(m int, d int, page string) (DayRecord, error) {
	var dayBirths DayRecord
	dayBirths.DD = d
	dayBirths.MM = m
	dayBirths.Records = make([]PersonRecord, 0)
	lines := strings.Split(page, "\n")
	if verbose {
		fmt.Printf("File has %d lines\n", len(lines))
	}
	birthsBlock := false
	deadsBlock := false
	for _, line := range lines {
		if strings.Contains(line, "id=\"Births\"") {
			birthsBlock = true
		}
		if strings.Contains(line, "id=\"Deaths\"") {
			birthsBlock = false
			deadsBlock = true
		}
		if strings.Contains(line, "id=\"Holidays_and_observances\"") {
			birthsBlock = false
			deadsBlock = false
		}
		if birthsBlock {
			r := extractInfo(line)
			if r != nil {
				dayBirths.Records = append(dayBirths.Records, *r)
			}
		}
	}
	_ = deadsBlock
	return dayBirths, nil
}
