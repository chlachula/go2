package aptitles

import "fmt"

func LoadAPODarchive() error {
	return nil
}

func SearchTitle(yymmdd string) (string, error) {
	if len(yymmdd) != 6 {
		return "", fmt.Errorf("yymmdd: %s does not have length 6", yymmdd)
	}
	return "OK", nil
}
