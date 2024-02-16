package dirreport

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type HtmlDataType = struct {
	DirName string
}
type DirInf struct {
	Files     []os.FileInfo
	Dirs      []DirInf
	Name      string
	FilesNums int
	FilesSize int
	TotalSize int
}

var Dir string = "."
var DI DirInf
var ExcludeDotDirs = true

const htmlPage2 = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">
<html>
 <head>
  <title>Index of %s</title>
 </head>
 <body>
<h1>Index of %s</h1>
<!--SVG dirs image--> %s
<pre>
      <a href="?C=N;O=D">Name</a>                    <a href="?C=L;O=A">Last modified</a>              <a href="?C=S;O=A">Size</a> <a href="?C=M;O=A">Mode</a>      <a href="?C=D;O=A">Description</a>
<hr />      <a href="%s">Parent Directory</a>                             -   
%s

      <a href="APODstyle.css">APODstyle.css</a>           11-Dec-2019 15:23  143   
      <a href="IoTest.html">IoTest.html</a>             09-Feb-2021 21:50  2.7K  
<hr></pre>
</body></html>
 `
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

	fmt.Fprintf(w, "<h1>Home</h1><h1>Show dir <a href=\"%s\">%s</a></h1><h1>Show dir <a href=\"%s\">%s</a></h1>",
		"/show-dir", Dir, "/show-dir2", Dir)

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
func findDI(di *DirInf, name string) *DirInf {
	var emptyDI DirInf
	for _, d := range di.Dirs {
		if name == d.Name {
			return &d
		}
	}
	fmt.Println("ERROR, not found DI ", name)
	return &emptyDI
}
func spaces23(name string) string {
	s := ""
	for ; len(s)+len(name) < 23; s += " " {

	}
	return s
}
func dirInfPath2string(dirInf *DirInf, path string) string {
	DItoShow := dirInf
	if path != "" {
		if subDinf := findDI(dirInf, path); subDinf != nil {
			DItoShow = subDinf
		}
	}
	s := ""
	f0 := "      <a href=\"%s\">%s</a>%s %-18s %12d %s \n"
	f1 := "      %-23s %-18s %12d %s \n"
	for _, f := range DItoShow.Files {
		modTime := f.ModTime().Format("2006-Jan-01 15:04")
		if f.IsDir() {
			link := "?d=" + f.Name()
			di := findDI(dirInf, f.Name())
			s += fmt.Sprintf(f0, link, f.Name(), spaces23(f.Name()), modTime, di.TotalSize, f.Mode())
		} else {
			s += fmt.Sprintf(f1, f.Name(), modTime, f.Size(), f.Mode())
		}
	}
	return s
}
func dirInf2string(dirInf DirInf) string {
	s := ""
	f0 := "      <a href=\"%s\">%s</a>%s %-18s %12d %s \n"
	f1 := "      %-23s %-18s %12d %s \n"
	for _, f := range dirInf.Files {
		modTime := f.ModTime().Format("2006-Jan-01 15:04")
		if f.IsDir() {
			link := "?d=" + f.Name()
			di := findDI(&dirInf, f.Name())
			s += fmt.Sprintf(f0, link, f.Name(), spaces23(f.Name()), modTime, di.TotalSize, f.Mode())
		} else {
			s += fmt.Sprintf(f1, f.Name(), modTime, f.Size(), f.Mode())
		}
	}
	return s
}
func HandleShowDir2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.URL =", r.URL)
	pageBody := ""
	if d := r.URL.Query().Get("d"); d != "" {
		pageBody = dirInfPath2string(&DI, d)
	} else {
		pageBody = dirInf2string(DI)
	}
	fmt.Fprintf(w, htmlPage2, Dir, Dir, "", "..", pageBody)
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
	s += fmt.Sprintf("\nTotal of %d files of size %d bytes in the directory %s\n", numberOfFiles, totalFilesSize, dirPath)
	return s, nil
}
func SummarizeDirectory(dirPath string) DirInf {
	var dirInf DirInf
	dir, err := os.Open(dirPath) // Open the directory
	if err != nil {
		fmt.Printf("error opening %s: %s\n", dirPath, err.Error())
		return dirInf
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1) // Read the contents of the directory
	if err != nil {
		fmt.Printf("error Readdir %s: %s\n", dirPath, err.Error())
		return dirInf
	}
	dirInf.Files = fileInfos
	dirInf.Dirs = make([]DirInf, 0)
	// Loop through each file/directory
	for _, fileInfo := range fileInfos {
		if ExcludeDotDirs && fileInfo.IsDir() && strings.HasPrefix(fileInfo.Name(), ".") {
			continue
		}
		fullPath := filepath.Join(dirPath, fileInfo.Name())
		// Check if the file/directory is a directory
		if fileInfo.IsDir() {
			di := SummarizeDirectory(fullPath) // If it is a directory, recursively call this function
			di.Name = fileInfo.Name()
			dirInf.TotalSize += di.TotalSize
			dirInf.Dirs = append(dirInf.Dirs, di)
		} else {
			dirInf.FilesNums++
			dirInf.FilesSize += int(fileInfo.Size())
		}
	}
	dirInf.TotalSize += dirInf.FilesSize
	return dirInf
}

func SetDirInf() {
	DI = SummarizeDirectory(Dir)
}
