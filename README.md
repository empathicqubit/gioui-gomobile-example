# Minimal Gio-Powered Android/iOS/desktop app

This project serves as an example how to build minimal working Gio-powered
Android/iOS/desktop application **without** Android Studio or Gradle. The example
was intended to connect to a BLE device with a specific service identifier.
Please search for FIXME in the project and replace information as appropriate.
Alternatively you can remove this stuff and replace it with your own.

The interface INativeBridge in native/native.go contains the methods which
bridge go to either Java on Android, ObjC/Swift on iOS, or more Go classes
on desktop.

Makefile targets:
* **build**: Build for both iOS and Android. Only works on Mac OS because of iOS
* **build-ios**, **build-android**: Build for single platform
* **install-ios**, **install-android**: Install to a real device. Your signing
team ID in /ios/export-options.plist must be correct.
* **run-android**: Tells the application to start on Android
* **debug-android**: Doesn't work yet.

## Requirements

To be able to build, following dependencies are
required:

* Android SDK version 32, be sure to set ANDROID\_HOME environment variable
* Android NDK version 24.0.8215888
* Android build-tools version 30.0.3
* Java on your PATH
* gomobile `go install "golang.org/x/mobile/cmd/gomobile"`

## Description of different project files
* **build**: All the build stuff goes here
* **android**: All the Android-specific files here. If you want to include
images in your application, it's probably better to do that in Go 
with //go:embed, instead of including images in the res folder. GodObject
is the main class which implements the bridge with Go. The Gio classes are
overridden to work around some issues with using Gio with `gomobile bind`
* **ios**: XCode project which contains iOS specific INativeBridge implementation.
* **mobile**: The entrypoint for `gomobile bind`
* **mobile/gio**: A stub entrypoint for Gio. You shouldn't need to change this.
* **desktop**: Entrypoint for the desktop application. The example only works
on Mac OS because the cbgo library is made for Mac OS bluetooth only, but you
can replace it with whatever you want and that should work on MacOS/Linux/Windows.
* **gio**: Stubs for Gio packages which fail to build on Linux when targeting Android.

## Credits

Thanks to [seletskiy/ebiten-android-minimal](https://github.com/seletskiy/ebiten-android-minimal) and [GioUI](https://gioui.org/)
