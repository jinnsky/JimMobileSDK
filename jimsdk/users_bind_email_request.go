package jimsdk

import (
  "encoding/json"
)

type BindEmailParams struct {
  UserID int `json:"users-id"`
  Email string `json:"email"`
  VerificationCode string `json:"email-verification-code"`
}

type BindEmailResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendBindEmail(userID int, email string, verificationCode string) (*BindEmailResponse) {
  payload := BindEmailParams{ UserID: userID, Email: email, VerificationCode: verificationCode }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + BindEmailRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(payload).
                                       End()

  respData := &BindEmailResponse{}

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
