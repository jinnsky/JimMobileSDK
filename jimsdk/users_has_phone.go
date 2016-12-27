package jimsdk

import (
	"encoding/json"
)

type HasPhoneParams struct {
  AppID int `json:"app-id"`
  Phone string `json:"phone"`
}

type HasPhoneResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendHasPhone(phone string) (*HasPhoneResponse) {
  payload := HasPhoneParams{ AppID: c.AppID, Phone: phone }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + HasPhoneRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &HasPhoneResponse{}
  
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
