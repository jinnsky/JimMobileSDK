package main

import (
	"fmt"
  "JimMobileSDK/jimsdk"
	"io/ioutil"
	"encoding/base64"
)

const clusterURL = "http://api2.jimyun.com"
const appID = 23
const jimAppID = "iu3TKjwRUCGfIwtTH9gXeYsq"
const jimAppSecret = "kJek81coyFG4V3eSg79b82HU"

type tempListener struct {}

func (t *tempListener) OnSuccess(respData *jimsdk.VerifyEmailResponse) {
	if respData.Result {
		fmt.Println("Sent verification email - OK.")
	} else {
		fmt.Println("Sent verification email - Failed.")
	}
}

func (t *tempListener) OnFailure(respErr *jimsdk.ResponseError) {
	fmt.Println(respErr.Key, respErr.Message)
}

func EncodeJPEGImageFile(path string) (string, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return "data:image/jpeg;base64," + base64.StdEncoding.EncodeToString(buff), nil
}

func main() {
	client, error := jimsdk.NewClient(clusterURL, appID, jimAppID, jimAppSecret, ".go-cookies")

	if error != nil {
		fmt.Println(error)
		return
	} 

	// Test request timeout
	client.RequestTimeout = 2

	listener := &tempListener{}
	client.SendVerifyEmailAsync("yangjingtian@oudmon.com", listener)	

	// Remove request timeout
	client.RequestTimeout = -1
	
	registerParams := &jimsdk.RegisterParams{ Username: "testerv1", 
																						Password: "123456", 
																						Email: "testerv1@oudmon.com" }	
	registerResponseData := client.SendRegister(registerParams)

	if registerResponseData != nil {
		if jimsdk.CatchResponseError(registerResponseData.Error) {
			fmt.Println(registerResponseData.Error.Key, registerResponseData.Error.Message)
		} else {
			fmt.Println("Username: ", registerResponseData.Username)	
			fmt.Println("UserID: ", registerResponseData.ID)
			fmt.Println("Register time: ", registerResponseData.RegisterTime)
		}
	}

	if client.HasValidSession() {
		infoResponseData := client.SendUserInfo(0, 0)
		
		if infoResponseData != nil {
			if jimsdk.CatchResponseError(infoResponseData.Error) {
				fmt.Println(infoResponseData.Error.Key, infoResponseData.Error.Message)
			} else {
				fmt.Println("Username: ", infoResponseData.Username)
				fmt.Println("UserID: ", infoResponseData.ID)
				fmt.Println("Register time: ", infoResponseData.RegisterTime)
				fmt.Println("Avatar URL: ", infoResponseData.AvatarURL)
			}
		}
	} else {
		loginParams := &jimsdk.LoginParams{ Username: "testerv2",
																				Password: "654321" }
		loginResponseData := client.SendLogin(loginParams)

		if loginResponseData != nil {
			if jimsdk.CatchResponseError(loginResponseData.Error) {
				fmt.Println(loginResponseData.Error.Key, loginResponseData.Error.Message)
			} else {
				fmt.Println("Username: ", loginResponseData.Username)	
				fmt.Println("UserID: ", loginResponseData.ID)
				fmt.Println("Register time: ", loginResponseData.RegisterTime)
				fmt.Println("Avatar URL: ", loginResponseData.AvatarURL)
			}
		}
	}

	changePasswordResponseData := client.SendChangePassword("123456", "654321")

	if changePasswordResponseData != nil {
		if jimsdk.CatchResponseError(changePasswordResponseData.Error) {
			fmt.Println(changePasswordResponseData.Error.Key, changePasswordResponseData.Error.Message)
		} else {
			if changePasswordResponseData.Result {
				fmt.Println("Changed password - OK.")
			} else {
				fmt.Println("Changed password - Failed.")
			}
		}
	}

	uploadAvatarResponseData := client.SendUploadAvatar("avatar.png")

	// encodedStr, _ := EncodeJPGImageFile("avatar.jpg")
	// uploadAvatarResponseData := client.SendUploadAvatarBase64(encodedStr)

	if uploadAvatarResponseData != nil {
		if jimsdk.CatchResponseError(uploadAvatarResponseData.Error) {
			fmt.Println(uploadAvatarResponseData.Error.Key, uploadAvatarResponseData.Error.Message)
		} else {
			fmt.Println("Avatar URL: ", uploadAvatarResponseData.URL)
			fmt.Println("Avatar message: ", uploadAvatarResponseData.Message)
		}
	}

	// feedbackSubmitResponseData := client.SendFeedback("tester@oudmon.com", "test feedback api")

	// if feedbackSubmitResponseData != nil {
	// 	if jimsdk.CatchResponseError(feedbackSubmitResponseData.Error) {
	// 		fmt.Println(feedbackSubmitResponseData.Error.Key, feedbackSubmitResponseData.Error.Message)
	// 	} else {
	// 		fmt.Println("ContactInfo: ", feedbackSubmitResponseData.ContactInfo)
	// 		fmt.Println("Feedback Content: ", feedbackSubmitResponseData.Content)
	// 	}
	// }

	newsDigestParams := &jimsdk.NewsDigestParams{ FromPage: 0, PageSize: 5, ThumbWidth: 200, ThumbHeight: 100, Language: "zh" }
	newsDigestResponseData := client.SendNewsDigest(newsDigestParams)

	if newsDigestResponseData != nil {
		if jimsdk.CatchResponseError(newsDigestResponseData.Error) {
			fmt.Println(newsDigestResponseData.Error.Key, newsDigestResponseData.Error.Message)
		} else {
			if len(newsDigestResponseData.Collection.Items) > 0 {
				fmt.Println(newsDigestResponseData.Collection.Items[0].ArticleURL)
			}
		}
	}

	client.SendLogout()
}
