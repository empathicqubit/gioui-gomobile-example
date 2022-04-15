package giouibind

import (
	"gioui.org/app"
	"github.com/empathicqubit/giouibind/native"
)

type AppSettings struct {
}

type AppState struct {
	settings         *AppSettings
	window           *app.Window
	bluetoothEnabled bool
	loaded           bool
	connected        bool
	connecting       bool
	initted          bool
	deviceName       string
	nativeBridge     native.INativeBridge
}
