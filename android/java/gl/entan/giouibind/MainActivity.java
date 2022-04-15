package gl.entan.giouibind;

import org.gioui.GioView;

import android.app.Activity;
import android.app.ActivityManager;
import android.content.Context;
import android.content.res.Configuration;
import android.graphics.BitmapFactory;
import android.os.Build;
import android.os.Bundle;
import android.util.Log;
import go.Seq;
import gl.entan.giouibind.mobile.Mobile;

public class MainActivity extends Activity {
  private GioView view;

	@Override public void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.main);
		Context context = getApplicationContext();
		Seq.setContext(context);

		this.view = findViewById(R.id.gioview);

		if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.LOLLIPOP) {
			this.setTaskDescription(
				new ActivityManager.TaskDescription(
					null, // Leave the default title.
					BitmapFactory.decodeResource(getResources(), R.drawable.icon)
			));
		}
	}

	@Override public void onDestroy() {
		view.destroy();
		super.onDestroy();
	}

	@Override public void onStart() {
		super.onStart();
		view.start();
	}

	@Override public void onStop() {
		view.stop();
		super.onStop();
	}

	@Override public void onConfigurationChanged(Configuration c) {
		super.onConfigurationChanged(c);
		view.configurationChanged();
	}

	@Override public void onLowMemory() {
		super.onLowMemory();
		GioView.onLowMemory();
	}

	@Override public void onBackPressed() {
		if (!view.backPressed())
			super.onBackPressed();
	}
}
