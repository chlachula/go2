package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	a "github.com/chlachula/go2/wikilang"
)

func quotesRemove(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}
func wikiLanguageVersion(urlWiki, langVersion string) {
	urlString := quotesRemove(urlWiki)
	lang := strings.ToLower(langVersion)

	page, err := a.DownloadPage(urlString)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	langLink, name, err1 := a.FindLink(page, lang)
	if err != nil {
		fmt.Println(err1.Error())
		os.Exit(1)
	}
	if langLink == "" {
		langLink = "NOT FOUND"
	}
	if name == "" {
		name = "N/A"
	}
	fmt.Println("For language", lang, "found \nfor : ", urlString, "\nlink: ", langLink, "\nname: ", name)

}
func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `Displays link to the alternative language version
	Usage:
	go2 url lang
	Example:
	go2 "https://en.wikipedia.org/wiki/Cassiopeia_(constellation)" cs
	go2 https://en.wikipedia.org/wiki/Tau_function cs
	`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	if len(os.Args) < 2 {
		help("Not enough arguments")
		os.Exit(1)
	}
	if strings.HasPrefix(os.Args[1], "-h") {
		help("")
		os.Exit(0)
	}
	if len(os.Args) < 3 {
		help("Not enough arguments")
		os.Exit(1)
	}

	wikiLanguageVersion(os.Args[1], os.Args[2])

}
