package jimsdk

import (
  "github.com/antonholmquist/jason"
)

type RegisterInfoParams struct {
  AppID int `json:"app-id"`
  Username string `json:"username,omitempty"`
  Password string `json:"password,omitempty"`
  Phone string `json:"phone,omitempty"`
  Email string `json:"email,omitempty"`
  WeixinOpenID string `json:"weixin-openid,omitempty"`
  QqOpenID string `json:"qq-openid,omitempty"`
  SinaUID string `json:"sina-uid,omitempty"`
  FacebookID string `json:"facebook-id,omitempty"`
  TwitterID string `json:"twitter-id,omitempty"`
  LinkedInID string `json:"linkin-id,omitempty"`
  VerificationCode string `json:"code,omitempty"`
  Nickname string `json:"nickname,omitempty"`
  Height float32 `json:"height,omitempty"`
  Weight float32 `json:"weight,omitempty"`
  Gender int `json:"sex,omitempty"`
  CaseHistory string `json:"case-history,omitempty"`
  Birthday string `json:"birthday,omitempty"`
  AvatarURL string `json:"avatar-url,omitempty"`
}

func NewRegisterInfoParams() (*RegisterInfoParams) {
  return &RegisterInfoParams{}
}

func (c *Client) SendRegisterInfo(params *RegisterInfoParams) (*UserInfoResponse) {
  params.AppID = c.AppID

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + RegisterInfoRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(params).
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

  c.saveCookieJar()
  
  return respData
}
