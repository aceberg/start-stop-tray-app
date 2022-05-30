package main

import (
	"fmt"
	"os/exec"
	"github.com/getlantern/systray"
	"github.com/aceberg/start-stop-tray-app/iconon"
	"github.com/aceberg/start-stop-tray-app/iconoff"
)

var (
	stopIcon = iconoff.Data
	startIcon = iconon.Data
	mycliApp = "myAppName"
)

func main() {
	systray.Run(onReady, onExit)
}

func runCmd(arg0 string) {
	cmd := exec.Command(mycliApp, arg0)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
}

func onReady() {
	systray.SetIcon(stopIcon)
	systray.SetTitle("Start/Stop")
	systray.SetTooltip("Start/Stop")

	mStart := systray.AddMenuItem("Start", "Start")
	mStop := systray.AddMenuItem("Stop", "Stop")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	for {
		select {
		case <-mStart.ClickedCh:
			systray.SetIcon(startIcon)
			fmt.Println("Start")
			runCmd("start")
		case <-mStop.ClickedCh:
			systray.SetIcon(stopIcon)
			fmt.Println("Stop")
			runCmd("stop")
		case <-mQuit.ClickedCh:
			fmt.Println("Quit now...")
			runCmd("stop")
			systray.Quit()
			return
		}
	}
}

func onExit() {}
