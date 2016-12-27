package jimsdk

import (
  "encoding/json"
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

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + VerifySmsRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
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
