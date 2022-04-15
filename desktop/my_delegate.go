package main

import (
	"log"

	"github.com/JuulLabs-OSS/cbgo"
)

type MyDelegate struct {
	cbgo.CentralManagerDelegateBase
	cbgo.PeripheralDelegateBase
	connectedPeripheral *cbgo.Peripheral
	service             *cbgo.Service
	readChr             cbgo.Characteristic
	writeChr            cbgo.Characteristic
	encryptedChr        cbgo.Characteristic
	pairingChr          cbgo.Characteristic
	peripheral          *cbgo.Peripheral
	finishedConnect     bool
}

var shortServiceUuid, _ = cbgo.ParseUUID16("FIXME")
var longServiceUuid, _ = cbgo.ParseUUID("FIXME")

var writeUuid, _ = cbgo.ParseUUID("FIXME")
var readUuid, _ = cbgo.ParseUUID("FIXME")

func (d *MyDelegate) writeChar(data []byte) bool {
	if d.connectedPeripheral == nil {
		return false
	}

	d.connectedPeripheral.WriteCharacteristic(data, d.writeChr, true)

	return true
}

func (d *MyDelegate) readChar() []byte {
	if d.connectedPeripheral == nil {
		return []byte{}
	}

	d.connectedPeripheral.ReadCharacteristic(d.readChr)

	return d.readChr.Value()
}

func (d *MyDelegate) CentralManagerDidUpdateState(cmgr cbgo.CentralManager) {
	if cmgr.State() == cbgo.ManagerStatePoweredOn {
		log.Println("Start scanning")
		go func() {
			nativeBridge.central.Scan([]cbgo.UUID{shortServiceUuid, longServiceUuid}, &cbgo.CentralManagerScanOpts{
				AllowDuplicates:       false,
				SolicitedServiceUUIDs: []cbgo.UUID{longServiceUuid, shortServiceUuid},
			})
		}()
	}
}

func (d *MyDelegate) DidDiscoverPeripheral(cm cbgo.CentralManager, prph cbgo.Peripheral,
	advFields cbgo.AdvFields, rssi int) {
	log.Println("Found peripheral", prph.Name())
	nativeBridge.central.Connect(prph, nil)
}

func (d *MyDelegate) DidConnectPeripheral(cm cbgo.CentralManager, prph cbgo.Peripheral) {
	if d.connectedPeripheral == nil {
		prph.SetDelegate(d)
		prph.DiscoverServices([]cbgo.UUID{longServiceUuid, shortServiceUuid})
	}
}

func (d *MyDelegate) DidFailToConnectPeripheral(cm cbgo.CentralManager, prph cbgo.Peripheral, err error) {
	log.Printf("failed to connect: %v", err)
}

func (d *MyDelegate) DidDisconnectPeripheral(cm cbgo.CentralManager, prph cbgo.Peripheral, err error) {
	log.Printf("peripheral disconnected: %v", err)
}

func (d *MyDelegate) DidDiscoverServices(prph cbgo.Peripheral, err error) {
	if err != nil || len(prph.Services()) == 0 {
		log.Println("Error discovering services", err)
		dumbApp.FinishedConnect(false, "")
		return
	}

	d.service = &prph.Services()[0]
	prph.DiscoverCharacteristics([]cbgo.UUID{writeUuid, readUuid, encryptedUuid, pairingUuid}, *d.service)
}

func (d *MyDelegate) DidDiscoverCharacteristics(prph cbgo.Peripheral, svc cbgo.Service, err error) {
	if svc.UUID().String() != d.service.UUID().String() {
		return
	}

	if err != nil || len(svc.Characteristics()) == 0 {
		log.Println("No characteristics")
		dumbApp.FinishedConnect(false, "")
		return
	}

	foundCount := 0
	for _, chr := range d.service.Characteristics() {
		log.Println("Found char", chr.UUID().String())

		uuid := chr.UUID().String()
		if uuid == writeUuid.String() {
			d.writeChr = chr
			foundCount += 1
		} else if uuid == readUuid.String() {
			d.readChr = chr
			foundCount += 1
                }
	}

	if foundCount < 2 || d.service == nil {
		log.Println("Couldn't find matching service or characteristics")
		dumbApp.FinishedConnect(false, "")
		return
	}

	nativeBridge.central.StopScan()
	d.connectedPeripheral = &prph

	log.Println("Finished connect. Writing test command")

	d.connectedPeripheral.SetNotify(true, d.readChr)
}

func (d *MyDelegate) DidDiscoverDescriptors(prph cbgo.Peripheral, chr cbgo.Characteristic, err error) {
}

func (d *MyDelegate) DidUpdateValueForCharacteristic(prph cbgo.Peripheral, chr cbgo.Characteristic, err error) {
	if err != nil {
		log.Printf("Error getting descriptor value: %s", chr.UUID().String())
		return
	}

	if chr.Value() == nil {
		return
	}
	data := chr.Value()

	if chr.UUID().String() != d.readChr.UUID().String() {
		return
	}

	dumbApp.BluetoothGotData(data)

	if !d.finishedConnect {
		d.finishedConnect = true
		dumbApp.FinishedConnect(true, d.connectedPeripheral.Name())
	}
}

func (d *MyDelegate) DidUpdateValueForDescriptor(prph cbgo.Peripheral, dsc cbgo.Descriptor, err error) {
	if dsc.Characteristic().Service().UUID().String() != d.service.UUID().String() {
		return
	}

	d.DidUpdateValueForCharacteristic(prph, dsc.Characteristic(), err)
}
