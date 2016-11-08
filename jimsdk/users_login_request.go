package jimsdk

import (
  "github.com/antonholmquist/jason"
	"github.com/parnurzeal/gorequest"
)

type LoginParams struct {
  AppID int `json:"app-id"`
  Username string `json:"username,omitempty"`
  Password string `json:"password,omitempty"`
  Email string `json:"email,omitempty"`
  Phone string `json:"phone,omitempty"`
  WeixinOpenID string `json:"weixin-openid,omitempty"`
  QqOpenID string `json:"qq-openid,omitempty"`
  SinaUID string `json:"sina-uid,omitempty"`
  FacebookID string `json:"facebook-id,omitempty"`
  TwitterID string `json:"twitter-id,omitempty"`
  LinkedInID string `json:"linkin-id,omitempty"`
}

func NewLoginParams() (*LoginParams) {
  return &LoginParams{}
}

type LoginResponseListener interface {
  OnSuccess(respData *UserInfoResponse)
  OnFailure(respErr *ResponseError)
}

func (c *Client) SendLogin(params *LoginParams) (*UserInfoResponse) {
  params.AppID = c.AppID

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + LoginRouter).
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

func (c *Client)SendLoginAsync(params *LoginParams, listener LoginResponseListener) {
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

        c.saveCookieJar()

        listener.OnSuccess(respData)
      }
    }
  }

  params.AppID = c.AppID

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + LoginRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(params).
                                       End(callback)

  if listener != nil {
    respErr := c.processResponse(resp, errs)

    if respErr != nil {
      listener.OnFailure(respErr)
    }
  }
}
