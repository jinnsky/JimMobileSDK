package jimsdk

import (
	"encoding/json"
)

type AppVersionInfoResponse struct {
  Type string `json:"type"`
  DownloadURL string `json:"download-url"`
  Limit string `json:"limit"`
  Current string `json:"curr"`
  Description string `json:"desc"`
  Error *ResponseError
}

func (c *Client) SendAppVersionInfo(appType string) (*AppVersionInfoResponse) {
  var payload = struct {
    AppID int `json:"app-id"`
    Type string `json:"type"`
  } { c.AppID, appType }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + AppVersionInfoRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &AppVersionInfoResponse{}

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
