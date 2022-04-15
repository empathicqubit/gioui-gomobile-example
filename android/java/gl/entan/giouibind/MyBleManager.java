package gl.entan.giouibind;

import java.util.List;
import java.util.UUID;

import android.bluetooth.BluetoothDevice;
import android.bluetooth.BluetoothGatt;
import android.bluetooth.BluetoothGattCallback;
import android.bluetooth.BluetoothGattCharacteristic;
import android.bluetooth.BluetoothGattService;
import android.bluetooth.BluetoothProfile;
import android.content.Context;
import android.os.Handler;
import android.util.Log;
import gl.entan.giouibind.mobile.Mobile;
import no.nordicsemi.android.ble.BleManager;

public class MyBleManager extends BleManager {
    public static UUID shortServiceUuid = UUID.fromString("FIXME");
    public static UUID longServiceUuid = UUID.fromString("FIXME");

    public static UUID writeUuid = UUID.fromString("FIXME");
    public static UUID readUuid = UUID.fromString("FIXME");

    private BluetoothGattService service;
    private BluetoothGattCharacteristic writeChr;
    private BluetoothGattCharacteristic readChr;

    public MyBleManager(Context context) {
        super(context);
    }

    public MyBleManager(Context context, Handler handler) {
        super(context, handler);
    }

    public void readChar() {
        readCharacteristic(readChr)
            .with((d, data) -> {
                Mobile.bluetoothGotData(data.getValue());
            })
            .enqueue();
    }

    public boolean writeChar(byte[] data) {
        writeCharacteristic(writeChr, data, BluetoothGattCharacteristic.WRITE_TYPE_DEFAULT).enqueue();
        return true;
    }

    @Override
    protected BleManagerGattCallback getGattCallback() {
        return new MyGattCallbackImpl();
    }

    @Override
    public int getMinLogPriority() {
        return Log.WARN;
    }

    @Override
    public void log(int priority, String message) {
        Log.println(priority, "bananas", message);
    }

    private class MyGattCallbackImpl extends BleManagerGattCallback {
        @Override
        protected void onCharacteristicNotified(BluetoothGatt gatt, BluetoothGattCharacteristic characteristic) {
            Mobile.bluetoothGotData(characteristic.getValue());
            super.onCharacteristicNotified(gatt, characteristic);
        }

        @Override
        protected boolean isRequiredServiceSupported(BluetoothGatt gatt) {
            BluetoothGattService s = gatt.getService(shortServiceUuid);
            if(s == null) {
                s = gatt.getService(longServiceUuid);
            }
            if(s == null) {
                Log.i("bananas", "Could not find GATT service");
                Mobile.finishedConnect(false, "");
                return false;
            }

            service = s;
            writeChr = s.getCharacteristic(writeUuid);
            readChr = s.getCharacteristic(readUuid);

            return service != null
                && writeChr != null 
                && readChr != null;
        }

        @Override
        protected void initialize() {
            requestMtu(165)
                .done(d -> {
                    enableNotifications(readChr).enqueue();

                    writeCharacteristic(writeChr, new byte[] { 0x04, (byte)0x95, 0x06, 0x03 }, BluetoothGattCharacteristic.WRITE_TYPE_DEFAULT)
                    .done(g -> {
                        Mobile.finishedConnect(true, g.getName());
                    })
                    .enqueue();
                })
                .enqueue();
        }

        @Override
        protected void onServicesInvalidated() {
            readChr = writeChr = null;
        }
    }

}