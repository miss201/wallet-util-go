outDir=out/android
libName=libwallet.so

export GOARCH=arm
export GOOS=android
export CGO_ENABLED=1
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/darwin-x86_64/bin/armv7a-linux-androideabi21-clang
go build -ldflags "-w -s" -buildmode=c-shared -o $outDir/armeabi-v7a/$libName

echo "Build armeabi-v7a success"

export GOARCH=arm64
export GOOS=android
export CGO_ENABLED=1
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang
go build -ldflags "-w -s" -buildmode=c-shared -o $outDir/arm64-v8a/$libName

echo "Build arm64-v8a success"

export GOARCH=386
export GOOS=android
export CGO_ENABLED=1
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/darwin-x86_64/bin/i686-linux-android21-clang
go build -ldflags "-w -s" -buildmode=c-shared -o $outDir/x86/$libName

echo "Build x86 success"

export GOARCH=amd64
export GOOS=android
export CGO_ENABLED=1
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/darwin-x86_64/bin/x86_64-linux-android21-clang
go build -ldflags "-w -s" -buildmode=c-shared -o $outDir/x86_64/$libName

echo "Build x86_64 success"