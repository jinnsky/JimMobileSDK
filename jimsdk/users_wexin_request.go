package jimsdk

import (
  "encoding/json"
)

type WexinUserParams struct {
  AppID int `json:"app-id"`
  OpenID string `json:"open-id"`
}

type WeixinUserResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendWeixinUser(openID string) (*WeixinUserResponse) {
  payload := WexinUserParams{ AppID: c.AppID, OpenID: openID }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + WeixinUserRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(payload).
                                       End()

  respData := &WeixinUserResponse{}
  
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
