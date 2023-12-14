package wikilang

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

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

	return string(bytes), nil
}

func FindLink(page string, lang string) (string, error) {
	langWiki := lang + ".wikipedia.org"
	needle := "https://" + langWiki
	i1 := strings.Index(page, needle)
	if i1 < 0 {
		return "", fmt.Errorf(langWiki + "not found")
	}
	i2 := i1 + len(needle)
	for page[i2] != '"' && i2 < len(page) {
		i2 += 1
	}
	if i2 >= len(page) {
		return "", fmt.Errorf("no '\"' after " + langWiki)
	}
	return page[i1:i2], nil
}
