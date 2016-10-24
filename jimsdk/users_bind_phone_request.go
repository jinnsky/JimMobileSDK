package jimsdk

import (
  "encoding/json"
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

  resp, _, errs := c.getRequest().Post(c.ClusterURL + BindPhoneRouter).
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
