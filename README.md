# JimMobileSDK

This project is using [Gomobile](https://github.com/golang/mobile) to build a cross platform SDK for iOS and Android.

## Get Golang environment prepared

```
$ brew install go
```

After installing, config your own $GOPATH by following `go help gopath`.

## Install the dependences

```
$ go get github.com/parnurzeal/gorequest
$ go get golang.org/x/mobile/cmd/gomobile
$ gomobile init
```

## Build the project

Clone the project to $GOPATH/src

```
$ cd $GOPATH/src
$ git clone https://github.com/jinnsky/JimMobileSDK.git
$ cd JimMobileSDK
```

Use the following command to generate a framework that can be used by Xcode project

```
$ gomobile bind -target=ios -o JimSdk.framework JimMobileSDK/jimsdk
```

Use the following command to generate a module that can be used by Android Studio project

```
$ gomobile bind -target=android -o JimSdk.aar JimMobileSDK/jimsdk
``` 