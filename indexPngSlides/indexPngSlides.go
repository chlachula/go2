package indexPngSlides

import (
	"fmt"
	"os"
	"strings"
)

var Hello string = "Hello World"

func FindPngFiles(dir string) {
	var PNGfiles []string = make([]string, 0)
	d, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := d.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, f := range files {
		fmt.Printf("name=%s, isDir=%t   %+v\n", f.Name(), f.IsDir(), f)
		if strings.HasSuffix(f.Name(), ".png") {
			PNGfiles = append(PNGfiles, f.Name())
		}
	}
	fmt.Printf("PNG files  %+v\n", PNGfiles)
}
