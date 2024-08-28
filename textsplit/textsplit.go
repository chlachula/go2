package textsplit

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

/*

Content of the newsletters
Number Date
121	11/1973
1   Actions in November
1   Joe Doe: Bussiness stuff
2   Talk about Solar System
3   Public session
---
122	1/1974
....
---END //end of processing

*/

var (
	OutFilePrefix = "Prefix"
	InpFile       = "InpFile.txt"
	InpDir        = "."
	OutDir        = "."
)

func createFile(headerLine, fileLines string) {
	NNN := headerLine[:3]
	Date := headerLine[4:]
	Date = strings.Trim(Date, " ")
	fname := OutFilePrefix + "-" + NNN + ".txt"

	f, err := os.Create(fname)
	if err != nil {
		fmt.Printf("Error creating file %s: %s\n", fname, err.Error())
		return
	}
	defer f.Close()

	bytes := []byte(NNN + " " + Date + "\n\n" + fileLines)
	if _, err := f.Write(bytes); err != nil {
		fmt.Printf("Error writting into file %s: %s\n", fname, err.Error())
	}
}

func ProcessFile() {
	fullPath := filepath.Join(InpDir, InpFile)
	bytes, err := os.ReadFile(fullPath) // Read entire file content
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	text := string(bytes)
	lines := strings.Split(text, "\n")
	linesCount := len(text)
	headerLine := ""
	fileLines := ""
	startNewFile := true
	countFiles := 0
	for i := 2; i < linesCount; i++ {
		line := lines[i]
		if startNewFile {
			headerLine = line
			startNewFile = false
			continue
		}
		if strings.HasPrefix(line, "---") {
			createFile(headerLine, fileLines)
			countFiles += 1
			startNewFile = true
			fileLines = ""
			if strings.HasPrefix(line, "---END") {
				break
			}
		}
		fileLines += line
		fileLines += "\n"
	}

	fmt.Printf("\n%d files has been created\n", countFiles)
}
