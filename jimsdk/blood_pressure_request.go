package jimsdk

import (
	"encoding/json"
)

type BloodPressureCommitParams struct {
  SystolicBP float64 `json:"sbp"`
  DiastolicBP float64 `json:"dbp"`
  Pulse int `json:"pulse"`
  DeviceID string `json:"device-id"`
  DeviceType string `json:"device-type"`
  Time int64 `json:"time,omitempty"`
}

type BloodPressureCommitResponse struct {
  ID int64 `json:"id"`
  UserID int64 `json:"users-id"`
  SystolicBP float64 `json:"sbp"`
  DiastolicBP float64 `json:"dbp"`
  Pulse int `json:"pulse"`
  DeviceID string `json:"device-id"`
  DeviceType string `json:"device-type"`
  Time int64 `json:"time"`
  Error *ResponseError
}

func NewBloodPressureCommitParams() (*BloodPressureCommitParams) {
  return &BloodPressureCommitParams{}
}

type BloodPressureCommitListParams struct {
  LastSyncID int `json:"last-sync-id"`
  Items []*BloodPressureCommitParams `json:"datas"`
}

func NewBloodPressureCommitListParams() (*BloodPressureCommitListParams) {
  newBpclp := &BloodPressureCommitListParams{}
  newBpclp.Items = make([]*BloodPressureCommitParams, 0)

  return newBpclp
}

func (bpclp *BloodPressureCommitListParams) AddParams(params *BloodPressureCommitParams) {
  bpclp.Items = append(bpclp.Items, params)
}

type BloodPressureCommitResponseCollection struct {
  Items []BloodPressureCommitResponse
}

func (bpcrc *BloodPressureCommitResponseCollection) GetSize() int {
  return len(bpcrc.Items)
}

func (bpcrc *BloodPressureCommitResponseCollection) GetItemAt(index int) (*BloodPressureCommitResponse) {
  if index < 0 || index >= len(bpcrc.Items) {
    return nil
  }

  return &bpcrc.Items[index]
}

type BloodPressureCommitListResponse struct {
  Collection *BloodPressureCommitResponseCollection
  Error *ResponseError
}

func (c *Client) SendBloodPressureCommit(params *BloodPressureCommitParams) (*BloodPressureCommitResponse) {
  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + BloodPressureCommitRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(params).
                                       End()

  respData := &BloodPressureCommitResponse{}

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

func (c *Client) SendBloodPressureCommitList(paramsList *BloodPressureCommitListParams) (*BloodPressureCommitListResponse) {
  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + BloodPressureCommitListRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(paramsList).
                                       End()

  respData := &BloodPressureCommitListResponse{}

  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }

  respData.Collection = &BloodPressureCommitResponseCollection{}

  if err := json.NewDecoder(resp.Body).Decode(&respData.Collection.Items); err != nil {
    respData.Error = &ResponseError{ Key: "JSON Decode", Message: "Decode failed" }
  }

  return respData
}

func (c *Client) SendBloodPressureSync(lastSyncId int, size int) (*BloodPressureCommitListResponse) {
  var payload = struct {
    LastSyncID int `json:"last-sync-id"`
    Size int `json:"size"`
  } { lastSyncId, size }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + BloodPressureSyncRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &BloodPressureCommitListResponse{}

  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }

  respData.Collection = &BloodPressureCommitResponseCollection{}

  if err := json.NewDecoder(resp.Body).Decode(&respData.Collection.Items); err != nil {
    respData.Error = &ResponseError{ Key: "JSON Decode", Message: "Decode failed" }
  }

  return respData
}

type BloodPressureDeleteParams struct {
  Items []int64 `json:"ids"`
}

func NewBloodPressureDeleteParams() (*BloodPressureDeleteParams) {
  return &BloodPressureDeleteParams{}
}

func (bpdp *BloodPressureDeleteParams) AddDeleteID(id int64) {
  bpdp.Items = append(bpdp.Items, id)
}

func (c *Client) SendBloodPressureDelete(params *BloodPressureDeleteParams) (*BloodPressureCommitListResponse) {
  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + BloodPressureDeleteRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(params).
                                       End()

  respData := &BloodPressureCommitListResponse{}

  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }

  respData.Collection = &BloodPressureCommitResponseCollection{}

  if err := json.NewDecoder(resp.Body).Decode(&respData.Collection.Items); err != nil {
    respData.Error = &ResponseError{ Key: "JSON Decode", Message: "Decode failed" }
  }

  return respData
}

type BloodPressureTotalCountResponse struct {
  Count int64 `json:"count"`
  Error *ResponseError
}

func (c *Client) SendBloodPressureTotalCount() (*BloodPressureTotalCountResponse) {
  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + BloodPressureTotalCountRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       End()

  respData := &BloodPressureTotalCountResponse{}

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

func (c *Client) SendBloodPressureList(index int, size int) (*BloodPressureCommitListResponse) {
  var payload = struct {
    Index int `json:"index"`
    Size int `json:"size"`
  } {index, size}

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + BloodPressureListRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &BloodPressureCommitListResponse{}

  respErr := c.processResponse(resp, errs)
  if respErr != nil {
    respData.Error = respErr
    return respData
  }

  respData.Collection = &BloodPressureCommitResponseCollection{}

  if err := json.NewDecoder(resp.Body).Decode(&respData.Collection.Items); err != nil {
    respData.Error = &ResponseError{ Key: "JSON Decode", Message: "Decode failed" }
  }

  return respData
}

type BloodPressureDeviceTroubleResponse struct {
  ID int64 `json:"id"`
  AppID int64 `json:"app-id"`
  UserID int64 `json:"users-id"`
  DeviceType string `json:"device-type"`
  DeviceID string `json:"device-id"`
  TroubleType string `json:"trouble-type"`
  Time int64 `json:"time"`
  Error *ResponseError
}

func (c *Client) SendBloodPressureDeviceTrouble(deviceType string, 
                                                deviceID string, 
                                                trouble string, 
                                                time int64) (*BloodPressureDeviceTroubleResponse) {
  var payload = struct {
    DeviceType string `json:"device-type"`
    DeviceID string `json:"device-id"`
    Trouble string `json:"trouble"`
    Time int64 `json:"time"`
  } { deviceType, deviceID, trouble, time }

  resp, _, errs := c.getRequestAgent().Post(c.ClusterURL + BloodPressureDeviceTroubleRouter).
                                       Set("JIM-APP-SIGN", c.getJimAppSign()).
                                       Set("JIM-APP-ID", c.JimAppID).
                                       Send(payload).
                                       End()

  respData := &BloodPressureDeviceTroubleResponse{}

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
