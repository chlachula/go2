package windrives 

import (
	"os"
)

func WinAvailableLetterDrives() []string {
	letters := make([]string, 0)
	for ch := 'A'; ch <= 'Z'; ch++ {
		if isDriveLetterAvailable(byte(ch)) {
			letters = append(letters, string(ch)+":")
		}
	}
	return letters
}
func isDriveLetterAvailable(ch byte) bool {
	root := string(ch) + ":\\"
	if _, err := os.Stat(root); err == nil {
		return true
	}
	return false
}
