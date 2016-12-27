package jimsdk

import (
  "encoding/json"
)

type UpdateBindPhoneParams struct {
  UserID int `json:"users-id"`
  Phone string `json:"phone"`
}

type UpdateBindPhoneResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendUpdateBindPhone(userID int, phone string) (*UpdateBindPhoneResponse) {
  payload := UpdateBindPhoneParams{ UserID: userID, Phone: phone }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + UpdateBindPhoneRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &UpdateBindPhoneResponse{}

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
