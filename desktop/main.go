package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/unit"
	"github.com/empathicqubit/giouibind"
	"github.com/empathicqubit/giouibind/native"
)

var nativeBridge *NativeBridge = &NativeBridge{}
var dumbApp *giouibind.AppState = &giouibind.AppState{}

func main() {
	var igod native.INativeBridge = nativeBridge
	log.Println("Called main")
	err := dumbApp.Load()
	if err != nil {
		panic(err)
	}
	err = dumbApp.InventGod(igod)
	if err != nil {
		panic(err)
	}

	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(450), unit.Dp(800)),
		)
		err := dumbApp.RunApp(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}
