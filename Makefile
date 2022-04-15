# FIXME Explain wth we're doing
MIN_SDK_VERSION ?= 21
ANDROID_SDK_VERSION ?= 30.0.3
ANDROID_HOME ?= $(HOME)/Android/Sdk
ANDROID_NDK_HOME ?= $(ANDROID_HOME)/ndk/24.0.8215888
MAX_SDK_VERSION ?= 32

GDB_PORT ?= 4334

# Target package name.
ANDROID_PACKAGE ?= gl.entan.giouibind

# Settings for keystore which is used to sign APK.
KEYS_DN=DC=gl,CN=entan
KEYS_PASS=123456
KEYS_VALIDITY=365
KEYS_ALGORITHM=RSA
KEYS_SIZE=2048

# Target build dir. Resulting APK file will be available here.
BUILD_DIR=build

_BUILD_TOOLS=$(ANDROID_HOME)/build-tools/$(ANDROID_SDK_VERSION)
_ANDROID_JAR_PATH=$(ANDROID_HOME)/platforms/android-$(MAX_SDK_VERSION)/android.jar
_SWIFT_SRC=$(shell find ios -iname '*.swift')
_JAVA_SRC=$(shell find android/java -iname '*.java')
_GO_SRC=$(shell find . -iname '*.go')
_JAVA_ROOT_PATH=android/java/$(subst .,/,$(ANDROID_PACKAGE))
_ADB_PATH=$(ANDROID_HOME)/platform-tools/adb
_JARS=build/jars/ble/classes.jar build/jars/go-android/classes.jar


build: build-ios build-android

install-ios: build/ios-app.ipa
	ideviceinstaller -i build/ios-app.ipa/giouibind.ipa

ios: build-ios
build-ios: build/ios-app.ipa

android: build-android
build-android: build/android-app.apk

run-android: install-android
	$(_ADB_PATH) shell am start -n $(ANDROID_PACKAGE)/.MainActivity
	$(_ADB_PATH) shell 'while ! dumpsys window windows | grep -o "$(ANDROID_PACKAGE)" 2>&1 > /dev/null ; do sleep 1 ; done'

debug-android: run-android
#FIXME Delve
	$(_ADB_PATH) push $(ANDROID_NDK_HOME)/prebuilt/android-arm64/gdbserver/gdbserver /data/local/tmp
	$(_ADB_PATH) shell "chmod 777 /data/local/tmp/gdbserver"
	$(_ADB_PATH) forward tcp:$(GDB_PORT) tcp:$(GDB_PORT)
	$(_ADB_PATH) shell 'su -c killall gdbserver || exit 0'
	$(_ADB_PATH) shell 'su -c set enforce 0'
	$(_ADB_PATH) shell 'su -c /data/local/tmp/gdbserver :$(GDB_PORT) --attach $$(ps -A -o NAME,PID | grep "$(ANDROID_PACKAGE)" | cut -F 2)'

install-android: build-android
	$(_ADB_PATH) install build/android-app.apk

# Initialize keystore to sign APK.
keys.store:
	keytool -genkeypair \
		-validity $(KEYS_VALIDITY) \
		-keystore $@ \
		-keyalg $(KEYS_ALGORITHM) \
		-keysize $(KEYS_SIZE) \
		-storepass $(KEYS_PASS) \
		-keypass $(KEYS_PASS) \
		-dname $(KEYS_DN) \
		-deststoretype pkcs12

build/Mobile.xcframework: $(_GO_SRC)
	CGO_ENABLED=1 GO386=softfloat gomobile bind \
		-target ios \
		-o "$@" \
		github.com/empathicqubit/giouibind/mobile

build/ios-app.ipa: build/ios.xcarchive
	xcodebuild -allowProvisioningUpdates -exportArchive -archivePath build/ios.xcarchive -exportOptionsPlist ios/export-options.plist -exportPath "$@"

