package gl.entan.giouibind;

import java.util.ArrayList;
import java.util.List;

import android.Manifest.permission;
import android.app.Activity;
import android.bluetooth.BluetoothAdapter;
import android.bluetooth.BluetoothDevice;
import android.bluetooth.BluetoothProfile;
import android.bluetooth.BluetoothSocket;
import android.content.Context;
import android.content.pm.PackageManager;
import android.util.Log;
import gl.entan.giouibind.mobile.IGodObject;
import gl.entan.giouibind.mobile.Mobile;

public class GodObject implements IGodObject {
    private static GodObject instance = null;

    private BluetoothSocket socket = null;
    private Context context = null;
    private List<MyBleManager> managers;

    private GodObject(Context context) {
        this.context = context;
        this.managers = new ArrayList<>();
    }

    public static GodObject getGodObject(Context context) {
        if(GodObject.instance == null) {
            GodObject.instance = new GodObject(context);
        }

        return GodObject.instance;
    }
    
    public void disconnectFromDevice() {
        if(this.socket != null) {
            try {
                this.socket.close();
            }
            catch (Exception e) {
                Log.e("bananas", "Couldn't close socket", e);
            }
        }
    }

    public void connectToDevice() {
        GodObject self = this;

        this.disconnectFromDevice();

        BluetoothAdapter.getDefaultAdapter().getProfileProxy(this.context, new BluetoothProfile.ServiceListener() {
            @Override
            public void onServiceDisconnected(int arg0) {
            }

            @Override
            public void onServiceConnected(int profile, BluetoothProfile proxy) {
                Log.i("bananas", "Called service connected");

                for (BluetoothDevice device : proxy.getConnectedDevices()) {
                    MyBleManager manager = new MyBleManager(context);

                    manager
                        .connect(device)
                        .retry(20)
                        .timeout(5000)
                        .useAutoConnect(true)
                        .enqueue();

                    self.managers.add(manager);
                }

                BluetoothAdapter.getDefaultAdapter().closeProfileProxy(profile, proxy);
            }
        }, BluetoothProfile.HEADSET);
    }

    @Override
    public boolean enableBluetooth() {
        Context context = this.context;

        if(context.checkSelfPermission(permission.BLUETOOTH) == PackageManager.PERMISSION_DENIED) {
            ((Activity)context).requestPermissions(new String[] { permission.BLUETOOTH }, 0);
            return false;
        }

        if(context.checkSelfPermission(permission.ACCESS_BACKGROUND_LOCATION) == PackageManager.PERMISSION_DENIED) {
            ((Activity)context).requestPermissions(new String[] { permission.ACCESS_BACKGROUND_LOCATION }, 0);
            return false;
        }

        return true;
    }

    @Override
    public boolean writeChar(byte[] data) {
        for(MyBleManager manager: this.managers) {
            if(manager.isReady()) {
                manager.writeChar(data);
                break;
            }
        }
        return false;
    }
}
