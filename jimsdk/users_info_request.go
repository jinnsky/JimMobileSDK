package jimsdk

import (
  "github.com/antonholmquist/jason"
)

type UserInfoParams struct {
  UserID int `json:"users-id"`
  SubUserID int `json:"sub-users-id,omitempty"`
}

type UserInfoResponse struct {
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

func (c *Client) SendUserInfo(userID int, subUserID int) (*UserInfoResponse) {
  payload := UserInfoParams{ UserID: userID }

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

  v, _ := jason.NewObjectFromReader(resp.Body)
  
  id, _ := v.GetInt64("id")
  username, _ := v.GetString("username")
  registerTime, _ := v.GetInt64("register-time")
  emailChecked, _ := v.GetBoolean("email", "checked")
  email, _ := v.GetString("email", "email")
  phoneChecked, _ := v.GetBoolean("phone", "checked")
  phone, _ := v.GetString("phone", "phone")
  birthday, _ := v.GetString("info", "birthday")
  caseHistory, _ := v.GetString("info", "case-history")
  nickname, _ := v.GetString("info", "nickname")
  height, _ := v.GetInt64("info", "height")
  weight, _ := v.GetInt64("info", "weight")
  gender, _ := v.GetInt64("info", "sex")
  
  respData.ID = id
  respData.Username = username
  respData.RegisterTime = registerTime
  respData.Email = email
  respData.EmailChecked = emailChecked
  respData.Phone = phone
  respData.PhoneChecked = phoneChecked
  respData.InfoBirthday = birthday
  respData.InfoCaseHistory = caseHistory
  respData.InfoNickname = nickname
  respData.InfoHeight = int(height)
  respData.InfoWeight = int(weight)
  respData.InfoGender = int(gender)

  return respData
}
