package jimsdk

import (
  "encoding/json"

  "github.com/parnurzeal/gorequest"
)

type VerifySmsParams struct {
  AppID int `json:"app-id"`
  Phone string `json:"phone"`
}

type VerifySmsResponse struct {
  Result bool `json:"result"`
}

func (c *Client) SendVerifySms(phone string) (*VerifySmsResponse) {
  payload := VerifySmsParams{ AppID: c.AppID, Phone: phone }

  resp, _, errs := gorequest.New().Post(c.ClusterURL + "/v1/users/send-verify-sms").
                                   Set("Content-Type", "application/json").
                                   Set("JIM-APP-ID", c.JimAppID).
                                   Set("JIM-APP-SIGN", c.getJimAppSign()).
                                   Send(payload).
                                   End()

  if errs != nil {
    return nil
  }

  respData := &VerifySmsResponse{}

	if err := json.NewDecoder(resp.Body).Decode(respData); err != nil {
    return nil
  }

  return respData
}
