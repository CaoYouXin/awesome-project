package main

import (
	"awesomeProject/cmd"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/signal"
)

//go:embed web/build
var webFs embed.FS

func main() {
	webSubFs, err := fs.Sub(webFs, "web/build")
	if err != nil {
		panic(err)
	}

	ln := cmd.ServeWebBuild(webSubFs)
	defer ln.Close()

	////cmd.StartServer()
	ui := cmd.StartLorca(fmt.Sprintf("http://%s", ln.Addr()))
	defer ui.Close()

	log.Printf("main: started.")

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Printf("main: done. exiting")
}
