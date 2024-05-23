package svgImagesJs

import (
	"fmt"
	"net/http"
	"os"
)

var Atext string = "text sample"

func HandlerHome(w http.ResponseWriter, r *http.Request) {
	//writeHtmlHeadAndMenu(w, "/", "Home")
	fmt.Fprint(w, `<html>
 <head>
 <meta http-equiv="refresh" content="0; url=/htmlSvgImagesJs">
 	  <title>redirect to EyepieceLabels</title>
 </head>
 <body>
  <h1>Click to: <a href="/htmlSvgImagesJs">htmlSvgImagesJs</a></h1>
 </body>
</html>
	`)
}
func HandlerHtmlJsSvgPages(w http.ResponseWriter, r *http.Request) {
	//writeHtmlHeadAndMenu(w, "/svg-roundlogo-color", "Color")

	fmt.Fprint(w, "<html><head><title>HTML javascript SVG pages</title></head>\n")
	fmt.Fprint(w, "<body style=\"text-align: center;\">\n")
	fmt.Fprint(w, "<h1>HTML javascript SVG pages</h1>\n")
	fmt.Fprint(w, "<h1><a href=\"/eyepieceLabels\">Eyepiece Labels</a></h1>\n")
	fmt.Fprint(w, "<h1><a href=\"/eyepieceLabels2\">Eyepiece Labels 2</a></h1>\n")

	fmt.Fprint(w, "</body></html>\n")

}
func HandlerEyepieceLabels(w http.ResponseWriter, r *http.Request) {
	//writeHtmlHeadAndMenu(w, "/svg-roundlogo-color", "Color")

	fmt.Fprint(w, "<html><head><title>Eyepiece Labels</title></head>\n")
	fmt.Fprint(w, "<body style=\"text-align: center;\">\n")
	fmt.Fprint(w, "<h1>Eyepiece Labels</h1>\n")

	fmt.Fprint(w, "</body></html>\n")

}
func HandlerEyepieceLabels2(w http.ResponseWriter, r *http.Request) {
	str := LoadTextFile("eyepieceLabels.svg")
	fmt.Fprint(w, str)

}
func LoadTextFile(filename string) string {
	bytes, err := os.ReadFile(filename) //Read entire file content. No need to close
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return string(bytes)
}
