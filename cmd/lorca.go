package cmd

import (
	"github.com/zserge/lorca"
	"log"
)

func StartLorca(url string, api string) (lorca.UI, error) {
	args := []string{"--remote-allow-origins=*", "--disable-infobars", "--disable-features=automation"}
	//if runtime.GOOS == "linux" {
	//	args = append(args, "--class=Lorca")
	//}
	ui, err := lorca.New("", "", 1240, 826, args...)
	if err != nil {
		return ui, err
	}

	// A simple way to know when UI is ready (uses body.onload event in JS)
	err = ui.Bind("start", func() {
		log.Println("UI is ready")
	})
	if err != nil {
		return ui, err
	}

	err = ui.Bind("getApiRoot", func() string {
		return api
	})
	if err != nil {
		return ui, err
	}

	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: ui.Load("data:text/html," + url.PathEscape(html))
	if err = ui.Load(url); err != nil {
		return ui, err
	}

	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	ui.Eval(`
		console.log("Hello, world!");
		console.log('Multiple values:', [1, false, {"x":5}]);
	`)

	return ui, nil
}
