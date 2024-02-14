package main

import (
	"awesomeProject/cmd"
	"awesomeProject/db"
	"database/sql"
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

func runMain() {
	if err := db.InitInst(); err != nil {
		log.Fatal(err)
		return
	}
	defer func(Inst *sql.DB) {
		err := Inst.Close()
		if err != nil {
			fmt.Println("DB conn close failed.")
		}
	}(db.Inst)

	webSubFs, err := fs.Sub(webFs, "web/build")
	if err != nil {
		log.Fatal(err)
		return
	}

	lnFile := cmd.ServeWebBuild(webSubFs)
	defer func(ln net.Listener) {
		if err = ln.Close(); err != nil {
			log.Fatal("File Server Shutdown:", err)
		}
	}(lnFile)

	log.Printf("main: file server started. @%s", lnFile.Addr())

	lnBk, err := cmd.StartServer("0")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(ln net.Listener) {
		if err = ln.Close(); err != nil {
			log.Fatal("Backend Server Shutdown:", err)
		}
	}(lnBk)

	log.Printf("main: backend server started. @%s", lnBk.Addr())

	ui, err := cmd.StartLorca(
		fmt.Sprintf("http://%s", lnFile.Addr()),
		fmt.Sprintf("http://%s", lnBk.Addr()),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(ui lorca.UI) {
		if err = ui.Close(); err != nil {
			log.Fatal("Lorca UI Close:", err)
		}
	}(ui)

	log.Printf("main: lorca ui started.")

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

func runBk() {
	if err := db.InitInst(); err != nil {
		log.Fatal(err)
		return
	}
	defer func(Inst *sql.DB) {
		err := Inst.Close()
		if err != nil {
			fmt.Println("DB conn close failed.")
		}
	}(db.Inst)

	lnBk, err := cmd.StartServer("8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func(ln net.Listener) {
		if err = ln.Close(); err != nil {
			log.Fatal("Backend Server Shutdown:", err)
		}
	}(lnBk)

	log.Printf("main: backend server started. @%s", lnBk.Addr())

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
	}

	log.Printf("main: done. exiting")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		runMain()
	} else if args[0] == "-bk" {
		runBk()
	} else {
		fmt.Println("Unknown Args", args)
	}
}
