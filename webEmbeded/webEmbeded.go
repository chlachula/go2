package webEmbeded

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed to_embed
var content embed.FS
var dirName string = "to_embed"

func FsHandler() http.Handler {
	sub, err := fs.Sub(content, dirName)
	if err != nil {
		panic(err)
	}

	return http.FileServer(http.FS(sub))
}

func main() {
	colonPort := ":8080"
	print("Serving from embeded files. Listenning at " + colonPort + ". CTRL+C to stop.")
	http.ListenAndServe(colonPort, FsHandler())
}
