package jimsdk

import (
	"encoding/json"
)

type FirmwareOtaInfoResponse struct {
  Limit string `json:"limit"`
  DownloadURL string `json:"download-url"`
  Version string `json:"version"`
  HardwareVersion string `json:"hardware-version"`
  Description string `json:"desc"`
  Date string `json:"date"`
  Error *ResponseError
}

func (c *Client) SendFirmwareLastOta(curHwVer string) (*FirmwareOtaInfoResponse) {
  var payload = struct {
    HardwareVersion string `json:"hardware-version"`
  } { curHwVer }
  
  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + FirmwareLastOtaRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &FirmwareOtaInfoResponse{}

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

func (c *Client) SendFirmwareOtaInfoByVersion(curHwVer string, curFwVer string) (*FirmwareOtaInfoResponse) {
  var payload = struct {
    HardwareVersion string `json:"hardware-version"`
    Version string `json:"version"`
  } { curHwVer, curFwVer }
  
  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + FirmwareOtaInfoRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &FirmwareOtaInfoResponse{}

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
