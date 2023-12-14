package main

import (
	"fmt"
	"os"
	"time"

	a "github.com/chlachula/go2/wikilang"
)

func help(msg string) {
	if msg != "" {
		fmt.Println(msg)
	}
	helptext := `Displays link to the alternative language version
	Usage:
	go2 url lang
	Example:
	go2 https://en.wikipedia.org/wiki/Cassiopeia_(constellation) cs
	`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())

	urlString := "https://en.wikipedia.org/wiki/Cassiopeia_(constellation)"
	lang := "cs"
	//help("")
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
	fmt.Println("For language", lang, "found for ", urlString, "\nlink: ", langLink, "\nname: ", name)

}
