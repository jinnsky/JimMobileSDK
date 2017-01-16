package jimsdk

import (
	"encoding/json"
)

type SystemInfoResponse struct {
  Result bool
  Error *ResponseError
}

func (c *Client) SendSystemInfo(model string, os string, resolution string, network string) (*SystemInfoResponse) {
  type InfoObject struct {
    TerminalType string `json:"terminal-type"`
    OS string `json:"os"`
    Resolution string `json:"dpi"`
    NetworkType string `json:"network-type"`
  }
  
  infoObj := InfoObject{ model, os, resolution, network }

  var payload = struct {
    SystemInfo InfoObject `json:"system-info"`
  } { infoObj }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + SystemInfoRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &SystemInfoResponse{}

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
