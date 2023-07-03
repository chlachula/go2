package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"

	a "github.com/chlachula/go2/file_YMD_match"
)

func f1(sourceDir, targetDir string) {
	err := filepath.Walk(sourceDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		//fmt.Printf("visited file or dir: %q\n", path)
		if f, err := os.Stat(path); err == nil {
			if !f.IsDir() {
				a.ExistsYMD(strings.ReplaceAll(path, "\\", "/"), targetDir)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", sourceDir, err)
		return
	}
}
func f2(toProcessFile, processed string) {
	tArr := a.SumsToArr(toProcessFile)
	pArr := a.SumsToArr(processed)
	fmt.Printf("Zpracovat: %d, zpracovano: %d \n", len(tArr), len(pArr))
	a.FindMatches(tArr, pArr)
}

func help() {
	helptext := `
Usage:
go2 sourceDir targetDir 
Example:
go2  /e/Arc-Pics/2003 /e/Arc-Pics/Zpracovat
	`
	fmt.Println(helptext)
}
func main() {
	defer func(start time.Time) {
		fmt.Printf("Elapsed time %s\n", time.Since(start))
	}(time.Now())
	if len(os.Args) < 3 {
		help()
		os.Exit(1)
	}
	//fmt.Printf("%t\n", a.ExistsYMD(os.Args[1], os.Args[2]))
	//f1(os.Args[1], os.Args[2])
	f2(os.Args[1], os.Args[2])
}
