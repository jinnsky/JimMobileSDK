package jimsdk

import (
  "encoding/json"
)

type QqUserParams struct {
  AppID int `json:"app-id"`
  OpenID string `json:"open-id"`
}

type QqUserResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendQqUser(openID string) (*QqUserResponse) {
  payload := QqUserParams{ AppID: c.AppID, OpenID: openID }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + QqUserRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(payload).
                                       End()

  respData := &QqUserResponse{}
  
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
