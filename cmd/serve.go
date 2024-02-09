package cmd

import (
	"io/fs"
	"log"
	"net"
	"net/http"
)

func ServeWebBuild(webFs fs.FS) net.Listener {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err = http.Serve(ln, http.FileServer(http.FS(webFs))); err != nil {
			log.Fatal(err)
		}
	}()

	return ln
}
