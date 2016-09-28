package main

import (
	"fmt"
  "JimMobileSDK/jimsdk"
)

func main() {
	client, error := jimsdk.NewClient("http://api2.jimyun.com", "", "", "")

	if error != nil {
		fmt.Println(error)
	} 

	fmt.Println(client.ServerTimestamp)
}
