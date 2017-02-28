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
$ go get github.com/juju/persistent-cookiejar
$ go get golang.org/x/mobile/cmd/gomobile
$ gomobile init
```

## Build the project

Clone the project to $GOPATH/src

```
$ cd $GOPATH/src
$ git clone [repo]
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

## How to use the SDK

Create the Client instance

Swift version
```
struct SdkConstants {
    static let ClusterURL = "http://api2.jimyun.com"
    static let AppID = 23
    static let JimAppID = "iu3TKjwRUCGfIwtTH9gXeYsq"
    static let JimAppSecret = "kJek81coyFG4V3eSg79b82HU"
    static let CookieJarFile = "client_cookie_jar"
}

let documentsDirectory = FileManager.default.urls(for: .documentDirectory, in: .userDomainMask).first
let filepath = documentsDirectory?.appendingPathComponent(CookieJarFile).path
        
GoJimsdkNewClient(SdkConstants.ClusterURL,
                  SdkConstants.AppID,
                  SdkConstants.JimAppID,
                  SdkConstants.JimAppSecret,
                  filepath,
                  &client,
                  nil)
```

Java version
```
File cacheDir = this.getApplicationContext().getExternalCacheDir();
File cookieFile = new File(cacheDir, "client_cookie_jar");

Client client = null;

try {
    client = Jimsdk.newClient("http://api2.jimyun.com", 23, "iu3TKjwRUCGfIwtTH9gXeYsq", "kJek81coyFG4V3eSg79b82HU", cookieFile.getPath());
} catch (Exception e) {
    e.printStackTrace();
}
```

Create request parameters, send the request, and handle the response

Swift version
```
DispatchQueue.global().async { [weak self] in
    guard let `self` = self, let sdkClient = self.client else { return }
    guard let loginParams = GoJimsdkNewLoginParams() else { return }
    
    loginParams.setUsername("username")
    loginParams.setPassword("password")
    
    if let responseData = sdkClient.sendLogin(loginParams) {
        if let responseError = responseData.error() {
            DispatchQueue.main.async {
                print(responseError.message())
            }
        } else {
            DispatchQueue.main.async {
                // Update UI
                print(responseData.id_())
            }
        }
    } else {
        DispatchQueue.main.async {
            // Update UI
            print("Failed")
        }
    }
}
```

Java version
```
final Client finalClient = client;

new Thread(new Runnable() {
    @Override
    public void run() {
        LoginParams loginParams = Jimsdk.newLoginParams();
        loginParams.setUsername("username");
        loginParams.setPassword("password");

        final UserInfoResponse response = finalClient.sendLogin(loginParams);

        runOnUiThread(new Runnable() {
            @Override
            public void run() {
                // Update UI
                if (response.getError() != null) {
                    Toast.makeText(getApplicationContext(), "Login - Failed. " + response.getError().getMessage(), Toast.LENGTH_LONG).show();
                } else {
                    Toast.makeText(getApplicationContext(), "Login - OK. UserID = " + Long.toString(response.getID()), Toast.LENGTH_LONG).show();
                }
            }
        });
    }
});
```
