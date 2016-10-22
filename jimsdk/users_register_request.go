package jimsdk

import (
  "github.com/parnurzeal/gorequest"
  "github.com/antonholmquist/jason"
)

type RegisterParams struct {
  AppID int `json:"app-id"`
  Username string `json:"username,omitempty"`
  Password string `json:"password"`
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
  Error *ResponseError
}

func (c *Client) SendRegister(params *RegisterParams) (*RegisterResponse) {
  params.AppID = c.AppID
  
  resp, _, errs := gorequest.New().Post(c.ClusterURL + "/v1/users/register").
                                   Set("Content-Type", "application/json").
                                   Set("JIM-APP-ID", c.JimAppID).
                                   Set("JIM-APP-SIGN", c.getJimAppSign()).
                                   Send(params).
                                   End()
                                   
  respData := &RegisterResponse{}
  
  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }
                                   
  v, _ := jason.NewObjectFromReader(resp.Body)
  
  id, _ := v.GetInt64("id")
  username, _ := v.GetString("username")
  registerTime, _ := v.GetInt64("register-time")
  emailChecked, _ := v.GetBoolean("email", "checked")
  email, _ := v.GetString("email", "email")
  phoneChecked, _ := v.GetBoolean("phone", "checked")
  phone, _ := v.GetString("phone", "phone")
  
  respData.ID = id
  respData.Username = username
  respData.RegisterTime = registerTime
  respData.Email = email
  respData.EmailChecked = emailChecked
  respData.Phone = phone
  respData.PhoneChecked = phoneChecked
  
  return respData
}