build/ios.xcarchive: $(_SWIFT_SRC) build/Mobile.xcframework
	xcodebuild -project ios/giouibind/giouibind.xcodeproj -scheme giouibind -sdk iphoneos -configuration AppStoreDistribution archive -archivePath "$@"

build/jars/ble/classes.jar:
	@mkdir -p build/jars/ble

	curl -L -o "build/ble.aar" https://repo1.maven.org/maven2/no/nordicsemi/android/ble/2.4.0/ble-2.4.0.aar
	cd build/jars/ble && unzip -o ../../ble.aar
	touch "$@"

build/jars/go-android/classes.jar: $(_GO_SRC) $(ANDROID_NDK_HOME)
	@mkdir -p build/jars/go-android

	CGO_ENABLED=1 ANDROID_NDK_HOME=$(ANDROID_NDK_HOME) GO386=softfloat gomobile bind \
		-target android \
		-javapkg $(ANDROID_PACKAGE) \
		-o "build/go-android.aar" \
		github.com/empathicqubit/giouibind/mobile

	# Unpack resulting AAR library to link it to APK during further stages.
	@unzip -o -qq "build/go-android.aar" -d build/jars/go-android
	@ln -sf jars/go-android/jni build/lib
	touch "$@"

# Collect resources and generate R.java.
$(_JAVA_ROOT_PATH)/R.java: $(wildcard android/res/*/*.*) android/AndroidManifest.xml $(_ANDROID_JAR_PATH)
	$(_BUILD_TOOLS)/aapt package \
		-f \
		-m \
		-J android/java \
		-M android/AndroidManifest.xml \
		-S android/res \
		-I $(_ANDROID_JAR_PATH)

# Generate a JAR suitable for code completion (less java.* classes)
build/jars/android-meta.jar: $(_ANDROID_JAR_PATH)
	@mkdir -p build/jars

	cp "$(_ANDROID_JAR_PATH)" "$@"
	zip -d "$@" 'java/*'

build/obj.jar: build/jars/android-meta.jar $(_JARS) $(_JAVA_SRC) $(_JAVA_ROOT_PATH)/R.java
	@mkdir -p build/obj

	javac \
		-source 8 \
		-target 8 \
		-d build/obj \
		-classpath $(_ANDROID_JAR_PATH):android/java:$(subst $() $(),:,$(_JARS)) \
		$(_JAVA_SRC)
	
	jar cvf "$@" -C build/obj/ .

# Convert compiled Java code into DEX file (required by Android).
build/classes.dex: build/d8.jar
	$(_BUILD_TOOLS)/dx \
		--dex \
		--min-sdk-version $(MIN_SDK_VERSION) \
		--output build/classes.dex \
		build/d8.jar

build/d8.jar: build/obj.jar
	$(_BUILD_TOOLS)/d8 \
		--output build/d8.jar \
		--classpath $(_ANDROID_JAR_PATH) \
		$(_JARS) \
		build/obj.jar

# Package everything into unaligned APK file.
build/app.apk.unaligned: build/classes.dex
	$(_BUILD_TOOLS)/aapt package \
		-f \
		-m \
		-F build/app.apk.unaligned \
		-M android/AndroidManifest.xml \
		-S android/res \
		-I $(_ANDROID_JAR_PATH)

	cd build && $(_BUILD_TOOLS)/aapt add \
		app.apk.unaligned \
		classes.dex \
		lib/*/*

# Align unaligned APK file and sign it using keystore.
build/android-app.apk: keys.store build/app.apk.unaligned
	$(_BUILD_TOOLS)/zipalign \
		-f 4 \
		build/app.apk.unaligned \
		"$@"

	$(_BUILD_TOOLS)/apksigner sign \
		--ks keys.store \
		--ks-pass pass:$(KEYS_PASS) \
		"$@"

clean:
	@rm -rf build
	@rm -rf $(_JAVA_ROOT_PATH)/R.java
