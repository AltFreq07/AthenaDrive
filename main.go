package main

import (
	"fmt"
	"os"
	"time"

	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/getlantern/systray"
)

var exit chan bool = make(chan bool)
var fyneApp fyne.App = app.New()

func main() {
	fmt.Println("Starting app")
	checkRClone()
	checkConfig()
	getMountLocation()
	go systray.Run(onReady, onExit)
	go mountDrive()
	for {
		fmt.Println("Sleep")
		time.Sleep(time.Second)
		select {
		case <-exit:
			fmt.Println("Exiting")
			os.Exit(0)
		}
	}
}
