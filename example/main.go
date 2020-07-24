package main

import (
	"sync"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go runSystray(&wg)
	wg.Wait()
}

func runSystray(wg *sync.WaitGroup) {
	onReady := func() {
		systray.SetIcon(icon.Data)
		systray.SetTitle("Awsome App")
		systray.SetTooltip("Lantern")
		mQuitOrig := systray.AddMenuItem("Quit", "Quit App")
		go func() {
			<-mQuitOrig.ClickedCh
			systray.Quit()
		}()
	}

	onExit := func() {
		wg.Done()
	}

	systray.Run(onReady, onExit)
}
