package jimsdk

import (
  "encoding/json"
)

type ChangePasswordParams struct {
  OldPassword string `json:"old"`
  NewPassword string `json:"new"`
}

type ChangePasswordResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendChangePassword(oldPwd string, newPwd string) (*ChangePasswordResponse) {
  payload := ChangePasswordParams{ OldPassword: oldPwd, NewPassword: newPwd }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + ChangePasswordRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(payload).
                                       End()

  respData := &ChangePasswordResponse{}

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
