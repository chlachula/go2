package dotdir

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirAbsInfo(dir string) string {
	abs, err := filepath.Abs(dir)
	f1 := "filepath.Abs(\"%s\"): %s"
	s := ""
	if err != nil {
		return fmt.Sprintf(f1, dir, err.Error())
	} else {
		s += fmt.Sprintf(f1, dir, abs)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		s += fmt.Sprintf("\nos.ReadDir(\"%s\"):%s", dir, err.Error())
		return s
	}
	var countFiles, countDirs int
	var files, dirs string
	for _, entry := range entries {
		if entry.IsDir() {
			countDirs++
			dirs += entry.Name() + " "
		} else {
			countFiles++
			files += entry.Name() + " "
		}
	}
	if countDirs == 0 && countFiles == 0 {
		return s + "\nThere are no entries."
	}
	s += fmt.Sprintf(": %d directories, %d files.", countDirs, countFiles)
	s += "\nFiles: " + files
	s += "\nDirectories: " + dirs
	return s
}
