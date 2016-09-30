package jimsdk

import (
  "encoding/json"
	"errors"
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
