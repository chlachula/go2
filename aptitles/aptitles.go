package aptitles

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

type ApodArchiveTitle struct {
	YMD   string `json:"ymd"`
	Title string `json:"title"`
}

type ApodArchiveTitles struct {
	TitlesName string
	Titles     []ApodArchiveTitle
}

var mapYMDtitles map[string]string
var re = regexp.MustCompile(`(?m).*"ap(?P<yymmdd>\d{6})\.html">(?P<title>.*)<\/a>.*`)

var tmpfile = `2021 August 02:  <a href="ap210802.html">The Hubble Ultra Deep Field in Light and Sound</a><br>
2021 August 01:  <a href="ap210801.html">Pluto in Enhanced Color</a><br>
2021 July 31:  <a href="ap210731.html">Remembering NEOWISE</a><br>
2021 July 30:  <a href="ap210730.html">Mimas in Saturnlight</a><br>
2021 July 29:  <a href="ap210729.html">The Tulip and Cygnus X 1("black hole")</a><br>`

func writeStr(f *os.File, str string) {
	if _, err := f.Write([]byte(str)); err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
}
func Create() {
	fname := "aptitles.json"
	if err := createAPtitlesJson(fname, tmpfile); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
func createAPtitlesJson(fname string, fstr string) error {
	f, err := os.Create(fname)
	if err != nil {
		return err
	}
	defer f.Close()
	writeStr(f, "{\"titles\": [")
	lines := strings.Split(fstr, "\n")
	comma := "\n"
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 2 {
			escapedTitle := strings.ReplaceAll(match[2], "\"", "\\\"")
			jsonLine := fmt.Sprintf(`%s {"ymd":"%s", "title":"%s"}`, comma, match[1], escapedTitle)
			writeStr(f, jsonLine)
			comma = ",\n"
		}
	}
	writeStr(f, "\n ]\n}\n")
	return nil
}
func readEntireFileToBytes(fname string) []byte {
	reader, err := os.Open(fname)
	if err != nil {
		fmt.Println(err)
		return make([]byte, 0)
	}
	if bytes, err := io.ReadAll(reader); err != nil {
		fmt.Println(err)
		return make([]byte, 0)
	} else {
		return bytes
	}
}
func loadAPtitlesJson(fname string) (map[string]string, error) {
	m := make(map[string]string, 0)
	bytes := readEntireFileToBytes(fname)
	var archiveTitles ApodArchiveTitles

	if err := json.Unmarshal(bytes, &archiveTitles); err != nil {
		return m, err
	}

	for _, s := range archiveTitles.Titles {
		m[s.YMD] = s.Title
	}
	return m, nil
}
func LoadAPODarchive() error {
	fname := "aptitles.json"
	var err error
	if mapYMDtitles, err = loadAPtitlesJson(fname); err != nil {
		return err
	}
	fmt.Printf("Total %d APOD titles loaded\n", len(mapYMDtitles))
	return nil
}

func SearchTitle(yymmdd string) (string, error) {
	if len(yymmdd) != 6 {
		return "", fmt.Errorf("yymmdd: %s does not have length 6", yymmdd)
	}
	return "OK", nil
}
