package dirreport

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"
)

type HtmlDataType = struct {
	DirName string
}
type DirInf struct {
	Path      string
	Files     []os.FileInfo
	Dirs      []DirInf
	Name      string
	FilesNums int
	FilesSize int64
	TotalSize int64
}
type FileInfo2Sort struct {
	Name    string
	Size    int64
	ModTime time.Time
	Mode    os.FileMode
	IsDir   bool
}

var Dir string = "."
var DI DirInf
var ExcludeDotDirs = true
var verbose = false

const htmlHead = `<html><head><title>%s</title>
<style>
body {
	font-family: Arial, sans-serif;
	margin: 0;
	padding: 0;
}
.textCenter{
	text-align: center;
}
nav {
	background-color: #333;
	color: #fff;
	text-align: center;
	padding: 10px 0;
}

nav ul {
	list-style: none;
	margin: 0;
	padding: 0;
}

nav li {
	display: inline;
	margin-right: 20px;
}

nav a {
	text-decoration: none;
	color: #fff;
	font-weight: bold;
	font-size: 16px;
}

nav a:hover {
	color: #ffd700; /* Change the color on hover */
}
</style>
</head>
<body>`

const htmlPage2 = `
<h1>Index of %s</h1>
<!--SVG dirs image-->
Sub dir: %s
<pre>
      <a href="?C=N;O=D">Name</a>                              <a href="?C=L;O=A">Last modified</a>       <a href="?C=S;O=A">Size</a> <a href="?C=M;O=A">Mode</a>      <a href="?C=D;O=A">Description</a>
<hr/>%s%s<hr/></pre>
</body></html>
 `
const parentDirectory = "      <a href=\"%s\">Parent Directory</a>\n"

const htmlTemplateDir = `<h1>Hello dir {{.DirName}}</h1>
`
const htmlEnd = `</body></html>`

func verbosePrint(s string) {
	if verbose {
		println(s)
	}
}

func numberPower3Units() []string {
	// unmutable slice returned
	return []string{" ", "K", "M", "G", "T", "P", "E", "Z"}
}

// three chars mantissa and unit
func num10p3str(n int64) string {
	u := numberPower3Units()
	i := 0
	n1 := n
	var mod int64
	for n1 >= 1000 {
		mod = n1 % 1000
		n1 = n1 / 1000
		i += 1
	}
	s := fmt.Sprintf("%3d", n1)
	if n1 < 10 && i > 0 {
		if n1 == 9 && mod > 949 {
			s = " 10"
		} else {
			s = fmt.Sprintf("%3.1f", float64(n1)+float64(mod)*0.001)
		}
	}
	s += u[i]
	return s
}

func getHtmlData() HtmlDataType {
	data := HtmlDataType{
		DirName: Dir,
	}
	return data
}

type MenuItem struct {
	Link string
	Name string
}

var menuItems = []MenuItem{{Link: "/", Name: "Home"},
	{Link: "/show-dir", Name: "One page"},
	{Link: "/show-dir2", Name: "Summarized subdirectories"},
	{Link: "/#contact", Name: "Contact"},
	{Link: "/#about", Name: "About"}}

func navMenu(ActiveLink string) string {
	s := "<nav>\n    <ul>\n"
	for _, item := range menuItems {
		link := item.Link
		if item.Link == ActiveLink {
			link = "#"
		}
		s += fmt.Sprintf("<li><a href=\"%s\">%s</a></li>", link, item.Name)
	}
	s += "    </ul>\n</nav>\n\n"
	return s
}

