#!/bin/sh
export LANG=en_US.UTF-8
outDir=out/ios
libNameSimulator=libwallet_simulator.a
libNameIphone=libwallet_iphone.a
libName=libwallet.a
headerName=libwallet.h
headerNameIphone=libwallet_iphone.h

rm -f $outDir/$libNameSimulator
rm -f $outDir/$libNameIphone
rm -f $outDir/$libName
rm -f $outDir/$headerName

export CFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphoneos --show-sdk-path)
CGO_LDFLAGS="-arch arm64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphoneos --show-sdk-path)
CGO_ENABLED=1 GOARCH=arm64 GOOS=darwin CC="clang $CFLAGS" go build -tags ios -ldflags=-w -trimpath -v -o $outDir/$libNameIphone -buildmode c-archive

cp ${outDir}/$headerNameIphone ${outDir}/$headerName

export CFLAGS="-arch x86_64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphonesimulator --show-sdk-path)
CGO_LDFLAGS="-arch x86_64 -miphoneos-version-min=9.0 -isysroot "$(xcrun -sdk iphonesimulator --show-sdk-path)
CGO_ENABLED=1 GOARCH=amd64 GOOS=darwin CC="clang $CFLAGS" go build -tags ios -ldflags=-w -trimpath -v -o $outDir/$libNameSimulator -buildmode c-archive

lipo -create  ${outDir}/$libNameIphone ${outDir}/$libNameSimulator -output ${outDir}/$libName
lipo -info ${outDir}/$libName