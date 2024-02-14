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
	MM       int
	DD       int
	Filename string
	Deaths   []PersonRecord
	Births   []PersonRecord
}
type PersonRecord struct {
	Age  int
	YoB  int
	YoD  int
	Wiki string
}

var verbose bool = true

// https://regex101.com/r/4WLTRe/1
var reBirth = regexp.MustCompile(`(?m).*<li>.*title="(\d{4})".*?<a href="/wiki/(.*?)".*\(d\. (.*)?\)`)
var reDeath = regexp.MustCompile(`(?m).*<li>.*title="(\d{4})".*?<a href="/wiki/(.*?)".*\(b\. (.*)?\)`)

func getMMDD(mmdd string) (int, int) {
	if t, err := time.Parse("1/2", mmdd); err != nil {
		return 0, 0
	} else {
		return int(t.Month()), t.Day()
	}
}
func makeWikiDayFilename(mm, dd int) string {
	t := time.Date(2024, time.Month(mm), dd, 0, 0, 0, 0, time.UTC)
	return fmt.Sprintf("%s_%d", t.Format("January"), dd)
}

func WikiDay(mmdd string) {
	jsonStr := ""
	mm, dd := getMMDD(mmdd)
	filename := makeWikiDayFilename(mm, dd)
	filenameJson := filename + ".json"
	if bytes, err := os.ReadFile(filenameJson); err == nil { //Read entire file content. No need to close
		jsonStr = string(bytes)
	} else {
		jsonStr = downloadWikiDay(mm, dd, filename)
	}
	fmt.Println(jsonStr)
}
func downloadWikiDay(m, d int, filename string) string {
	prefix := "https://en.wikipedia.org/wiki/" // February_13
	urlString := prefix + filename
	page, err := DownloadPage(urlString)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	dayRec, err1 := Find(m, d, filename, page)
	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(1)
	}

	sort.SliceStable(dayRec.Births, func(i, j int) bool {
		//i,j are represented for two value of the slice .
		return dayRec.Births[i].Age < dayRec.Births[j].Age
	})
	sort.SliceStable(dayRec.Deaths, func(i, j int) bool {
		//i,j are represented for two value of the slice .
		return dayRec.Deaths[i].Age < dayRec.Deaths[j].Age
	})
	bytes, _ := json.Marshal(dayRec)
	s := strings.ReplaceAll(string(bytes), "},", "},\n")
	s = strings.ReplaceAll(s, "[", "\n[")
	x := "\"Births\":"
	s = strings.ReplaceAll(s, x, "\n"+x)
	x = "\"Deaths\":"
	s = strings.ReplaceAll(s, x, "\n"+x)

	bytes = []byte(s)
	if f, err := os.Create(filename + ".json"); err == nil {
		if _, err := f.Write(bytes); err != nil {
			fmt.Printf("error: %s\n", err.Error())
		}
	} else {
		fmt.Printf("error: %s\n", err.Error())
	}

	return s
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

func extractBirthInfo(line string) *PersonRecord {
	var r PersonRecord
	match := reBirth.FindStringSubmatch(line)
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
func extractDeathInfo(line string) *PersonRecord {
	var r PersonRecord
	match := reDeath.FindStringSubmatch(line)
	if len(match) > 3 {
		y2, err1 := strconv.Atoi(match[1])
		y1, err2 := strconv.Atoi(match[3])
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

func Find(m int, d int, filename string, page string) (DayRecord, error) {
	var wikiDay DayRecord
	wikiDay.DD = d
	wikiDay.MM = m
	wikiDay.Filename = filename
	wikiDay.Births = make([]PersonRecord, 0)
	wikiDay.Deaths = make([]PersonRecord, 0)
	lines := strings.Split(page, "\n")
	if verbose {
		fmt.Printf("File has %d lines\n", len(lines))
	}
	birthsBlock := false
	deathsBlock := false
	for _, line := range lines {
		if strings.Contains(line, "id=\"Births\"") {
			birthsBlock = true
		}
		if strings.Contains(line, "id=\"Deaths\"") {
			birthsBlock = false
			deathsBlock = true
		}
		if strings.Contains(line, "id=\"Holidays_and_observances\"") {
			birthsBlock = false
			deathsBlock = false
		}
		if birthsBlock {
			r := extractBirthInfo(line)
			if r != nil {
				wikiDay.Births = append(wikiDay.Births, *r)
			}
		}
		if deathsBlock {
			r := extractDeathInfo(line)
			if r != nil {
				wikiDay.Deaths = append(wikiDay.Deaths, *r)
			}
		}
	}
	return wikiDay, nil
}
