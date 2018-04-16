package main 

import "fmt"

func main(){
	fmt.Println("Hello World!")
}




/*C:\GoCode\testProject>set GOPATH=C:\GoCode\testProject

C:\GoCode\testProject>set GOBIN=C:\GoCode\testProject\bin*/


/*


g@DESKTOP-6FIH9CD MINGW64 ~
$ go get golang.org/x/mobile/cmd/gomobile

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile init

g@DESKTOP-6FIH9CD MINGW64 ~
$ ~/android-sdk/tools/android sdk
bash: /c/Users/g/android-sdk/tools/android: No such file or directory

g@DESKTOP-6FIH9CD MINGW64 ~
$


g@DESKTOP-6FIH9CD MINGW64 ~
$

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile build [-target android|ios] [-o output] [build flags] [package]
bash: ios]: command not found

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile build [-target android] [-o output] [build flags] [package]
usage: C:\Users\g\go\bin\gomobile.exe build [-target android|ios] [-o output] [b                                                    uild flags] [package]

Build compiles and encodes the app named by the import path.

The named package must define a main function.

The -target flag takes a target system name, either android (the
default) or ios.

For -target android, if an AndroidManifest.xml is defined in the
package directory, it is added to the APK output. Otherwise, a default
manifest is generated. By default, this builds a fat APK for all supported
instruction sets (arm, 386, amd64, arm64). A subset of instruction sets can
be selected by specifying target type with the architecture name. E.g.
-target=android/arm,android/386.

For -target ios, gomobile must be run on an OS X machine with Xcode
installed. Support is not complete.

If the package directory contains an assets subdirectory, its contents
are copied into the output.

The -o flag specifies the output file name. If not specified, the
output file name depends on the package built.

The -v flag provides verbose output, including the list of packages built.

The build flags -a, -i, -n, -x, -gcflags, -ldflags, -tags, and -work are
shared with the build command. For documentation, see 'go help build'.

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile version
gomobile version +2f6753b Mon Oct 23 18:25:25 2017 +0000 (android); androidSDK=

g@DESKTOP-6FIH9CD MINGW64 ~
$ export ANDROID_HOME=$HOME"/android-sdk"

g@DESKTOP-6FIH9CD MINGW64 ~
$ ~/android-sdk/tools/android sdk
bash: /c/Users/g/android-sdk/tools/android: No such file or directory

g@DESKTOP-6FIH9CD MINGW64 ~
$ go get golang.org/x/mobile/cmd/gomobile

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile init

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile command [arguments]
C:\Users\g\go\bin\gomobile.exe: unknown subcommand "command"
Run 'gomobile help' for usage.

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile clean

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile init [-u]

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile install [-target android] [build flags] [package]
usage: C:\Users\g\go\bin\gomobile.exe install [-target android] [build flags] [package]

Install compiles and installs the app named by the import path on the
attached mobile device.

Only -target android is supported. The 'adb' tool must be on the PATH.

The build flags -a, -i, -n, -x, -gcflags, -ldflags, -tags, and -work are
shared with the build command.
For documentation, see 'go help build'.

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile version

gomobile version +2f6753b Mon Oct 23 18:25:25 2017 +0000 (android); androidSDK=

g@DESKTOP-6FIH9CD MINGW64 ~
$

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile bind --target=android .
C:\Users\g\go\bin\gomobile.exe: no Android NDK path is set. Please run gomobile init with the ndk-bundle installed through the Android SDK manager or with the -ndk flag set.

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile init -ndk ~/Library/Android/sdk/ndk-bundle/
C:\Users\g\go\bin\gomobile.exe: "C:\\Users\\g\\Library\\Android\\sdk\\ndk-bundle" does not point to an Android NDK.

g@DESKTOP-6FIH9CD MINGW64 ~
$ ^C

g@DESKTOP-6FIH9CD MINGW64 ~
$ ^C

g@DESKTOP-6FIH9CD MINGW64 ~
$ export ANDROID_HOME=$HOME"/android-sdk"

g@DESKTOP-6FIH9CD MINGW64 ~
$ ~/AppData/Local/Android/Sdk/toolsLocal/Android/Sdk/tools
bash: /c/Users/g/AppData/Local/Android/Sdk/toolsLocal/Android/Sdk/tools: No such file or directory

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile init -ndk ~/AppData/Local/Android/Sdk/ndk-bundle/

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile build -target=android<PATH>
bash: syntax error near unexpected token `newline'

g@DESKTOP-6FIH9CD MINGW64 ~
$ gomobile build -target=android golang.org/x/mobile/example/basic

g@DESKTOP-6FIH9CD MINGW64 ~
$*/
