package windrives

import "os"

// available letter drives in windows OS like C: etc.
func WinAvailableLetterDrives() (letters []string) {
	for ch := 'A'; ch <= 'Z'; ch++ {
		d := string(ch) + ":"
		if _, err := os.Stat(d + "\\"); err == nil {
			letters = append(letters, d)
		}
	}
	return letters
}
