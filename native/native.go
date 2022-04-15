// Interfaces for the binding bridge. This should only contain a bare minimum
// of functions to communicate with the host and should avoid storing anything
package native

type INativeBluetooth interface {
	WriteChar([]byte) bool
	ConnectToDevice()
	EnableBluetooth() bool
}

type INativeBridge interface {
	INativeBluetooth
}
