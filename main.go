package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/getlantern/systray"
	"github.com/sqweek/dialog"
)

const VersionFormat = "20060102T150405MST"

var StopRunningFlStore context.CancelFunc

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(flaarumLogoBytes)
	systray.SetTitle("Flaarum: a comfortable database")
	updates := systray.AddMenuItem("Updates", "Check for updates")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quits this app")

	ctx, cancel := context.WithCancel(context.Background())
	StopRunningFlStore = cancel

	go func() {
		exec.CommandContext(ctx, "C:\\Program Files (x86)\\Flaarum\\flstore.exe").Run()
	}()

	go func() {
		for {
			select {

			case <-updates.ClickedCh:
				checkAndNofityOfUpdates()

			case <-mQuit.ClickedCh:
				systray.Quit()
				return

			}
		}
	}()
}

func onExit() {
	StopRunningFlStore()
}

func checkAndNofityOfUpdates() {
	newVersionStr := ""
	resp, err := http.Get("https://sae.ng/static/wapps/flaarum.txt")
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err == nil && resp.StatusCode == 200 {
			newVersionStr = string(body)
		}
	}

	newVersionStr = strings.TrimSpace(newVersionStr)
	currentVersionStr = strings.TrimSpace(currentVersionStr)

	hnv := false
	if newVersionStr != "" && newVersionStr != currentVersionStr {
		time1, err1 := time.Parse(VersionFormat, newVersionStr)
		time2, err2 := time.Parse(VersionFormat, currentVersionStr)

		if err1 == nil && err2 == nil && time2.Before(time1) {
			hnv = true
		}
	}

	if hnv {
		dialog.Message("%s", "Please visit 'https://sae.ng/flaarumtuts/install' to download a new installer.").Title("Flaarum has an Update").YesNo()
	}

}
