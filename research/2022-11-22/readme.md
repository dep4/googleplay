# November 22 2022

## Android 11.0 (Google APIs)

then start the Google APIs device:

~~~
emulator -avd Pixel_3a_XL_API_30 -writable-system
~~~

and push:

~~~
adb root
adb remount
adb reboot
~~~

After two minutes, the device has not rebooted. Same result with `x86` or
`x86_64`:

https://issuetracker.google.com/issues/260134707

## Android 12.0 (Google APIs)

Start the Google Play device and pull:

~~~
adb pull /product/priv-app/Phonesky/Phonesky.apk Phonesky-31.apk
~~~

then start the Google APIs device:

~~~
emulator -avd Pixel_3a_XL_API_31 -writable-system
~~~

and push:

~~~
adb root
adb remount
adb reboot

adb root
adb remount
adb push Phonesky-31.apk /product/priv-app
adb reboot
~~~

## Android 12L (Google APIs)

Start the Google Play device and pull:

~~~
adb pull /product/priv-app/Phonesky/Phonesky.apk Phonesky-32.apk
~~~

then start the Google APIs device:

~~~
emulator -avd Pixel_3a_XL_API_32 -writable-system
~~~

and push:

~~~
adb root
adb remount
adb reboot

adb root
adb remount
adb push Phonesky-32.apk /product/priv-app
adb reboot
~~~

## Android Tiramisu (Google APIs)

Start the Google Play device and pull:

~~~
adb pull /product/priv-app/Phonesky/Phonesky.apk Phonesky-33.apk
~~~

then start the Google APIs device:

~~~
emulator -avd Pixel_3a_XL_API_33 -writable-system
~~~

and push:

~~~
adb root
adb remount
adb reboot

adb root
adb remount
adb push Phonesky-33.apk /product/priv-app
adb reboot
~~~

## Android Tiramisu (Google Play)

we cannot run as root:

~~~
> adb root
adbd cannot run as root in production builds
~~~
