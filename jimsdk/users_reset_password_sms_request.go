package jimsdk

import (
  "encoding/json"
)

type ResetPasswordSmsParams struct {
  AppID int `json:"app-id"`
  Phone string `json:"phone"`
}

type ResetPasswordSmsResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendResetPasswordSms(phone string) (*ResetPasswordSmsResponse) {
  payload := ResetPasswordSmsParams{ AppID: c.AppID, Phone: phone }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + ResetPasswordSmsRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()
                                   
  respData := &ResetPasswordSmsResponse{}
  
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