func findDI(di *DirInf, relPathName string) *DirInf {
	name := relPathName
	i := strings.Index(relPathName, "/")
	if i > -1 {
		name = relPathName[:i]
		relPathName = relPathName[i+1:]
	}
	verbosePrint("DEBUG-findDI: Path=" + di.Path)
	for _, d := range di.Dirs {
		verbosePrint("DEBUG-findDI: " + name + ": " + d.Name)
		if name == d.Name {
			verbosePrint("DEBUG-findDI FOUND!!!: " + name + ", Path=" + d.Path)
			if i < 0 {
				return &d
			} else {
				return findDI(&d, relPathName)
			}

		}
	}
	return nil
}
func spaces(name string, length int) string {
	s := ""
	for ; len(s)+len(name) < length; s += " " {

	}
	return s
}
func sizeSpan(size int64) string {
	return fmt.Sprintf("<span title=\"%d\">%5s</span>", size, num10p3str(size))
}
func createFileInfo2(dirInf DirInf, f os.FileInfo) FileInfo2Sort {
	var f2 FileInfo2Sort
	f2.Name = f.Name()
	f2.ModTime = f.ModTime()
	f2.Mode = f.Mode()
	f2.Size = f.Size()
	f2.IsDir = false
	if f.IsDir() {
		f2.IsDir = true
		if di := findDI(&dirInf, f.Name()); di != nil {
			f2.Size = di.TotalSize
		}
	}
	return f2
}
func sortFiles2(dirInf DirInf) []FileInfo2Sort {
	files := make([]FileInfo2Sort, 0)
	for _, f := range dirInf.Files {
		if f.IsDir() {
			if !(ExcludeDotDirs && strings.HasPrefix(f.Name(), ".")) {
				f2 := createFileInfo2(dirInf, f)
				files = append(files, f2)
			}
		} else {
			f2 := createFileInfo2(dirInf, f)
			files = append(files, f2)
		}
	}
	sort.SliceStable(files, func(i, j int) bool { return files[i].Size > files[j].Size })
	return files
}
func dirInfPath2string(dirInf *DirInf, rootpath string, path string) string {
	DItoShow := dirInf
	if path != "" {
		if subDinf := findDI(dirInf, path); subDinf != nil {
			DItoShow = subDinf
		}
	}
	s := ""
	f0 := "      <a href=\"%s\">%s</a>%s %-18s %s %s \n"
	f1 := "      %-33s %-18s %s %s \n"
	if rootpath != "" {
		rootpath += path + "/"
	} else {
		if path != "" {
			rootpath = path + "/"
		}
	}
	files := sortFiles2(*DItoShow)

	for _, f := range files {
		modTime := f.ModTime.Format("2006-Jan-01 15:04")
		if f.IsDir {
			if !(ExcludeDotDirs && strings.HasPrefix(f.Name, ".")) {
				link := "?d=" + rootpath + f.Name
				if di := findDI(DItoShow, f.Name); di != nil {
					s += fmt.Sprintf(f0, link, f.Name, spaces(f.Name, 33), modTime, sizeSpan(di.TotalSize), f.Mode)
				} else {
					fmt.Printf("error findDI rootpath:%s, path:%s, name:%s\n", rootpath, path, f.Name)
				}
			}
		} else {
			s += fmt.Sprintf(f1, f.Name, modTime, sizeSpan(f.Size), f.Mode)
		}
	}
	s += fmt.Sprintf("<hr/>      %-52s %s\n", "Total size", sizeSpan(dirInf.TotalSize))
	return s
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlHead, "Home")
	fmt.Fprint(w, navMenu("/"))

	fmt.Fprintf(w, "<h1>Home</h1><h1>Show dir %s</h1><a href=\"%s\">%s</a> -  <a href=\"%s\">%s</a> <br/>", Dir,
		"/show-dir", "One page", "/show-dir2", "Summarized subdirectories")

	fmt.Fprint(w, htmlEnd)
}
func HandleShowDir(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlHead, "One page dir report")
	fmt.Fprint(w, navMenu("/show-dir"))

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

func HandleShowDir2(w http.ResponseWriter, r *http.Request) {
	verbosePrint("\n\nr.URL = " + r.URL.String())
	rootPath := ""
	leaf := ""
	pageBody := ""
	if d := r.URL.Query().Get("d"); d != "" {
		d = strings.TrimSuffix(d, "/")
		subDI := findDI(&DI, d)
		i := strings.LastIndex(d, "/")
		if i > -1 {
			rootPath = d[:i+1]
			leaf = d[i+1:]
			verbosePrint("--Handle 3 second and more subdirs")
			pageBody = dirInfPath2string(subDI, rootPath, leaf)
		} else {
			verbosePrint("--Handle 2 first subdir")
			leaf = d
			pageBody = dirInfPath2string(subDI, rootPath, leaf) // rootPath==""
		}
	} else {
		verbosePrint("--Handle 1 root")
		pageBody = dirInfPath2string(&DI, rootPath, d) // rootPath==""
	}
	parentDirLink := "/show-dir2?d=" + rootPath
	parentDirLink = fmt.Sprintf(parentDirectory, parentDirLink)
	currentDir := rootPath
	if rootPath == "" {
		currentDir = leaf
	} else {
		currentDir += leaf
	}
	if currentDir == "" {
		currentDir = "."
		parentDirLink = "      .\n"
	}
	fmt.Fprintf(w, htmlHead, "Summarized subdirectories")
	fmt.Fprint(w, navMenu("/show-dir2"))

	fmt.Fprintf(w, htmlPage2, Dir, currentDir, parentDirLink, pageBody)
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
	verbosePrint("SummarizeDirectory p=" + dirPath)
	var dirInf DirInf
	dirInf.Path = dirPath
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
			dirInf.FilesSize += fileInfo.Size()
		}
	}
	dirInf.TotalSize += dirInf.FilesSize
	return dirInf
}

func SetDirInf() {
	DI = SummarizeDirectory(Dir)
}
