package main

import (
    "fmt"
    "io"
    "strings"
    "time"
    "net/http"
    "github.com/sqweek/dialog"
    "github.com/getlantern/systray"
)

const VersionFormat = "20060102T150405MST"

func main() {
    systray.Run(onReady, onExit)
}

func loadAllServices() {

}

func onReady() {
  systray.SetIcon(flaarumLogoBytes)
  systray.SetTitle("Flaarum: a comfortable database")
  reload := systray.AddMenuItem("Reload Services", "Reloads all Services")
  updates := systray.AddMenuItem("Updates", "Check for updates")
  systray.AddSeparator()
  mQuit := systray.AddMenuItem("Quit", "Quits this app")

  go func() {
    for {
      select {
      case <-reload.ClickedCh:
        loadAllServices()
      // case <-hcmcTime.ClickedCh:
      //   timezone = "Asia/Ho_Chi_Minh"
      // case <-sydTime.ClickedCh:
      //   timezone = "Australia/Sydney"
      // case <-gdlTime.ClickedCh:
      //   timezone = "America/Mexico_City"
      // case <-sfTime.ClickedCh:
      //   timezone = "America/Los_Angeles"
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

  if hnv == true {
    dialog.Message("%s", "Please visit 'https://sae.ng/flaarumtuts/install' to download a new installer.").Title("Flaarum has an Update").YesNo()
  }

}
