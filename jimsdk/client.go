package jimsdk

import (
  "crypto/md5"
	"encoding/hex"
  "encoding/json"
	"errors"
  "fmt"
	"strconv"
  "time"

  "github.com/parnurzeal/gorequest"
)

type Client struct {
  ClusterURL string
  AppID int
  JimAppID string
  JimAppSecret string
  ServerTimestampDiff int64
}

func (c *Client) getJimAppSign() (string) {
  serverTime := strconv.FormatInt(time.Now().UnixNano() / (1000 * 1000) + c.ServerTimestampDiff, 10)
  hasher := md5.New()
  hasher.Write([]byte(c.JimAppSecret + serverTime))

  return hex.EncodeToString(hasher.Sum(nil)) + "," + serverTime
}

func NewClient(clusterURL string, appID int, jimAppID string, jimAppSecret string) (*Client, error) {
  client := &Client {
    ClusterURL: clusterURL,
    AppID: appID,
    JimAppID: jimAppID,
    JimAppSecret: jimAppSecret,
  }

  if clusterURL == "" {
    return client, errors.New("ClusterURL must be indicated")
  }

  resp, _, errs := gorequest.New().
                             Post(clusterURL + "/v1/system/base-info").
                             Set("Content-Type", "application/json").
                             End()

  if errs != nil {
    return client, errors.New(clusterURL + " isn't reachable")
  }

  type apiResultType struct {
		Time int64 `json:"time"`
	}

  apiResult := &apiResultType{}
	err := json.NewDecoder(resp.Body).Decode(apiResult)

  if err != nil {
    return client, errors.New(clusterURL + " can't response basic info: " + err.Error())
  }

  client.ServerTimestampDiff = apiResult.Time - time.Now().UnixNano() / (1000 * 1000)

	return client, nil
}

type ResponseError struct {
  Key string `json:"key"`
  Message string `json:"message"`
}

func CatchResponseError(respError *ResponseError) bool {
    return (respError != nil) 
}

func (c *Client) processResponse(resp gorequest.Response, errs []error) *ResponseError {
  respData := &ResponseError{}

  if errs != nil {
    respData.Key = "Unexpected errors"
    respData.Message = fmt.Sprint(errs)

    return respData
  }

  if resp.StatusCode == 422 {
    if err := json.NewDecoder(resp.Body).Decode(respData); err == nil {
      return respData
    }
  }

  return nil
}
