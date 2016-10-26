package jimsdk

import (
  "encoding/json"
)

type WeiboUserParams struct {
  AppID int `json:"app-id"`
  SinaUID string `json:"sina-uid"`
}

type WeiboUserResponse struct {
  Result bool `json:"result"`
  Error *ResponseError
}

func (c *Client) SendWeiboUser(sinaUID string) (*WeiboUserResponse) {
  payload := WeiboUserParams{ AppID: c.AppID, SinaUID: sinaUID }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + WeiboUserRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Send(payload).
                                       End()

  respData := &WeiboUserResponse{}
  
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
