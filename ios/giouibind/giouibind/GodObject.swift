//
//  GodObject.swift
//  giouibind
//
//  Created by Jessica Fleming on 11.04.22.
//

import Foundation
import Mobile
import ExternalAccessory
import os
import CoreBluetooth

// FIXME This should come from gocode
let shortServiceUuid = CBUUID(string: "FIXME")
let longServiceUuid = CBUUID(string: "FIXME")

let writeUuid = CBUUID(string: "FIXME")
let readUuid = CBUUID(string: "FIXME")

class MyDelegate : NSObject, CBCentralManagerDelegate, CBPeripheralDelegate {
    var service : CBService?
    var writeChr : CBCharacteristic?
    var readChr : CBCharacteristic?
    var connectedPeripheral : CBPeripheral?
    var finishedConnect = false
    
    override init() {
        super.init()
    }
    
    func peripheral(_ peripheral: CBPeripheral, didDiscoverServices error:
        Error?) {
        if error != nil || (peripheral.services?.isEmpty ?? true) {
            os_log("Error discovering services", log: .default, type: .info)
            MobileFinishedConnect(false, "")
            return
        }
        
        service = peripheral.services![0]
        peripheral.discoverCharacteristics([readUuid, writeUuid], for: service!)

    }
    func peripheral(_ peripheral: CBPeripheral, didUpdateValueFor descriptor: CBDescriptor, error: Error?) {
        if descriptor.characteristic == nil || descriptor.characteristic?.service?.uuid.uuidString != service?.uuid.uuidString {
            return
        }
        
        self.peripheral(peripheral, didUpdateValueFor: descriptor.characteristic!, error: error)
    }
    
    func peripheral(_ peripheral: CBPeripheral, didUpdateValueFor characteristic: CBCharacteristic, error: Error?) {
                
        if error != nil {
            os_log("Error getting descriptor value: %s", log: .default, type: .info, characteristic.uuid.uuidString)
            return
        }
        
        if characteristic.value == nil {
            return
        }
        let data = characteristic.value!
 
        if characteristic.uuid.uuidString != readChr?.uuid.uuidString {
            return
        }
        
        MobileBluetoothGotData(data)
    }
    
    func centralManager(_ central: CBCentralManager, didConnect peripheral: CBPeripheral) {
        if connectedPeripheral == nil {
            peripheral.discoverServices([longServiceUuid, shortServiceUuid])
            return
        }
    }
    
    func centralManager(_ central: CBCentralManager, didDisconnectPeripheral peripheral: CBPeripheral, error: Error?) {
        connectedPeripheral = nil
    }
    
    func peripheral(_ peripheral: CBPeripheral, didDiscoverCharacteristicsFor s: CBService, error: Error?) {
        if s.uuid.uuidString != service?.uuid.uuidString {
            return
        }
        
        if(error != nil || (service?.characteristics?.isEmpty ?? true)) {
            os_log("No chars", log: .default, type: .info)
            MobileFinishedConnect(false, "")
            return
        }
        
        var foundCount = 0
        for chr in service!.characteristics! {
            os_log("Found char: %s", log: .default, type: .info, chr.uuid.uuidString)

            let uuid = chr.uuid.uuidString
            if uuid == writeUuid.uuidString {
                writeChr = chr
                foundCount += 1
            } else if uuid == readUuid.uuidString {
                readChr = chr
                foundCount += 1
            }
        }
        
        if foundCount < 2 || service == nil {
            os_log("Couldn't find matching service or characteristics", log: .default, type: .info)
            MobileFinishedConnect(false, "")
            return
        }
        
        connectedPeripheral = peripheral
        
        os_log("Finished connect. Writing test command", log: .default, type: .info)
        
        peripheral.setNotifyValue(true, for: readChr!)
        
        if !finishedConnect {
            finishedConnect = true
            MobileFinishedConnect(true, peripheral.name)
        }
    }
    
    func writeChar(_ d: Data) -> Bool {
        if connectedPeripheral == nil {
            return false
        }
                
        connectedPeripheral!.writeValue(d, for: writeChr! , type: CBCharacteristicWriteType.withResponse)
        
        return true
    }
    
    func readChar() -> Data {
        if connectedPeripheral == nil {
            return Data()
        }
        
        let ret = readChr!.value ?? Data()
                
        return ret
    }
    
    func centralManagerDidUpdateState(_ central: CBCentralManager){
        os_log("Central Manager intialized", log: .default, type: .info)
        
        switch central.state{
        case CBManagerState.unauthorized:
            os_log("Central Manager unauthorized", log: .default, type: .info)
        case CBManagerState.poweredOff:
            os_log("Central Manager powered off", log: .default, type: .info)
        case CBManagerState.poweredOn:
            os_log("Central Manager powered on", log: .default, type: .info)
        default:break
        }
    }
}

class GodObject : NSObject, MobileIGodObjectProtocol {
    func readChar() -> Data? {
        return delegate.readChar()
    }
    
    func writeChar(_ data: Data?) -> Bool {
        if data == nil {
            return false
        }
        
        return delegate.writeChar(data!)
    }
    
    let delegate : MyDelegate
    let manager : CBCentralManager
    var peris : [CBPeripheral] = []
        
    override init() {
        delegate = MyDelegate()
        manager = CBCentralManager(delegate: delegate, queue: DispatchQueue(label: "BT_queue"))
        super.init()
    }
    
    func connectToDevice() {
        let peris = manager.retrieveConnectedPeripherals(withServices: [shortServiceUuid, longServiceUuid])
        
        if(peris.isEmpty) {
            os_log("No connected devices", log: .default, type: .info)
            MobileFinishedConnect(false, "")
            return
        }
        
        self.peris = peris
        for peri in peris {
            os_log("Found device: %s", log: .default, type: .info, peri.name!)
            peri.delegate = delegate
            manager.connect(peri)
        }
    }
    
    func enableBluetooth() -> Bool {
        return true
    }
}
