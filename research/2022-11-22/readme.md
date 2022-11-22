# November 22 2022

## Android API 33 (Google APIs)

If you instead start the Google Play device and pull:

~~~
adb pull /system/priv-app/Phonesky/Phonesky.apk Phonesky.apk
~~~

then start the Google APIs device:

~~~
emulator -avd Pixel_3a_XL_API_24 -writable-system
~~~

and push:

~~~
adb root

adb remount
adb push Phonesky.apk /system/priv-app
adb reboot
~~~

we get this result:

> Google Play services has stopped

So the pulled `Phonesky.apk` is poisoned somehow. Get the `versionCode` from
Google Play device:

~~~
> adb shell dumpsys package com.android.vending | rg versionCode
versionCode=80671500 minSdk=14 targetSdk=23
~~~

we should try pulling from API 33 (Google Play) instead, so that the Play Store
version matches the one used with Open GApps.

## Android API 33 (Google Play)

If we start the Google Play device, writable:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_24 -writable-system
~~~

or not:

~~~
emulator -avd Pixel_3a_XL_API_24
~~~

we cannot run as root:

~~~
> adb root
adbd cannot run as root in production builds
~~~
