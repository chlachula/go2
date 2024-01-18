package dotdir

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func DotWebHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><head><title>dot dir</title></head><body><h1>dot dir</h1>\n<pre>")
	absDot := DirAbsInfo(".")
	fmt.Fprintf(w, "%s\n</pre>\n", absDot)
}

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
	return s + "\n"
}

func Web(colonPort string) {

	fs := http.FileServer(http.Dir("files"))
	prefixHandler := http.StripPrefix("/files1/", fs)
	http.Handle("/files1/", prefixHandler)
	fmt.Print("OK: handle to /files1/ \n")

	relFiles := "./files"
	sF2s := "/files2/"
	if stat, err := os.Stat(relFiles); err == nil && stat.IsDir() {
		http.Handle(sF2s, http.StripPrefix(sF2s, http.FileServer(http.Dir(relFiles))))
		fmt.Print("OK: handle to /files2/ \n")
	} else {
		fmt.Printf("error:  %s\n\n", err.Error())
	}

	if absFiles, err := filepath.Abs(relFiles); err == nil {
		sF3s := "/files3/"
		fmt.Printf("filepath.Abs(\"%s\"): %s\n", relFiles, absFiles)
		http.Handle(sF3s, http.StripPrefix(sF3s, http.FileServer(http.Dir(absFiles))))
		fmt.Print("OK: handle to /files3/ \n")
	} else {
		fmt.Println(err.Error())
	}

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/dot", DotHandler)

	print("...listening at " + colonPort)
	if err := http.ListenAndServe(colonPort, nil); err != nil {
		fmt.Println("error ", err.Error())
	}
}

func home(homelink bool) string {
	link := "home"
	part := `<html><head><title>dot dir</title></head><body style="text-align:center">
	%s
	<hr size="1" noshadow="noshadow" />
	`
	if homelink {
		link = `<a href="/">home</a>`
	}
	return fmt.Sprintf(part, link)
}
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, home(false))
	part := `<h1>Hello <a href="/dot">dot</a>!</h1>
	Handles to the subdirectory files<br/>
	<a href="/files1/">/files1/</a><br/>
	<a href="/files2/">/files2/</a><br/>
	<a href="/files3/">/files3/</a><br/>
	</body></html>`
	fmt.Fprint(w, part)
}
func DotHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, home(true))
	part1 := "<h1>Dot directory - current directory information</h1>\n<pre>\n"
	fmt.Fprint(w, part1)
	fmt.Fprint(w, DirAbsInfo("."))
	fmt.Fprint(w, "</pre></body></html>")
}
