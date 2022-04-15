package mobile

import (
	"log"
	"os"

	"gioui.org/app"
	"github.com/empathicqubit/giouibind"
	"github.com/empathicqubit/giouibind/native"
)

type IGodObject interface {
	native.INativeBridge
}

var dumbApp *giouibind.AppState = &giouibind.AppState{}

func init() {
	log.Println("Called gomobile entrypoint")
	Main()
}

func Main() {
	log.Println("Called main")
	err := dumbApp.Load()
	if err != nil {
		panic(err)
	}

	go func() {
		w := app.NewWindow()
		err := dumbApp.RunApp(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func FinishedConnect(success bool, name string) {
	dumbApp.FinishedConnect(success, name)
}

func BluetoothGotData(data []byte) {
	dumbApp.BluetoothGotData(data)
}

func InventGod(god IGodObject) {
	log.Println("Inventing God...")
	dumb, ok := god.(native.INativeBridge)
	if !ok {
		log.Println("God is not okay!")
	} else {
		log.Println("God is fine and dandy!")
	}
	dumbApp.InventGod(dumb)
}
