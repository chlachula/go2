package dirreport

import (
	"fmt"
	"net/http"
	"text/template"
)

var Dir string

type HtmlDataType = struct {
	DirName string
}

const htmlHead = `<html><head><title>%s</title></head>
<body>`
const htmlTemplate = `<h1>Hello dir {{.DirName}}</h1>
`
const htmlEnd = `</body></html>`

func getHtmlData() HtmlDataType {
	data := HtmlDataType{
		DirName: Dir,
	}
	return data
}

func ShowDir(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlHead, "Dir Report")

	if t, err := template.New("webpage1").Parse(htmlTemplate); err == nil {
		data := getHtmlData()
		if err = t.Execute(w, data); err != nil {
			fmt.Fprintf(w, "<h1>error %s</h1>", err.Error())
		}
	}

	fmt.Fprint(w, htmlEnd)
}
