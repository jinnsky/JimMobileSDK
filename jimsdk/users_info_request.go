package jimsdk

import (
  "github.com/antonholmquist/jason"
	"github.com/parnurzeal/gorequest"
)

type UserInfoParams struct {
  UserID int `json:"users-id,omitempty"`
  SubUserID int `json:"sub-users-id,omitempty"`
}

type UserInfoResponseListener interface {
  OnSuccess(respData *UserInfoResponse)
  OnFailure(respErr *ResponseError)
}

func (c *Client) SendUserInfo(userID int, subUserID int) (*UserInfoResponse) {
  var payload = UserInfoParams{}

  if userID > 0 {
    payload.UserID = userID
  }

  if subUserID > 0 {
    payload.SubUserID = subUserID
  }

  if (subUserID > 0) {
    payload.SubUserID = subUserID
  }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + UserInfoRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(payload).
                                       End()

  respData := &UserInfoResponse{}
  
  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }

  obj, _ := jason.NewObjectFromReader(resp.Body)

  respData.ID, respData.Username, respData.RegisterTime, respData.AvatarURL,
  respData.Email, respData.EmailChecked, respData.Phone, respData.PhoneChecked, 
  respData.InfoBirthday, respData.InfoCaseHistory, respData.InfoNickname, 
  respData.InfoHeight, respData.InfoWeight, respData.InfoGender = c.decodeUserInfoObject(obj)

  return respData
}

func (c *Client) SendUserInfoAsync(userID int, subUserID int, listener UserInfoResponseListener) {
  callback := func (resp gorequest.Response, body string, errs []error) {
    if listener != nil {
      respErr := c.processResponse(resp, errs)

      if respErr != nil {
        listener.OnFailure(respErr)
      } else {
        respData := &UserInfoResponse{}

        obj, _ := jason.NewObjectFromReader(resp.Body)

        respData.ID, respData.Username, respData.RegisterTime, respData.AvatarURL,
        respData.Email, respData.EmailChecked, respData.Phone, respData.PhoneChecked, 
        respData.InfoBirthday, respData.InfoCaseHistory, respData.InfoNickname, 
        respData.InfoHeight, respData.InfoWeight, respData.InfoGender = c.decodeUserInfoObject(obj)

        listener.OnSuccess(respData)
      }
    }
  }

  var payload = UserInfoParams{}

  if userID > 0 {
    payload.UserID = userID
  }

  if subUserID > 0 {
    payload.SubUserID = subUserID
  }

  if (subUserID > 0) {
    payload.SubUserID = subUserID
  }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + UserInfoRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(payload).
                                       End(callback)

  if listener != nil {
    respErr := c.processResponse(resp, errs)

    if respErr != nil {
      listener.OnFailure(respErr)
    }
  }
}
