package jimsdk

import (
  "encoding/json"
)

type ResetPasswordEmailParams struct {
  AppID int `json:"app-id"`
  Email string `json:"email"`
}

type ResetPasswordEmailResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendResetPasswordEmail(email string) (*ResetPasswordEmailResponse) {
  payload := ResetPasswordEmailParams{ AppID: c.AppID, Email: email }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + ResetPasswordEmailRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()
                                   
  respData := &ResetPasswordEmailResponse{}
  
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
