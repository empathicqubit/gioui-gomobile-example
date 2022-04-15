package main

import (
	"strconv"
	"strings"

	"github.com/JuulLabs-OSS/cbgo"
	"github.com/empathicqubit/giouibind/native"
)

var _ native.INativeBridge = (*NativeBridge)(nil)

type NativeBridge struct {
	fd       int
	central  *cbgo.CentralManager
	delegate *MyDelegate
}

// str2ba converts MAC address string representation to little-endian byte array
func str2ba(addr string) [6]byte {
	a := strings.Split(addr, ":")
	var b [6]byte
	for i, tmp := range a {
		u, _ := strconv.ParseUint(tmp, 16, 8)
		b[len(b)-1-i] = byte(u)
	}
	return b
}

func (god *NativeBridge) EnableBluetooth() bool {
	central := cbgo.NewCentralManager(&cbgo.ManagerOpts{})
	god.central = &central
	delegate := &MyDelegate{}
	god.delegate = delegate
	god.central.SetDelegate(delegate)

	return true
}

var localName string = ""

func (god *NativeBridge) WriteChar(data []byte) bool {
	if data == nil {
		return false
	}

	return god.delegate.writeChar(data)
}

func (god *NativeBridge) ConnectToDevice() {
}
