package main

import (
	"fmt"
  "JimMobileSDK/jimsdk"
)

const clusterURL = "http://api2.jimyun.com"
const appID = 23
const jimAppID = "iu3TKjwRUCGfIwtTH9gXeYsq"
const jimAppSecret = "kJek81coyFG4V3eSg79b82HU"

type tempListener struct {}

func (t *tempListener) OnSuccess(respData *jimsdk.ResponseData) {
	if respData.Result {
		fmt.Println("Sent verification email - OK.")
	} else {
		fmt.Println("Sent verification email - Failed.")
	}
}

func (t *tempListener) OnFailure(err string) {
	fmt.Println(err)
}

func main() {
	client, error := jimsdk.NewClient(clusterURL, appID, jimAppID, jimAppSecret)

	if error != nil {
		fmt.Println(error)
	} 

	listener := &tempListener{}
	client.SendVerifyEmailAsync("yangjingtian@oudmon.com", listener)	
}
