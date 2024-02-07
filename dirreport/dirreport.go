package dirreport

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var Dir string = "."
var ExcludeDotDirs = true

type HtmlDataType = struct {
	DirName string
}

const htmlHead = `<html><head><title>%s</title></head>
<body>`
const htmlTemplateDir = `<h1>Hello dir {{.DirName}}</h1>
`
const htmlEnd = `</body></html>`

func getHtmlData() HtmlDataType {
	data := HtmlDataType{
		DirName: Dir,
	}
	return data
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlHead, "Home")

	fmt.Fprintf(w, "<h1>Home</h1><h1>Show dir <a href=\"%s\">%s</a></h1>", "/show-dir", Dir)

	fmt.Fprint(w, htmlEnd)
}
func HandleShowDir(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlHead, "Dir Report")

	if t, err := template.New("webpage2").Parse(htmlTemplateDir); err == nil {
		data := getHtmlData()
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}
	if str, err := displayDirectoryContents(Dir); err == nil {
		fmt.Fprint(w, "\n<pre>\n"+str+"\n</pre>\n")
	} else {
		fmt.Fprintf(w, "<h2>%s</h2>", err.Error())
	}
	fmt.Fprint(w, htmlEnd)
}
func displayDirectoryContents(dirPath string) (string, error) {
	numberOfFiles := 0
	totalFilesSize := 0
	s := ""
	// Open the directory
	dir, err := os.Open(dirPath)
	if err != nil {
		return "", err
	}
	defer dir.Close()

	// Read the contents of the directory
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return "", err
	}

	// Loop through each file/directory
	for _, fileInfo := range fileInfos {
		if ExcludeDotDirs && fileInfo.IsDir() && strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}
		fullPath := filepath.Join(dirPath, fileInfo.Name())

		// Print file/directory information
		s += fmt.Sprintln("Name:", fileInfo.Name())
		s += fmt.Sprintln("Size:", fileInfo.Size(), "bytes")
		s += fmt.Sprintln("Modified Time:", fileInfo.ModTime())
		s += fmt.Sprintln("Mode:", fileInfo.Mode())

		// Check if the file/directory is a directory
		if fileInfo.IsDir() {
			// If it is a directory, recursively call this function
			if s2, err := displayDirectoryContents(fullPath); err != nil {
				return s, err
			} else {
				s += s2
			}
		} else {
			numberOfFiles++
			totalFilesSize += int(fileInfo.Size())
		}
		s += "\n"
	}
	s += fmt.Sprintf("\nTotal of %d files of size %d in the directory %s\n", numberOfFiles, totalFilesSize, dirPath)
	return s, nil
}
