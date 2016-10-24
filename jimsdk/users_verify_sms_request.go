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
  Error *ResponseError
}

func (c *Client) SendVerifySms(phone string) (*VerifySmsResponse) {
  payload := VerifySmsParams{ AppID: c.AppID, Phone: phone }

  resp, _, errs := gorequest.New().Post(c.ClusterURL + VerifySmsRouter).
                                   Set("Content-Type", "application/json").
                                   Set("JIM-APP-ID", c.JimAppID).
                                   Set("JIM-APP-SIGN", c.getJimAppSign()).
                                   Send(payload).
                                   End()
  
  respData := &VerifySmsResponse{}

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
