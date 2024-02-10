package main

import (
	"awesomeProject/cmd"
	"embed"
	"fmt"
	"github.com/zserge/lorca"
	"io/fs"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

//go:embed web/build
var webFs embed.FS

func main() {
	webSubFs, err := fs.Sub(webFs, "web/build")
	if err != nil {
		panic(err)
	}

	lnFile := cmd.ServeWebBuild(webSubFs)
	defer func(ln net.Listener) {
		if err = ln.Close(); err != nil {
			log.Fatal("File Server Shutdown:", err)
		}
	}(lnFile)

	log.Printf("main: file server started. @%s", lnFile.Addr())

	lnBk := cmd.StartServer()
	defer func(ln net.Listener) {
		if err = ln.Close(); err != nil {
			log.Fatal("Backend Server Shutdown:", err)
		}
	}(lnBk)

	log.Printf("main: backend server started. @%s", lnBk.Addr())

	ui := cmd.StartLorca(
		fmt.Sprintf("http://%s", lnFile.Addr()),
		fmt.Sprintf("http://%s", lnBk.Addr()),
	)
	defer func(ui lorca.UI) {
		if err = ui.Close(); err != nil {
			log.Fatal("Lorca UI Close:", err)
		}
	}(ui)

	log.Printf("main: all started.")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
	case <-ui.Done():
	}

	log.Printf("main: done. exiting")
}
