package jimsdk

import (
  "encoding/json"

  "github.com/parnurzeal/gorequest"
)

type ChangePasswordParams struct {
  OldPassword string `json:"old"`
  NewPassword string `json:"new"`
}

type ChangePasswordResponse struct {
  Result bool `json:"result"`
}

func (c *Client) SendChangePassword(oldPwd string, newPwd string) (*ChangePasswordResponse) {
  payload := ChangePasswordParams{ OldPassword: oldPwd, NewPassword: newPwd }

  resp, _, errs := gorequest.New().Post(c.ClusterURL + "/v1/users/change-password").
                                   Set("Content-Type", "application/json").
                                   Set("JIM-APP-ID", c.JimAppID).
                                   Set("JIM-APP-SIGN", c.getJimAppSign()).
                                   Send(payload).
                                   End()

  if errs != nil {
    return nil
  }

  respData := &ChangePasswordResponse{}

  if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
    return nil
  }

  return respData
}
