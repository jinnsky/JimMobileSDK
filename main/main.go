package main

import (
	"fmt"
  "JimMobileSDK/jimsdk"
	"io/ioutil"
	"encoding/base64"
	"math"
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

	// hasPhoneResponseData := client.SendHasPhone("13923561187")

	// if hasPhoneResponseData != nil {
	// 	if jimsdk.CatchResponseError(hasPhoneResponseData.Error) {
	// 		fmt.Println(hasPhoneResponseData.Error.Key, hasPhoneResponseData.Error.Message)
	// 	} else {
	// 		if hasPhoneResponseData.Result {
	// 			fmt.Println("Phone number has been registered.")
	// 		} else {
	// 			fmt.Println("Phone number is OK to register.")
	// 		}
	// 	}
	// }

	// bpCommitParams := &jimsdk.BloodPressureCommitParams{DeviceID: "123456", 
	// 																										DeviceType: "BloodPressureTypeC",
	// 																										SystolicBP: 12.34,
	// 																										DiastolicBP: 56.78,
	// 																										Pulse: 90}
	// bpCommitResponseData := client.SendBloodPressureCommit(bpCommitParams)

	// if bpCommitResponseData != nil {
	// 	if jimsdk.CatchResponseError(bpCommitResponseData.Error) {
	// 		fmt.Println(bpCommitResponseData.Error.Key, bpCommitResponseData.Error.Message)
	// 	} else {
	// 		fmt.Println(bpCommitResponseData.ID, bpCommitResponseData.Time)
	// 	}
	// }

	bpCommitParams1 := &jimsdk.BloodPressureCommitParams{DeviceID: "123456", 
																											DeviceType: "BloodPressureTypeC",
																											SystolicBP: 12.34,
																											DiastolicBP: 56.78,
																											Pulse: 90,
																										  Time: 123456}

  bpCommitParams2 := &jimsdk.BloodPressureCommitParams{DeviceID: "123456", 
																											DeviceType: "BloodPressureTypeC",
																											SystolicBP: 12.35,
																											DiastolicBP: 56.79,
																											Pulse: 91,
																										  Time: 123457}

  bpCommitParams3 := &jimsdk.BloodPressureCommitParams{DeviceID: "123456", 
																											DeviceType: "BloodPressureTypeC",
																											SystolicBP: 12.36,
																											DiastolicBP: 56.70,
																											Pulse: 92,
																											Time: 123458}

  bpCommitParamsList := jimsdk.NewBloodPressureCommitListParams()
	bpCommitParamsList.LastSyncID = math.MaxUint32
	bpCommitParamsList.AddParams(bpCommitParams1)
	bpCommitParamsList.AddParams(bpCommitParams2)
	bpCommitParamsList.AddParams(bpCommitParams3)

	bpCommitListResponseData := client.SendBloodPressureCommitList(bpCommitParamsList)

	if bpCommitListResponseData != nil {
		if jimsdk.CatchResponseError(bpCommitListResponseData.Error) {
			fmt.Println(bpCommitListResponseData.Error.Key, bpCommitListResponseData.Error.Message)
		} else {
			if bpCommitListResponseData.Collection.GetSize() > 0 {
				fmt.Println("Upload blood pressure data - Success.")
			} else {
				fmt.Println("Upload blood pressure data - Response is empty.")
			}
		}
	}

	// bpSyncResponseData := client.SendBloodPressureSync(0, 2)

	// if bpSyncResponseData != nil {
	// 	if jimsdk.CatchResponseError(bpSyncResponseData.Error) {
	// 		fmt.Println(bpSyncResponseData.Error.Key, bpSyncResponseData.Error.Message)
	// 	} else {
	// 		if bpSyncResponseData.Collection.GetSize() > 1 {
	// 			fmt.Println(bpSyncResponseData.Collection.GetItemAt(1).ID)
	// 		}
	// 	}
	// }

	// bpDeleteParams := jimsdk.NewBloodPressureDeleteParams()
	// bpDeleteParams.AddDeleteID(4541)
	// bpDeleteResponseData := client.SendBloodPressureDelete(bpDeleteParams)

	// if bpDeleteResponseData != nil {
	// 	if jimsdk.CatchResponseError(bpDeleteResponseData.Error) {
	// 		fmt.Println(bpDeleteResponseData.Error.Key, bpDeleteResponseData.Error.Message)
	// 	} else {
	// 		if bpDeleteResponseData.Collection.GetSize() > 0 {
	// 			fmt.Println(bpDeleteResponseData.Collection.GetItemAt(0).ID)
	// 		}
	// 	}
	// }

	bpTotalCountResponseData := client.SendBloodPressureTotalCount()

	if bpTotalCountResponseData != nil {
		if jimsdk.CatchResponseError(bpTotalCountResponseData.Error) {
			fmt.Println(bpTotalCountResponseData.Error.Key, bpTotalCountResponseData.Error.Message)
		} else {
			fmt.Println("Blood pressure total count:", bpTotalCountResponseData.Count)
		}
	}

	bpListResponseData := client.SendBloodPressureList(0, 5)

	if bpListResponseData != nil {
		if jimsdk.CatchResponseError(bpListResponseData.Error) {
			fmt.Println(bpListResponseData.Error.Key, bpListResponseData.Error.Message)
		} else {
			if bpListResponseData.Collection.GetSize() > 0 {
				for _, item := range bpListResponseData.Collection.Items {
					fmt.Println(item.ID)
				}
			} else {
				fmt.Println("Blood pressure list is empty.")
			}
		}
	}

	client.SendLogout()
}
