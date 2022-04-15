package giouibind

import (
	"fmt"
	"image/color"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/empathicqubit/giouibind/native"

	"log"
)

const (
	screenW = 320
	screenH = 320
)

func (a *AppState) FinishedConnect(success bool, name string) {
	a.connected = success
	a.connecting = false
	a.deviceName = name
	log.Printf("Device %s connected: %t\n", name, success)
	a.window.Invalidate()
}

func (a *AppState) FinishedDisconnect() {
}

func (app *AppState) BluetoothGotData(data []byte) {
	log.Printf("Received: %v\n", data)
}

func (app *AppState) Update() error {
	if !app.connected && !app.connecting {
		app.connecting = true
		app.nativeBridge.ConnectToDevice()
		return nil
	}

	if !app.connected {
		return nil
	}

	return nil
}

func (app *AppState) Load() error {
	app.loaded = true

	return nil
}

func (app *AppState) InventGod(god native.INativeBridge) error {
	log.Println("God has been invented")
	app.nativeBridge = god

	if !app.bluetoothEnabled {
		app.nativeBridge.EnableBluetooth()
		app.bluetoothEnabled = true
	}

	return nil
}

func (app *AppState) RunApp(w *app.Window) error {
	app.window = w
	th := material.NewTheme(gofont.Collection())
	th.TextSize = unit.Sp(8)
	var ops op.Ops
	for {
		go app.Update()
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			titleText := "Not connected"
			if app.connected {
				titleText = fmt.Sprintf("Connected to %s", app.deviceName)
			}
			title := material.H1(th, titleText)
			maroon := color.NRGBA{R: 127, G: 0, B: 0, A: 255}
			title.Color = maroon
			title.Alignment = text.Middle
			title.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}
