package jimsdk

import (
  "github.com/antonholmquist/jason"
	"github.com/parnurzeal/gorequest"
)

type RegisterParams struct {
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
}

func NewRegisterParams() (*RegisterParams) {
  return &RegisterParams{}
}

type RegisterResponse struct {
  ID int64
  Username string
  RegisterTime int64
  Email string
  EmailChecked bool
  Phone string
  PhoneChecked bool
  InfoBirthday string
  InfoCaseHistory string
  InfoNickname string
  InfoHeight int
  InfoWeight int
  InfoGender int
  Error *ResponseError
}

type RegisterResponseListener interface {
  OnSuccess(respData *RegisterResponse)
  OnFailure(respErr *ResponseError)
}

func (c *Client) SendRegister(params *RegisterParams) (*RegisterResponse) {
  params.AppID = c.AppID

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + RegisterRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(params).
                                       End()
                                   
  respData := &RegisterResponse{}
  
  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }
                                   
  obj, _ := jason.NewObjectFromReader(resp.Body)

  respData.ID, respData.Username, respData.RegisterTime, 
  respData.Email, respData.EmailChecked, respData.Phone, respData.PhoneChecked, 
  respData.InfoBirthday, respData.InfoCaseHistory, respData.InfoNickname, 
  respData.InfoHeight, respData.InfoWeight, respData.InfoGender = c.decodeUserInfoObject(obj)

  c.saveCookieJar()
  
  return respData
}

func (c *Client) SendRegisterAsync(params *RegisterParams, listener RegisterResponseListener) {
  callback := func (resp gorequest.Response, body string, errs []error) {
    if listener != nil {
      respErr := c.processResponse(resp, errs)

      if respErr != nil {
        listener.OnFailure(respErr)
      } else {
        respData := &RegisterResponse{}

        obj, _ := jason.NewObjectFromReader(resp.Body)

        respData.ID, respData.Username, respData.RegisterTime, 
        respData.Email, respData.EmailChecked, respData.Phone, respData.PhoneChecked, 
        respData.InfoBirthday, respData.InfoCaseHistory, respData.InfoNickname, 
        respData.InfoHeight, respData.InfoWeight, respData.InfoGender = c.decodeUserInfoObject(obj)

        c.saveCookieJar()

        listener.OnSuccess(respData)
      }
    }
  }

  params.AppID = c.AppID

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + RegisterRouter).
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
