package jimsdk

import (
  "encoding/json"
  
  "github.com/parnurzeal/gorequest"
)

type BindPhoneParams struct {
  UserID int `json:"users-id"`
  Phone string `json:"phone"`
  VerificationCode string `json:"sms-verification-code"`
}

type BindPhoneResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendBindPhone(userID int, phone string, verificationCode string) (*BindPhoneResponse) {
  payload := BindPhoneParams{ UserID: userID, Phone: phone, VerificationCode: verificationCode }

  resp, _, errs := gorequest.New().Post(c.ClusterURL + "/v1/users/bind-phone").
                                   Set("Content-Type", "application/json").
                                   Set("JIM-APP-ID", c.JimAppID).
                                   Set("JIM-APP-SIGN", c.getJimAppSign()).
                                   Send(payload).
                                   End()
                                   
  respData := &BindPhoneResponse{}

  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }

  if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
    return nil
  }

  return respData
}
