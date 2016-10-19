package jimsdk

import (
  "encoding/json"

  "github.com/parnurzeal/gorequest"
)

type BindEmailParams struct {
  UserID int `json:"users-id"`
  Email string `json:"email"`
  VerificationCode string `json:"email-verification-code"`
}

type BindEmailResponse struct {
  Result bool `json:"result"`
}

func (c *Client) SendBindEmail(userID int, email string, verificationCode string) (*BindEmailResponse) {
  payload := BindEmailParams{ UserID: userID, Email: email, VerificationCode: verificationCode }

  resp, _, errs := gorequest.New().Post(c.ClusterURL + "/v1/users/bind-email").
                                   Set("Content-Type", "application/json").
                                   Set("JIM-APP-ID", c.JimAppID).
                                   Set("JIM-APP-SIGN", c.getJimAppSign()).
                                   Send(payload).
                                   End()

  if errs != nil {
    return nil
  }

  respData := &BindEmailResponse{}

  if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
    return nil
  }

  return respData
}
