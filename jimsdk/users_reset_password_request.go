package jimsdk

import (
  "encoding/json"
)

type ResetPasswordParams struct {
  AppID int `json:"app-id"`
  Phone string `json:"phone,omitempty"`
  Email string `json:"email,omitempty"`
  VerificationCode string `json:"code"`
  Password string `json:"password"`
}

func NewResetPasswordParams() (*ResetPasswordParams) {
  return &ResetPasswordParams{}
}

type ResetPasswordResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendResetPassword(params *ResetPasswordParams) (*ResetPasswordResponse) {
  if len(params.Password) > 0 {
    params.Password = c.getMD5String(params.Password)
  }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + ResetPasswordRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(params).
                                       End()
                                   
  respData := &ResetPasswordResponse{}
  
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
