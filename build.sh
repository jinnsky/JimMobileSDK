#!/bin/sh

on_fail() {
  echo "$*" >&2
  exit 1
}

prepare() {
  if ! command -v gomobile > /dev/null 2>&1; then
    on_fail "can't find gomobile"
  fi  
}

build_ios_framework() {
  echo "building framework for iOS..."
  gomobile bind -target=ios -o JimSdk.framework JimMobileSDK/jimsdk || on_fail "build iOS framework failed"
}

build_android_library() {
  echo "building library for Android..."
  gomobile bind -target=android -o JimSdk.aar JimMobileSDK/jimsdk || on_fail "build Android library failed"
}

copy_target() {
  if [ -d $IOS_FRAMEWORK_OUTPUT ]; then
    echo "copying iOS framework to example project..."
    cp -r -f $IOS_FRAMEWORK_OUTPUT $IOS_FRAMEWORK_DIRECTORY || on_fail "copy iOS framework failed"
  fi

  if [ -f $ANDROID_LIBRARY_OUTPUT ]; then
    echo "copying Android library to example project..."
    cp -f $ANDROID_LIBRARY_OUTPUT $ANDROID_LIBRARY_DIRECTORY || on_fail "copy Android library failed"
  fi
}

clean() {
  rm -rf $IOS_FRAMEWORK_OUTPUT
  rm -rf $ANDROID_LIBRARY_OUTPUT
}

usage() {
  echo ""
  exit 1
}

IOS_FRAMEWORK_OUTPUT="JimSdk.framework"
IOS_FRAMEWORK_DIRECTORY="./SdkExample_iOS"
ANDROID_LIBRARY_OUTPUT="JimSdk.aar"
ANDROID_LIBRARY_DIRECTORY="./SdkExample_Android/JimSdk"

prepare
clean
build_ios_framework
build_android_library
copy_target

echo "build successfully"

exit 0
