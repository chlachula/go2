package indexPngSlides

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var formHtml1 = `<html>
 <head>
 <title>index file</title>
 </head>
 <body>
 %s
 </body>
</html> 
`

func CreateIndexFile(dir string) {
	filename := filepath.Join(dir, "index.html")
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	html := pngFilesToIndex(dir)

	if _, err := f.Write([]byte(html)); err != nil {
		panic(err)
	}
	fmt.Printf(" Created %s\n", filename)

}
func htmlErrorPage(errMsg string) string {
	return fmt.Sprintf(formHtml1, "<h1>"+errMsg+"</h1>")
}
func pngFilesToIndex(dir string) string {
	var PNGfiles []string = make([]string, 0)
	d, err := os.Open(dir)
	if err != nil {
		return htmlErrorPage(err.Error())
	}
	files, err := d.Readdir(0)
	if err != nil {
		return htmlErrorPage(err.Error())
	}

	for _, f := range files {
		//fmt.Printf("name=%s, isDir=%t   %+v\n", f.Name(), f.IsDir(), f)
		if strings.HasSuffix(f.Name(), ".png") {
			PNGfiles = append(PNGfiles, f.Name())
		}
	}
	sort.SliceStable(PNGfiles, func(i, j int) bool { return PNGfiles[i] < PNGfiles[j] })

	s := ""
	s += time.Now().Format("Generated 2006-01-02 Monday 15:04 <br/>\n")
	for _, f := range PNGfiles {
		s += "\n<br/>" + f[:3] + "<img valign=\"top\" src=\"" + f + "\" />"
	}
	fmt.Printf(" Found   %+v\n", PNGfiles)

	return s
}
